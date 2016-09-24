package liebot

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func searchDialogue(query string) (results []string, err error) {

	ONRURL := "http://www.ohnorobot.com/index.php?s=%s&Search=Search&comic=636"
	q := url.QueryEscape(query)
	queryURL := fmt.Sprintf(ONRURL, q)

	r, err := http.Get(queryURL)
	if err != nil {
		return results, err
	}

	z := html.NewTokenizer(r.Body)
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return results, nil
		case tt == html.StartTagToken:
			t := z.Token()
			isSearchLink := t.Data == "a" && getAttr(t, "class") == "searchlink"
			if isSearchLink {
				results = append(results, indexToComic(getAttr(t, "href")))
			}
		}
	}
}

func searchTitles(query string) (results []string, err error) {

	LISTURL := "http://achewood.com/list.php"

	r, err := http.Get(LISTURL)

	n, err := html.Parse(r.Body)
	if err != nil {
		return results, err
	}

	// Recurse through the whole body and look for string matches
	var findTitles func(*html.Node)
	findTitles = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			tn := n.FirstChild
			queryLower := strings.ToLower(query)
			textLower := strings.ToLower(tn.Data)
			if tn.Type == html.TextNode && strings.Contains(textLower, queryLower) {
				for _, attr := range n.Attr {
					if attr.Key == "href" {
						results = append(results, "http://achewood.com/"+indexToComic(attr.Val))
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findTitles(c)
		}
	}

	findTitles(n)
	return results, nil
}
