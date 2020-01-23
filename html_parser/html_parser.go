package html_parser

import (
	"golang.org/x/net/html"
	"io"
	"log"
	"strings"
)

// Link represents a link (<a href="...">Text</a>) in a html file
type Link struct {
	Href string
	Text string
}

func Parse(html_page io.Reader) []Link {

	doc, err := html.Parse(html_page)
	if err != nil {
		log.Panic("Package html could not parse")
	}

	// Search through doc tree for "a" tags
	result := searchHTMLElements(doc, []Link{})

	return result
}

func searchHTMLElements(n *html.Node, result []Link) []Link {
	if n.Type == html.ElementNode && n.Data == "a" {
		// do another dfs to get the text for any children
		text := strings.TrimSpace(searchTextElements(n, ""))
		href := ""
		for _, item := range n.Attr {
			if item.Key == "href" {
				href = item.Val
			}
		}
		result = append(result, Link{href, text})
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result = append(result, searchHTMLElements(c, []Link{})...)
	}

	return result
}

func searchTextElements(n *html.Node, result string) string {

	if n.Type == html.TextNode {
		result += n.Data
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result += searchTextElements(c, "")
	}

	return result
}
