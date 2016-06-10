package find_unsubmited

import (
	//	"fmt"
	"github.com/remotejob/bngfeeder/domains"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func Find(dbsession mgo.Session, site string) []string {

	dbsession.SetMode(mgo.Monotonic, true)

	c := dbsession.DB("bingwebmaster").C("sites")
	var bngdb domains.BngDb

	err := c.Find(bson.M{"site": site}).One(&bngdb)
	if err != nil {

		log.Fatal(err)
	}

	var linkstosubmit []string

	var count = 0
	for i, page := range bngdb.Sitemappages {

		if !page.Submited {
			//			fmt.Println(page.Loc)
			bngdb.Sitemappages[i].Submited = true
			count = count + 1
			linkstosubmit = append(linkstosubmit, page.Loc)
		}

		if count > 4 {

			break
		}

	}
	//	dbtoinsert := domains.BngDb{site, bngdb.Sitemappages}

	err = c.Update(bson.M{"site": site}, bngdb)
	if err != nil {
		panic(err)
	}

	return linkstosubmit
}
