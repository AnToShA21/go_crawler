package main

import (
	"time"
	"container/list"
)

type Limitator struct {
}

func (l *Limitator) Limit(in_links chan Link, out_links chan Link) {

	visited := map[string]bool{
	}

	//list := make([]Link, 0, 10000000)

	var trueList list.List

	tick := time.Tick(100 * time.Millisecond)
	for {
		select {
		case <-tick:
			//if len(list) > 0 {
			//	var temp_links []Link
			//
			//	lenght := len(list)
			//
			//	temp_links, list = list[:lenght], list[lenght:]
			//
			//	for _, temp_link := range (temp_links) {
			//
			//		out_links <- temp_link
			//	}
			//
			//}

			if len(out_links) < 20000 {
				for e := trueList.Front(); e != nil; e = e.Next() {
					temp_link := e.Value.(Link)
					out_links <- temp_link
				}

			}


		case link := <-in_links:
			if !visited[link.Url] {
				visited[link.Url] = true

				//list = append(list, link)
				trueList.PushBack(link)
			}


		default:
			time.Sleep(50 * time.Millisecond)
		}
	}

}
