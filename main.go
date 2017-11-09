package main

import (
	"fmt"
	"net/http"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func scrapeImgLinks(rootLink string) ([]string, error) {
	
	// request and parse the front page
	resp, err := http.Get(rootLink)
	if err != nil {
		return nil, err
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	// Find all nodes with the <img> tag
	nodes := scrape.FindAll(root, scrape.ByTag(atom.Img))

	// Extract and store the href attributes 
	// Note that the actual links are 2-level up from the <img> element, hence .Parent.Parent 
	var imgLinks []string
	for _, aNode := range nodes {
		for _, attr := range aNode.Parent.Parent.Attr {
			if attr.Key == "href" {
				imgLinks = append(imgLinks, attr.Val)
				break
			}
		}
	}

	return imgLinks, err
}

func main() {
	
	links, err := scrapeImgLinks("https://www.bing.com/images/search?q=apples&FORM=HDRSC2")
	if err != nil {
		panic(err)
	}

	for _, aLink := range links {
		fmt.Println(aLink)
	}
}
