package main

import (
	"fmt"

	"net/url"
	"time"
)


func monitoring(links_ch chan Link, links_in chan Link, to_crawl chan string) {
	for {

		fmt.Printf("ch:%d in:%d crawl:%d\n", len(links_ch), len(links_in), len(to_crawl))

		time.Sleep(2 * time.Second)

	}

}




func main() {

	fmt.Println("START")


	links_ch := make(chan Link, 1000000)
	links_in := make(chan Link, 1000000)

	to_crawl := make(chan string, 200)
	finish :=  make (chan int)


	base := "https://ru.wikipedia.org"

	baseUrl, _ :=  url.Parse(base)

	balancer :=  Balancer{BaseUrl:baseUrl}

	parser := Parser{}


	limitator := Limitator{}

	go balancer.Run(links_ch, to_crawl, finish)

	go limitator.Limit(links_in, links_ch)

	go monitoring(links_ch, links_in, to_crawl)

	for i := 0 ; i < 10 ; i++ {
		go parser.Parse(to_crawl, links_in)
	}

	firstUrl := Link{Url:base}

	links_ch <- firstUrl

	<- finish




}
