package main

import (
	"fmt"
	"github.com/remotejob/bngfeeder/domains"
	"github.com/tebeka/selenium"
	"gopkg.in/gcfg.v1"
	"log"
	"time"
)

var mlogin = ""
var mpass = ""
var sites []string

func init() {

	var cfg domains.ServerConfig
	if err := gcfg.ReadFileInto(&cfg, "config.gcfg"); err != nil {
		log.Fatalln(err.Error())

	} else {

		mlogin = cfg.Login.Mlogin
		mpass = cfg.Pass.Mpass
		sites = cfg.Sites.Site

	}

}

func main() {

	caps := selenium.Capabilities{"browserName": "phantomjs"}
	wd, _ := selenium.NewRemote(caps, "")
	defer wd.Quit()

	wd.Get("https://login.live.com/login.srf?wa=wsignin1.0&rpsnv=12&ct=1465178498&rver=6.7.6636.0&wp=MBI&wreply=https:%2F%2Fwww.bing.com%2Fsecure%2FPassport.aspx%3Frequrl%3Dhttps%253a%252f%252fwww.bing.com%252fwebmaster%252fWebmasterManageSitesPage.aspx%253frflid%253d1&lc=1033&id=264960")

	time.Sleep(time.Millisecond * 2000)

	elem, err := wd.FindElement(selenium.ByID, "i0116")
	if err != nil {
		fmt.Println(err.Error())
	}
	pass, err := wd.FindElement(selenium.ByID, "i0118")
	if err != nil {
		fmt.Println(err.Error())
	}
	time.Sleep(time.Millisecond * 1000)

	err = elem.SendKeys(mlogin)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = pass.SendKeys(mpass)
	if err != nil {
		fmt.Println(err.Error())
	}
	btm, err := wd.FindElement(selenium.ByID, "idSIButton9")
	if err != nil {
		fmt.Println(err.Error())
	}
	btm.Click()
	time.Sleep(time.Millisecond * 5000)
	
//	title, err := wd.FindElement(selenium.ByTagName,"title")
	
	fmt.Println(wd.Title())
		

}
