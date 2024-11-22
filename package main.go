package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// Fetch the HTML document
func fetchHTML(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// Extract quotes and authors from the HTML document
func extractQuotesAndAuthors(node *html.Node) {
	var findQuotes func(*html.Node)
	findQuotes = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" {
			var isQuote, isAuthor bool
			for _, attr := range n.Attr {
				if attr.Key == "class" && attr.Val == "quote" {
					isQuote = true
				}
			}
			if isQuote {
				quoteText := ""
				authorName := ""
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.ElementNode && c.Data == "span" {
						for _, attr := range c.Attr {
							if attr.Key == "class" && attr.Val == "text" {
								quoteText = c.FirstChild.Data
							}
							if attr.Key == "class" && attr.Val == "author" {
								isAuthor = true
							}
						}
						if isAuthor && c.FirstChild != nil {
							authorName = c.FirstChild.Data
						}
					}
				}
				if quoteText != "" && authorName != "" {
					fmt.Printf("Quote: %s\nAuthor: %s\n\n", quoteText, authorName)
				}
			}
		}
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			findQuotes(child)
		}
	}
	findQuotes(node)
}

func main() {
	// URL of the quotes website
	url := "https://quotes.toscrape.com"

	// Fetch the HTML document
	doc, err := fetchHTML(url)
	if err != nil {
		fmt.Printf("Error fetching HTML: %v\n", err)
		return
	}

	// Extract and print quotes and authors
	fmt.Print("Quotes and Authors:\n")
	extractQuotesAndAuthors(doc)
}
