package main

import{
  "fmt"
  "log"

  "github.com/gocolly/colly/v2"
}

func main() {
    // Create a new collector
    c := colly.NewCollector()

    // Set up a callback for when a visited HTML element is found
    c.OnHTML("article", func(e *colly.HTMLElement) {
        title := e.ChildText("h2") // Assume the title is within an <h2> tag
        link := e.ChildAttr("a", "href") // Assume there's a link in the article
        fmt.Printf("Title: %s\nLink: %s\n", title, link)
    })

    // Set up a callback for errors
    c.OnError(func(r *colly.Response, err error) {
        log.Printf("Request failed with response code %d: %s", r.StatusCode, err)
    })

    // Start the scraping process on a given URL
    err := c.Visit("https://example-blog.com")
    if err != nil {
        log.Fatal(err)
    }

    // Wait for the collector to finish
    c.Wait()
}
