// This is a headzoo Surf example applicaiton that grabs some information from
// Google and presents it.
//
// I'm attempting to make it Search a term, and list the reults to the terminal
// instead of all the XML goop that comes with HTML.
//
// Source URL: https://github.com/headzoo/surf/blob/master/docs/index.md
//
// GoDocs URL: https://godoc.org/github.com/headzoo/surf
//
// jQuery Selectors: https://api.jquery.com/category/selectors/
//
// SitePoint Easy to understand jQuery Selectors:
// https://www.sitepoint.com/comprehensive-jquery-selectors/
package main

import (
	"fmt"
	//"github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
)

// Used to reduce the number of basic Error checks in this script.
func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	bow := surf.NewBrowser()
	err := bow.Open("http://google.com")
	checkerr(err)

	fmt.Println(bow.Title())
	// https://www.sitepoint.com/comprehensive-jquery-selectors/ => Look for Attribute Selectors.
	search, err := bow.Form("[name='f']")
	checkerr(err)
	search.Input("q", "Cheese")
	err = search.Submit()
	checkerr(err)
	fmt.Println(bow.Body())
}
