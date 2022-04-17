package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	link_count := 0
	headings := make(map[string]int)
	c := colly.NewCollector()

	c.OnHTML("title", func(e *colly.HTMLElement) {
		title := e.Text
		fmt.Printf("Title: %s\n", title)
	})

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		headings["h1"] += 1
	})
	c.OnHTML("h2", func(e *colly.HTMLElement) {
		headings["h2"] += 1
	})
	c.OnHTML("h3", func(e *colly.HTMLElement) {
		headings["h3"] += 1
	})
	c.OnHTML("h4", func(e *colly.HTMLElement) {
		headings["h4"] += 1
	})
	c.OnHTML("h5", func(e *colly.HTMLElement) {
		headings["h5"] += 1
	})
	c.OnHTML("h6", func(e *colly.HTMLElement) {
		headings["h6"] += 1
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// link := e.Attr("href")
		link_count += 1
	})

	c.Visit("http://go-colly.org/")

	fmt.Printf("Number of links: %d\n", link_count)
	fmt.Println(headings)
}
