package main

import (
	"fmt"
	"link-master"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

func main() {
	input := "http://www.sitemaps.org/protocol.html"
	visitedMap := make(map[string]linkmaster.Link)
	_url, _ := url.Parse(input)
	buildSiteMap(input, visitedMap, &_url.Host, &_url.Scheme)
	fmt.Println(visitedMap)

}

func buildSiteMap(inputUrl string, visitedMap map[string]linkmaster.Link, host *string, scheme *string) {
	resp, error := http.Get(inputUrl)

	if error != nil {
		log.Printf("*********** Error: %s *******************", error.Error())
	} else {
		defer func() {
			fmt.Printf("Defered called")
			resp.Body.Close()
		}()
		_url, _ := url.Parse(inputUrl)
		visitedMap[linkmaster.BuildURL(_url.Scheme, _url.Host, _url.Path)] =
			linkmaster.Link{Host: _url.Host, Href: _url.Path, Scheme: _url.Scheme}
		hrefMap := make(map[string]linkmaster.Link)
		hrefMap[linkmaster.BuildURL(_url.Scheme, _url.Host, _url.Path)] =
			linkmaster.Link{Host: _url.Host, Scheme: _url.Scheme, Href: _url.Path}
		root, _ := html.Parse(resp.Body)
		linkmaster.CollectLinks(root, hrefMap, *host, *scheme)
		for k, v := range hrefMap {
			if _, ok := visitedMap[k]; !ok {
				buildSiteMap(linkmaster.BuildURL(v.Scheme, v.Host, v.Href), visitedMap, host, scheme)
			}
		}
	}
}
