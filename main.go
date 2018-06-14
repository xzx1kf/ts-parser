package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func parse(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println(a.Val)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parse(c)
	}
}

func parseByTokenizer(r io.Reader) {
	z := html.NewTokenizer(r)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := z.Token()

			isAnchor := t.Data == "a"
			if isAnchor {
				fmt.Println("We found a link!")
				for _, a := range t.Attr {
					fmt.Println("\t"+a.Key+":"+a.Val)
				}
			}
		}
	}
}

func main() {
	dat, err := ioutil.ReadFile("ts.html")
	if err != nil {
		panic(err)
	}
	parseByTokenizer(strings.NewReader(string(dat)))
	//fmt.Print(string(dat))

	doc, err := html.Parse(strings.NewReader(string(dat)))
	if err != nil {
		log.Fatal(err)
	}

	parse(doc)
}
