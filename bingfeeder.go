package main

import (
	"fmt"
	"github.com/remotejob/bngfeeder/dbhandler/createdb"
	"github.com/remotejob/bngfeeder/dbhandler/find_unsubmited"
	"github.com/remotejob/bngfeeder/domains"
	"github.com/tebeka/selenium"
	"gopkg.in/gcfg.v1"
	"gopkg.in/mgo.v2"
	"log"
	"strconv"
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
	// FireFox driver without specific version
	dbsession, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer dbsession.Close()

	caps := selenium.Capabilities{"browserName": "chrome"}
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

	for _, site := range sites {

		exist := createdb.CheckIfExist(*dbsession, site)

		if exist {

			fmt.Println("Feed ", site)
			var i_dailyQuota int =0
			wd.Get("https://www.bing.com/webmaster/configure/submit/urls?url=http%3A%2F%2F" + site + "%2F")
			time.Sleep(time.Millisecond * 6000)

			dailyQuota, err := wd.FindElement(selenium.ByID, "dailyQuota")
			if err != nil {
				fmt.Println(err.Error())
			} else {

				dailyQuota_txt, err := dailyQuota.Text()
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("dailyQuota->", dailyQuota_txt)
					i_dailyQuota, err = strconv.Atoi(dailyQuota_txt)
					if err != nil {
						// handle error
						fmt.Println(err)
					}

				}

			}

			if i_dailyQuota > 0 {

				linkstosubmit := find_unsubmited.Find(*dbsession, site,i_dailyQuota)
				
				var strtoSend string
				
				if len(linkstosubmit) > 0 {

					for _, link := range linkstosubmit {

						strtoSend = strtoSend + "\n" + link
					}

					insertlinks, err := wd.FindElement(selenium.ByID, "urls")
					if err != nil {
						fmt.Println(err.Error())
					}
					insertlinks.Clear()
					insertlinks.SendKeys(strtoSend)
					//		insertlinks.SendKeys("")
					time.Sleep(time.Millisecond * 3000)

					btmsubmit, err := wd.FindElement(selenium.ByID, "addParam")
					if err != nil {
						fmt.Println(err.Error())
					}
					btmsubmit.Click()
					time.Sleep(time.Millisecond * 4000)

				}
			} else {
				
				fmt.Println("No more to submit",site)
			}
		}

	}

}
