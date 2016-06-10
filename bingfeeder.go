package main

import (
	"fmt"
	"github.com/remotejob/bngfeeder/dbhandler/find_unsubmited"
	"github.com/remotejob/bngfeeder/domains"
	"github.com/tebeka/selenium"
	"gopkg.in/gcfg.v1"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

var mlogin = ""
var mpass = ""

func init() {

	var cfg domains.ServerConfig
	if err := gcfg.ReadFileInto(&cfg, "config.gcfg"); err != nil {
		log.Fatalln(err.Error())

	} else {

		mlogin = cfg.Login.Mlogin
		mpass = cfg.Pass.Mpass

	}

}

// Errors are ignored for brevity.

func main() {
	// FireFox driver without specific version
	dbsession, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer dbsession.Close()

	linkstosubmit := find_unsubmited.Find(*dbsession, "kaukotuo.fi")

	var strtoSend string

	if len(linkstosubmit) > 0 {

		for _, link := range linkstosubmit {

			strtoSend = strtoSend + "\n" + link
		}

		fmt.Println(strtoSend)
		caps := selenium.Capabilities{"browserName": "chrome"}
		wd, _ := selenium.NewRemote(caps, "")
		defer wd.Quit()

		wd.Get("https://login.live.com/login.srf?wa=wsignin1.0&rpsnv=12&ct=1465178498&rver=6.7.6636.0&wp=MBI&wreply=https:%2F%2Fwww.bing.com%2Fsecure%2FPassport.aspx%3Frequrl%3Dhttps%253a%252f%252fwww.bing.com%252fwebmaster%252fWebmasterManageSitesPage.aspx%253frflid%253d1&lc=1033&id=264960")

		time.Sleep(time.Millisecond * 3000)

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
		time.Sleep(time.Millisecond * 10000)

		wd.Get("https://www.bing.com/webmaster/configure/submit/urls?url=http%3A%2F%2Fkaukotyo.fi%2F")

		time.Sleep(time.Millisecond * 4000)

		insertlinks, err := wd.FindElement(selenium.ByID, "urls")
		if err != nil {
			fmt.Println(err.Error())
		}
		insertlinks.Clear()
		insertlinks.SendKeys(strtoSend)
//		insertlinks.SendKeys("")		
		time.Sleep(time.Millisecond * 10000)

		btmsubmit, err := wd.FindElement(selenium.ByID, "addParam")
		if err != nil {
			fmt.Println(err.Error())
		}
		btmsubmit.Click()
		time.Sleep(time.Millisecond * 10000)

	}

	//	insertlinks,err := wd.FindElement(selenium., "i0116")

	//	// Enter code in textarea
	//	elem, _ := wd.FindElement(selenium.ByCSSSelector, "#code")
	//	elem.Clear()
	//	elem.SendKeys(code)
	//
	//	// Click the run button
	//	btn, _ := wd.FindElement(selenium.ByCSSSelector, "#run")
	//	btn.Click()
	//
	//	// Get the result
	//	div, _ := wd.FindElement(selenium.ByCSSSelector, "#output")
	//
	//	output := ""
	//	// Wait for run to finish
	//	for {
	//		output, _ = div.Text()
	//		if output != "Waiting for remote server..." {
	//			break
	//		}
	//		time.Sleep(time.Millisecond * 100)
	//	}
	//
	//	fmt.Printf("Got: %s\n", output)
}
