package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	doc, err := goquery.NewDocument("https://www.monstercat.com/music")
	if err != nil {
		fmt.Print("url scarapping failed")
	}

	doc.Find("section[role=\"content\"]").Each(func(i int, s *goquery.Selection) {

		s.Find("div > div > ul > li").Each(func(n int, p *goquery.Selection) {
			tex := p.Find("a").Text()
			fmt.Print(tex)
		})

		url, _ := s.Find("a").Attr("href")

		fmt.Println(url)
	})

	/*
		exit := doc.Find(".content_title_single").Each(func(i int,* sText()
		println(exit)
	*/
}
