package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	link_count := 0
	c := colly.NewCollector()

	c.OnHTML("title", func(e *colly.HTMLElement) {
		title := e.Text
		fmt.Printf("Title: %s\n", title)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// link := e.Attr("href")
		link_count += 1
	})

	c.Visit("http://go-colly.org/")

	fmt.Printf("Number of links: %d\n", link_count)
}
