package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("title", func(e *colly.HTMLElement) {
		title := e.Text
		fmt.Printf("Title: %s", title)
	})
	c.Visit("http://go-colly.org/")
}
