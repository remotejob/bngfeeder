package main

import (
	"github.com/remotejob/bngfeeder/domains"
	"gopkg.in/gcfg.v1"
	"log"
	"fmt"
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
	
	for _,site := range sites {
		
		fmt.Println(site)
		
	}
	
	

}
