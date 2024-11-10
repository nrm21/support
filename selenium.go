package support

import (
	"log"
	"time"

	"github.com/tebeka/selenium"
)

/*
To avoid this error:
	github.com/coreos/etcd/etcdserver/etcdserverpb imports
	google.golang.org/genproto/googleapis/api/annotations: ambiguous import: found package google.golang.org/genproto/googleapis/api/annotations in multiple modules:
	google.golang.org/genproto v0.0.0-20200513103714-09dca8ec2884 (C:\Users\Nate\go\pkg\mod\google.golang.org\genproto@v0.0.0-20200513103714-09dca8ec2884\googleapis\api\annotations)
	google.golang.org/genproto/googleapis/api v0.0.0-20241104194629-dd2ea8efbc28 (C:\Users\Nate\go\pkg\mod\google.golang.org\genproto\googleapis\api@v0.0.0-20241104194629-dd2ea8efbc28\annotations)

You must run these commands in the package you are using (it will need to be done once in every new package you make that uses 'nrm21/support' package):
	go get google.golang.org/genproto@latest
	go mod tidy
*/

// Loops until we find an element and return it
func LoopUntilFindRowElement(we selenium.WebElement, cssselector string, timeout int) selenium.WebElement {
	var elem selenium.WebElement
	var err error

	then := time.Now().Add(time.Duration(timeout) * time.Second)
	for {
		elem, err = we.FindElement(selenium.ByCSSSelector, cssselector)
		if err != nil {
			time.Sleep(500 * time.Millisecond)
			if time.Now().After(then) {
				log.Fatalln("Timeout finding element on page")
			}
		} else { // we found the element break from loop
			break
		}
	}

	return elem
}

// Loops until we find an element and return it
func LoopUntilFindElement(driver *selenium.WebDriver, cssselector string, timeout int) selenium.WebElement {
	var elem selenium.WebElement
	var err error

	then := time.Now().Add(time.Duration(timeout) * time.Second)
	for {
		elem, err = (*driver).FindElement(selenium.ByCSSSelector, cssselector)
		if err != nil {
			time.Sleep(500 * time.Millisecond)
			if time.Now().After(then) {
				log.Fatalln("Timeout finding element on page")
			}
		} else { // we found the element break from loop
			break
		}
	}

	return elem
}

// Loops until we find a group of elements and return an array of them
func LoopUntilFindElements(driver *selenium.WebDriver, cssselector string, timeout int) []selenium.WebElement {
	var elem []selenium.WebElement
	var err error

	then := time.Now().Add(time.Duration(timeout) * time.Second)
	for {
		elem, err = (*driver).FindElements(selenium.ByCSSSelector, cssselector)
		if err != nil {
			time.Sleep(500 * time.Millisecond)
			if time.Now().After(then) {
				log.Fatalln("Timeout finding element on page")
			}
		} else { // we found the element break from loop
			break
		}
	}

	return elem
}
