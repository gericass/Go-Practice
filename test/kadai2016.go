package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type Page struct {
	Title       string
	Description string
}

func getContent(nd *html.Node) string {
	cont := ""
	for _, v := range nd.Attr {
		if v.Key == "content" {
			cont = v.Val
		}
	}
	return cont
}

func checkAttr(nd *html.Node) string {
	for _, v := range nd.Attr {
		if v.Key == "name" && v.Val == "description" {
			cnt := getContent(nd)
			return cnt
		}
	}
	return "nil"
}

func Get(url string) (*Page, error) {
	p := Page{Title: "nil", Description: "nil"}
	res, err := http.Get(url)
	body, err := ioutil.ReadAll(res.Body)
	buf := bytes.NewBuffer(body)
	htmls := buf.String()
	h := strings.NewReader(htmls)
	doc, err := html.Parse(h)
	if err != nil {
		// 実際にはちゃんとエラー処理しましょう
		return nil, err
	}
	var f func(*html.Node)
	// fはDOMツリーを再帰的にトラバースするための手続きです。
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			p.Title = n.FirstChild.Data
		}
		if n.Type == html.ElementNode && n.Data == "meta" {
			attr := checkAttr(n)
			if p.Description == "nil" {
				p.Description = attr
			}
		}
		// 再帰的にノードをおっていくために、次のノードを探し、
		// ノードが存在すれば再びこのfを実行する、ということをしています。
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	if p.Description == "nil" || p.Title == "nil" {
		f(doc)
	}

	return &p, nil
}

func main() {
	// Getを利用しているとします
	p, err := Get("http://voyagegroup.com")
	if err != nil && err != io.EOF {
		panic(err)
	}
	fmt.Printf("%#v\n", p)
}
