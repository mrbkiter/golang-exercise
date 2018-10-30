package linkmaster

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	customstrings "customutils/strings"

	html "golang.org/x/net/html"
)

func main() {
	//filepath := "D:/projects/golang/exercises/src/link-master/ex3.html"
	filepath := ".\\src\\link-master\\ex3.html"
	file, error := os.Open(filepath)
	if error != nil {
		panic(error)
	} else {

		root, error1 := html.Parse(file)
		pathMap := make(map[string]Link, 0)
		if error1 == nil {
			CollectLinks(root, pathMap, "", "")
			fmt.Println(pathMap)
		}
	}

}

//CollectLinks correct href, then put path to a map. Host would be used to filter out not related Href
func CollectLinks(n *html.Node, linkMap map[string]Link, host string, scheme string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" && a.Val != "#" {
				_url, error := url.Parse(a.Val)
				if error != nil {
					fmt.Printf(error.Error())
					return
				}
				if _url.Host == "" {
					//link := Link{Href: _url.Path, Text: n.FirstChild.Data, Host: host, Scheme: scheme}
					link := Link{Href: _url.Path, Host: host, Scheme: scheme}
					linkMap[BuildURL(scheme, host, _url.Path)] = link

				} else if _url.Host == host || customstrings.ContainEnd(_url.Host, "."+host) {
					_scheme := _url.Scheme
					if _url.Scheme == "" {
						_scheme = scheme
					}
					link := Link{Href: _url.Path, Host: _url.Host, Scheme: _scheme}
					linkMap[BuildURL(_scheme, _url.Host, _url.Path)] = link
					//*aList = append(*linkMap, link)
				} else if host == "" {
					link := Link{Href: _url.Path, Host: _url.Host, Scheme: _url.Scheme}
					linkMap[BuildURL(_url.Scheme, _url.Host, _url.Path)] = link
				} else {
					fmt.Printf("Host %s not accepted\n", _url.Host)
				}
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		CollectLinks(c, linkMap, host, scheme)
	}
}

//BuildURL build full url from scheme + host + path
func BuildURL(scheme string, host string, path string) string {
	return strings.Trim(scheme, " ") + "://" + strings.Trim(host, " ") + strings.Trim(path, " ")
}

//Link an object that hold href & Text of <a> element
type Link struct {
	Href   string
	Text   string
	Scheme string
	Host   string
}
