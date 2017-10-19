package main

import (
	"fmt"
	"net/url"
)

type Balancer struct {
	Counter int
	BaseUrl *url.URL
}

func (b *Balancer) Init() {
	b.Counter = 0
}

func (b *Balancer) Increment() {
	b.Counter += 1
}

func (b *Balancer) Run(links_ch chan Link, to_crawl chan string, finish chan int) {

	visitedURL := map[string]bool{
	}

	for link := range links_ch {

		link_url, _ := url.Parse(link.Url)

		url := b.BaseUrl.ResolveReference(link_url)

		if url.Host == b.BaseUrl.Host {

			url_str :=  url.String()

			if !visitedURL[url_str] {
				//fmt.Println("Already been here.")
				visitedURL[url_str] = true

				b.Increment()
				fmt.Printf("%d  - %s \n", b.Counter, url_str)

				to_crawl <- url_str

			}
		}

	}
}
