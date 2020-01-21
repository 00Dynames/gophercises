package html_parser

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
)

// Link represents a link (<a href="...">Text</a>) in a html file
type Link struct {
	Href string
	Text string
}

func Parse(html_page io.Reader) ([]Link, error) {

	doc, err := html.Parse(html_page)
	if err != nil {
		log.Panic("Package html could not parse")
	}

	// Search through doc tree for "a" tags
	search(doc)

	return nil, nil
}

func searchHTMLElements(n *html.Node) {
	if n.Type == html.ElementNode { //&& n.Data == "a" {
		fmt.Print(n.Data)
		fmt.Println(n.Namespace)
		// do another dfs to get the text for any children

	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		search(c)
	}
}
