package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func BaiduHotSearch() {
	res, err := http.Get("https://m.dytt8.net/index2.htm")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("ul").Each(func(i int, s *goquery.Selection) {
		content := s.Find("a").Text()
		fmt.Printf("%d: %s\n", i, content)
	})

}

func main() {
	BaiduHotSearch()
}
