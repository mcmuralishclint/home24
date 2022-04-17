package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

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
	// Read from index.gohtml
	t, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		panic(err)
	}
	t.Execute(w, "")

	// Get link inputted from index page
	r.ParseForm()
	link := r.PostForm.Get("link")

	getTitle(link)
	getLinks(link)
	getHeadings(link)
	loginForm(link)
}

func getTitle(link string) {
	c := colly.NewCollector()
	c.OnHTML("title", func(e *colly.HTMLElement) {
		title := e.Text
		fmt.Printf("Title: %s\n", title)
	})
	c.Visit(link)
}

func getLinks(link string) {
	link_count := 0
	accessible_link_count := 0
	c := colly.NewCollector(
		colly.MaxDepth(1),
		colly.Async(true),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		new_link := e.Attr("href")
		link_count += 1
		if strings.Contains(new_link, link) {
			fmt.Println(link)
			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	c.Visit(link)

	fmt.Printf("Number of links: %d, Number of accessible links: %d\n", link_count, accessible_link_count)
}

func getHeadings(link string) {
	c := colly.NewCollector()
	headings := make(map[string]int)

	// Output headings
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
	c.Visit(link)
	fmt.Println(headings)
}

func loginForm(link string) error {
	if link == "" {
		return nil
	}
	c := colly.NewCollector()

	// authenticate
	err := c.Post(link, map[string]string{"username": "admin", "password": "admin"})
	if err != nil {
		return nil
	}

	// attach callbacks after login
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
		if r.StatusCode == 200 {
			fmt.Println("Login form found")
		}
	})

	// start scraping
	c.Visit(link)
	return nil
}
