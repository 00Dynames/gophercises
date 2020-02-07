package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/00Dynames/gophercises/html_parser"
)

func main() {

	// Assign base url
	baseURL := os.Args[1]
	// TODO: check that the base url includes either http:// or https://
	fmt.Println(baseURL)

	// Build sitemap
	urls := buildSitemap(baseURL)
	//fmt.Println(urls)

	// Format xml
	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "	")
	if err := enc.Encode(); err != nil {
		log.Panic(err)
	}
	fmt.Println()
}

func buildSitemap(baseURL string) []string {

	// visited pages
	visited := make(map[string]bool)
	result := make([]string, 0)
	// pages to visit
	queue := make([]string, 0)
	queue = append(queue, baseURL)
	//fmt.Println(queue)

	httpPattern := regexp.MustCompile(fmt.Sprintf("^http://.+")) //"^%s", baseURL))
	basePattern := regexp.MustCompile(fmt.Sprintf("^%s", baseURL))

	currURL := ""
	// while pages to visit is not empty
	for len(queue) > 0 {
		//fmt.Println(queue)
		//fmt.Println(visited)
		currURL, queue = queue[0], queue[1:]

		// TODO: strings library has a hasPrefix function

		//fmt.Println(httpPattern.FindString(currURL))
		if httpPattern.FindString(currURL) == "" {
			currURL = baseURL + "/" + currURL
		}

		//fmt.Println(basePattern.FindString(currURL))
		if basePattern.FindString(currURL) == "" {
			//		fmt.Println("continue")
			continue
		}

		//		fmt.Println(currURL)

		if visited[currURL] {
			continue
		}

		// Mark currURL as visited
		visited[currURL] = true
		result = append(result, currURL)

		resp, err := http.Get(currURL)
		if err != nil {
			log.Panic(err)
		}

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
	return result
}
