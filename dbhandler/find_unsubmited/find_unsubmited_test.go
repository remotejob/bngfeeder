package find_unsubmited

import (
	"gopkg.in/mgo.v2"
	"testing"
	"fmt"
)

func TestFind(t *testing.T) {

	dbsession, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer dbsession.Close()

	linkstosubmit := Find(*dbsession, "kaukotuo.fi",5)

	var strtoSend string
	for _, link := range linkstosubmit {

		strtoSend = strtoSend + "\n" + link
	}

	fmt.Println(strtoSend)	
}
