package support

import (
	"time"

	"github.com/tebeka/selenium"
)

// Loops until we find an element and return it (this might need a timeout later on)
func LoopUntilFindRowElement(we selenium.WebElement, cssselector string) selenium.WebElement {
	elem, err := we.FindElement(selenium.ByCSSSelector, cssselector)
	for err != nil {
		time.Sleep(time.Millisecond * 500)

		elem, err = we.FindElement(selenium.ByCSSSelector, cssselector)
	}

	return elem
}

// Loops until we find an element and return it (this might need a timeout later on)
func LoopUntilFindElement(driver *selenium.WebDriver, cssselector string) selenium.WebElement {
	elem, err := (*driver).FindElement(selenium.ByCSSSelector, cssselector)
	for err != nil {
		time.Sleep(time.Millisecond * 500)

		elem, err = (*driver).FindElement(selenium.ByCSSSelector, cssselector)
	}

	return elem
}

// Loops until we find a group of elements and return an array of them (this might need a timeout later on)
func LoopUntilFindElements(driver *selenium.WebDriver, cssselector string) []selenium.WebElement {
	elem, err := (*driver).FindElements(selenium.ByCSSSelector, cssselector)
	for err != nil {
		time.Sleep(time.Millisecond * 500)

		elem, err = (*driver).FindElements(selenium.ByCSSSelector, cssselector)
	}

	return elem
}
