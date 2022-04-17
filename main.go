package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
)

var tmpl *template.Template

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		panic(err)
	}
	t.Execute(w, "fodm")
	r.ParseForm()
	link := r.PostForm.Get("link")

	link_count := 0
	headings := make(map[string]int)
	c := colly.NewCollector()

	getTitle(c)

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

	c.Visit(link)

	fmt.Printf("Number of links: %d\n", link_count)
	fmt.Println(headings)
}

func getTitle(c *colly.Collector) {
	c.OnHTML("title", func(e *colly.HTMLElement) {
		title := e.Text
		fmt.Printf("Title: %s\n", title)
	})
}
