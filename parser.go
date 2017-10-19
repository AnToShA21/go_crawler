package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	//"fmt"
)

type Parser struct {
}

func (p *Parser) Parse(urls chan string, links chan Link) {

	for url := range urls {

		// Load the HTML document
		doc, err := goquery.NewDocument(url)
		if err != nil {
			log.Println(err)
			continue
		}

		// Find the review items
		doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {

			href, _ := s.Attr("href")

			//fmt.Printf("Href %d -   %s\n", i, href)

			link := Link{Url: href}

			links <- link

		})
	}
}
