package createdb

import (
	"fmt"
	"github.com/remotejob/bngfeeder/domains"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func Create(dbsession mgo.Session, site string, sitemap []domains.SitemapObj) {

	dbsession.SetMode(mgo.Monotonic, true)

	c := dbsession.DB("bingwebmaster").C("sites")

	count, err := c.Find(bson.M{"site": site}).Limit(1).Count()
	if err != nil {

		log.Fatal(err)
	}

	if count == 0 {

		dbtoinsert := domains.BngDb{site, sitemap}

		err := c.Insert(dbtoinsert)
		if err != nil {
			panic(err)
		}

	} else {

		fmt.Println("site DB exist ", site)

	}

}
