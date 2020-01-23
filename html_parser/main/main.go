package main

import (
	"fmt"
	"github.com/00Dynames/gophercises/html_parser"
	"log"
	"os"
)

func main() {

	html, err := os.Open(os.Args[1])
	if err != nil {
		log.Panic("Cannot open the given html file")
	}

	fmt.Println(html_parser.Parse(html))
}
