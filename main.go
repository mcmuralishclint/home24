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
	http.ListenAndServe(":3001", r)
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

	GetHTMLVersion(link)
	fmt.Println()
	GetTitle(link)
	fmt.Println()
	GetHeadings(link)
	fmt.Println()
	LoginForm(link)
	fmt.Println()
	GetLinks(link)
}

func GetLinks(link string) error {
	link_count := 0
	c := colly.NewCollector(
		colly.MaxDepth(1),
		colly.Async(true),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link_count += 1
		fmt.Printf("Link %d: %s\n", link_count, e.Attr("href"))
	})
	c.Visit(link)
	return nil
}

func GetTitle(link string) error {
	if link == "" {
		return nil
	}
	c := colly.NewCollector()
	c.OnHTML("title", func(e *colly.HTMLElement) {
		title := e.Text
		fmt.Printf("Title: %s\n", title)
	})
	err := c.Visit(link)
	if err != nil {
		return err
	}
	return nil
}

func GetHeadings(link string) error {
	if link == "" {
		return nil
	}
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
	return nil
}

func LoginForm(link string) error {
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

func GetHTMLVersion(link string) error {
	if link == "" {
		return nil
	}
	version, err := HTMLVersion.DetectFromURL(link)
	if err != nil {
		fmt.Println("HTML Version not found")
	}
	fmt.Printf("HTML Version: %s\n", version)
	return nil
}
