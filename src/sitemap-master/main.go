package main

import (
	"encoding/xml"
	"fmt"
	"link-master"
	"log"
	"net/http"
	"net/url"

	html1 "golang.org/x/net/html"
)

func main() {
	input := "https://google.com"
	visitedMap := make(map[string]linkmaster.Link)
	_url, _ := url.Parse(input)
	buildSiteMap(input, visitedMap, &_url.Host, &_url.Scheme, 2)
	fmt.Println("Done build sitemap. Publishing sitemap ....")
	urlElements := make([]URLElement, 0)
	for k := range visitedMap {
		ele := URLElement{Loc: k}
		urlElements = append(urlElements, ele)
	}
	xmlRoot := SitemapXml{Urlset: &urlElements}
	bcontent, _ := xml.MarshalIndent(xmlRoot, "", "	")
	fmt.Println(string(bcontent))
}

func buildSiteMap(inputUrl string, visitedMap map[string]linkmaster.Link, host *string, scheme *string, depth int) {
	if depth <= 0 {
		return
	}
	fmt.Printf("Getting content from URL: %s\n", inputUrl)
	root, error := getContent(inputUrl)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	_url, _ := url.Parse(inputUrl)
	visitedMap[linkmaster.BuildURL(_url.Scheme, _url.Host, _url.Path)] =
		linkmaster.Link{Host: _url.Host, Href: _url.Path, Scheme: _url.Scheme}
	hrefMap := make(map[string]linkmaster.Link)
	hrefMap[linkmaster.BuildURL(_url.Scheme, _url.Host, _url.Path)] =
		linkmaster.Link{Host: _url.Host, Scheme: _url.Scheme, Href: _url.Path}

	linkmaster.CollectLinks(root, hrefMap, *host, *scheme)
	for k, v := range hrefMap {
		if _, ok := visitedMap[k]; !ok {
			buildSiteMap(linkmaster.BuildURL(v.Scheme, v.Host, v.Href), visitedMap, host, scheme, depth-1)
		}
	}
}

func getContent(url string) (*html1.Node, error) {
	resp, error := http.Get(url)

	if error != nil {
		log.Printf("*********** Error: %s *******************", error.Error())
		return nil, error
	}
	defer resp.Body.Close()
	return html1.Parse(resp.Body)
}

//SitemapXml present a sitemap structure
type SitemapXml struct {
	Urlset *[]URLElement `xml:"urlset"`
}

//URLElement present presentation of Url Element
type URLElement struct {
	Loc string `xml:"loc"`
}
