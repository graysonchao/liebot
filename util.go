package liebot

import (
	"strings"

	"golang.org/x/net/html"
)

func getAttr(t html.Token, attrKey string) string {
	for _, a := range t.Attr {
		if a.Key == attrKey {
			return a.Val
		}
	}
	return ""
}

func indexToComic(link string) string {
	return strings.Replace(link, "index.php", "comic.php", 1)
}
