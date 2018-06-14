package main

import (
	"fmt"
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


func main() {
	dat, err := ioutil.ReadFile("ts.html")
	if err != nil {
		panic(err)
	}
	//fmt.Print(string(dat))

	doc, err := html.Parse(strings.NewReader(string(dat)))
	if err != nil {
		log.Fatal(err)
	}

	parse(doc)
}
