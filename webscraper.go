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
    c.OnHTML("div.disease", func(e *colly.HTMLElement) {
        diseaseName := e.ChildText("h2.name") // Adjust the selector based on the actual HTML structure
        symptoms := e.ChildText("p.symptoms") // Adjust the selector based on the actual HTML structure

        // Print the disease name and symptoms
        fmt.Printf("Disease: %s\nSymptoms: %s\n\n", diseaseName, symptoms)
    })

    // Set up error handling
    c.OnError(func(r *colly.Response, err error) {
        log.Printf("Request URL: %s failed with response: %s\n", r.Request.URL, err)
    })

    // Start the scraping process on the target website
    err := c.Visit("https://example.com/diseases") // Replace with the actual URL
    if err != nil {
        log.Fatalf("Failed to visit: %s", err)
    }
}
