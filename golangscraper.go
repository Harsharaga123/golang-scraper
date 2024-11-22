package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

type Quote struct {
	Text   string   `json:"text"`
	Author string   `json:"author"`
	Tags   []string `json:"tags"`
}

func main() {
	baseURL := "https://quotes.toscrape.com"
	var quotes []Quote
	var mu sync.Mutex

	// Initialize Colly collector
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
		colly.Async(true),
	)

	// Scrape quotes
	c.OnHTML(".quote", func(e *colly.HTMLElement) {
		quote := Quote{
			Text:   e.ChildText(".text"),
			Author: e.ChildText(".author"),
			Tags:   e.ChildAttrs(".tags .tag", "href"),
		}

		mu.Lock()
		quotes = append(quotes, quote)
		mu.Unlock()
	})

	// Handle pagination
	c.OnHTML("li.next a", func(e *colly.HTMLElement) {
		nextPage := e.Request.AbsoluteURL(e.Attr("href"))
		c.Visit(nextPage)
	})

	// Error handling
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error: %s\n", err)
	})

	fmt.Println("Starting the scraper...")
	c.Visit(baseURL)
	c.Wait()

	// Save the scraped data
	saveToJSON(quotes, "quotes.json")
	saveToCSV(quotes, "quotes.csv")
	fmt.Println("Scraping complete. Check 'quotes.json' and 'quotes.csv' for results.")
}

func saveToJSON(data []Quote, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Could not create JSON file: %s", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		log.Fatalf("Could not write JSON data: %s", err)
	}
}

func saveToCSV(data []Quote, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Could not create CSV file: %s", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Text", "Author", "Tags"}
	writer.Write(headers)

	for _, quote := range data {
		record := []string{
			quote.Text,
			quote.Author,
			strings.Join(quote.Tags, "; "),
		}
		writer.Write(record)
	}
}
