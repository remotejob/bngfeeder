package createdb

import (
//	"fmt"
	"github.com/remotejob/bngfeeder/domains"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func CheckIfExist(dbsession mgo.Session, site string) bool {

	var exist bool = true
	
	dbsession.SetMode(mgo.Monotonic, true)

	c := dbsession.DB("bingwebmaster").C("sites")

	count, err := c.Find(bson.M{"site": site}).Limit(1).Count()
	if err != nil {

		log.Fatal(err)
	}

	if count == 0 {
//	  fmt.Println("site DB NOT exist ", site)	
      exist = false     
	}

	return exist
}

func Create(dbsession mgo.Session, site string, sitemap []domains.SitemapObj) {

	dbsession.SetMode(mgo.Monotonic, true)

	c := dbsession.DB("bingwebmaster").C("sites")

		dbtoinsert := domains.BngDb{site, sitemap}

		err := c.Insert(dbtoinsert)
		if err != nil {
			panic(err)
		}


}
