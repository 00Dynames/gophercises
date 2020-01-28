package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/00Dynames/gophercises/html_parser"
)

func main() {

	// Assign base url
	base_url := os.Args[1]
	fmt.Println(base_url)

	// Build sitemap
	urls := buildSitemap(base_url)

	// Format xml
}

func buildSitemap(base_url string) []string {

	// visited pages
	visited := make(map[string]bool)
	result := make([]string, 0)
	// pages to visit
	queue := make([]string, 0)
	queue = append(queue, base_url)

	// while pages to visit is not empty
	for len(queue) > 0 {
		curr_url := base_url

		resp, err := http.Get(curr_url)
		if err != nil {
			log.Panic(err)
		}

		// Mark curr_url as visited
		visited[curr_url] = true
		result = append(result, curr_url)

		// visit a new page and parse the links on the page
		links := html_parser.Parse(resp.Body)

		for _, link := range links {
			// if not visited
			if !visited[link.Href] {
				// add to pages to visit
				queue = append(queue, link.Href)
			}

		}
	}

	// return visited pages
	return nil
}
