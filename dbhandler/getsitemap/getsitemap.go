package getsitemap

import (
	"encoding/xml"
	//	"fmt"
	"github.com/remotejob/bngfeeder/domains"
	"io/ioutil"
	"log"
	"net/http"
)

//type SitemapObj struct {
//	Changefreq string
//	//	Hoursduration float64
//	Loc      string
//	Lastmod  string
//	Submited bool
//}
type Pages struct {
	//	Version string   `xml:"version,attr"`
	XMLName xml.Name `xml:"urlset"`
	XmlNS   string   `xml:"xmlns,attr"`
	//	XmlImageNS string   `xml:"xmlns:image,attr"`
	//	XmlNewsNS  string   `xml:"xmlns:news,attr"`
	Pages []*Page `xml:"url"`
}

type Page struct {
	XMLName    xml.Name `xml:"url"`
	Loc        string   `xml:"loc"`
	Lastmod    string   `xml:"lastmod"`
	Changefreq string   `xml:"changefreq"`
	//	Name       string   `xml:"news:news>news:publication>news:name"`
	//	Language   string   `xml:"news:news>news:publication>news:language"`
	//	Title      string   `xml:"news:news>news:title"`
	//	Keywords   string   `xml:"news:news>news:keywords"`
	//	Image      string   `xml:"image:image>image:loc"`
}

func GetUrl(link string) []domains.SitemapObj {

	resp, err := http.Get(link)
	if err != nil {
		// handle error
		log.Fatal(err)

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	//	fmt.Println(string(body))

	var sitemap Pages
	xml.Unmarshal(body, &sitemap)

	//	fmt.Println(sitemap)

	var sitemapObjs []domains.SitemapObj

	for _, page := range sitemap.Pages {

		//		fmt.Println(page)
		sitemapObj := domains.SitemapObj{
			Loc:        page.Loc,
			Lastmod:    page.Lastmod,
			Changefreq: page.Changefreq,
			Submited:   false,
		}
		sitemapObjs = append(sitemapObjs, sitemapObj)

	}

	return sitemapObjs
}
