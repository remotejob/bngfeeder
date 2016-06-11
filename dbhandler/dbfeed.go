package main

import (
	"fmt"
	"github.com/remotejob/bngfeeder/dbhandler/createdb"
	"github.com/remotejob/bngfeeder/dbhandler/getsitemap"
	"github.com/remotejob/bngfeeder/domains"
	"gopkg.in/gcfg.v1"
	"gopkg.in/mgo.v2"
	"log"
)

var sites []string

func init() {

	var cfg domains.ServerConfig
	if err := gcfg.ReadFileInto(&cfg, "config.gcfg"); err != nil {
		log.Fatalln(err.Error())

	} else {

		sites = cfg.Sites.Site

	}

}

func main() {

	dbsession, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer dbsession.Close()

	for _, site := range sites {

		exist := createdb.CheckIfExist(*dbsession, site)

		if !exist {
			
			urlstr := "http://"+site+"/sitemap.xml"
			fmt.Println("Create", site,urlstr)

			sitemapObjs := getsitemap.GetUrl(urlstr)

			createdb.Create(*dbsession, site, sitemapObjs)

		}

	}

}
