package createdb

import (
	"gopkg.in/mgo.v2"
	"testing"
	"github.com/remotejob/bngfeeder/dbhandler/getsitemap"	
)

func TestCreate(t *testing.T) {

	dbsession, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer dbsession.Close()
	
	sitemapObjs := getsitemap.GetUrl("http://kaukotyo.fi/sitemap.xml")
		
	Create(*dbsession,"kaukotyo.fi",sitemapObjs )	

}
