package getsitemap

import (
    "testing"
)

func TestGetUrl(t *testing.T) {
	
	sitemapObjs := GetUrl("http://kaukotyo.fi/sitemap.xml")
	
	if len(sitemapObjs) <10 {
		
		t.Fatal("len(sitemapObjs) must be more 10")
	} 	

}

