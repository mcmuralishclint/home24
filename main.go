package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
	HTMLVersion "github.com/lestoni/html-version"
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

	getHTMLVersion(link)
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
	// link_count := 0
	// accessible_link_count := 0
	// // Instantiate default collector
	// c := colly.NewCollector()

	// // On every a element which has href attribute call callback
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	new_link := e.Attr("href")
	// 	// Print link
	// 	fmt.Printf("Link found: %q -> %s\n", e.Text, new_link)
	// 	c.Visit(e.Request.AbsoluteURL(new_link))
	// })

	// c.OnResponse(func(r *colly.Response) {
	// 	link_count += 1
	// 	fmt.Println(link_count)
	// })

	// c.Visit(link)

	// fmt.Printf("Number of links: %d, Number of accessible links: %d\n", link_count, accessible_link_count)
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
		if r.StatusCode == 200 {
			fmt.Println("Login form found")
		}
	})

	// start scraping
	c.Visit(link)
	return nil
}

func getHTMLVersion(link string) {
	version, err := HTMLVersion.DetectFromURL(link)
	if err != nil {
		fmt.Println("HTML Version not found")
	}
	fmt.Printf("HTML Version: %s\n", version)
}
