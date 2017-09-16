package main

import (
	"fmt"
	//"github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
)

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
