package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("http://sense-sapporo.jp/release")
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	doc.Find(".release_content_text").Each(func(i int, s *goquery.Selection) {

		title := s.Find(".content_title_single").Text()

		url, _ := s.Find("a").Attr("href")

		fmt.Println(title, url)
	})

	/*
		exit := doc.Find(".content_title_single").Each(func(i int,* sText()
		println(exit)
	*/
}
