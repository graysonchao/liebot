package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/nlopes/slack"

	"golang.org/x/net/html"
)

// Search returns an array of direct links to each comic URL returned by the given search query
func Search(query string) ([]string, error) {

	var results []string

	ONRURL := "http://www.ohnorobot.com/index.php?s=%s&Search=Search&comic=636"
	q := url.QueryEscape(query)
	queryURL := fmt.Sprintf(ONRURL, q)

	/*
		onrRes = searchONR(query)
		listRes = searchList(query)
	*/

	res, err := http.Get(queryURL)
	results = parseResponse(res)

	/*if len(results) < 1 {
		listSearch, err := http.Get(listURL)
		results = parseList(listSearch)
	}*/

	if err != nil {
		return results, err
	}
	return results, nil
}
func matchAttr(t html.Token, key string, test string) bool {
	for _, a := range t.Attr {
		if a.Key == key && strings.Index(a.Val, test) >= 0 {
			return true
		}
	}
	return false
}

func hasAttr(t html.Token, key string, value string) bool {
	for _, a := range t.Attr {
		if a.Key == key && a.Val == value {
			return true
		}
	}
	return false
}

func getHref(t html.Token) string {
	for _, a := range t.Attr {
		if a.Key == "href" {
			return a.Val
		}
	}
	return ""
}

func dedupe(strings []string) (results []string) {
	dedupeMap := make(map[string]bool)

	for _, s := range strings {
		dedupeMap[s] = true
	}

	for k := range dedupeMap {
		results = append(results, k)
	}

	return results
}

/*func parseList(r *http.Response) results []string {
	z := html.NewTokenizer(r.Body)
	for {
		tt := z.Next()
		switch {
			case tt == html.ErrorToken:
				return dedupe(results)
			case tt == html.StartTagToken:
				t := z.Token()
				isComicLink := t.Data == "a" && matchAttr(t, "href", "index.php")
		}
	}
}*/

func parseResponse(r *http.Response) (results []string) {
	z := html.NewTokenizer(r.Body)
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return dedupe(results)
		case tt == html.StartTagToken:
			t := z.Token()
			isSearchLink := t.Data == "a" && hasAttr(t, "class", "searchlink")
			if isSearchLink {
				results = append(results, getHref(t))
			}
		}
	}
}

func handleMessage(r *slack.RTM, mev *slack.MessageEvent, myID string) {

	mentionStr := fmt.Sprintf("<@%s>", myID)
	if strings.Index(mev.Text, mentionStr) != 0 {
		return
	}

	searchText := strings.Replace(mev.Text, mentionStr, "", -1)

	results, err := Search(searchText)

	cLink := "I didn't find anything."
	if len(results) > 0 && err == nil {
		indexLink := results[0]
		cLink = strings.Replace(indexLink, "index", "comic", 1)
	}

	msg := r.NewOutgoingMessage(cLink, mev.Channel)
	r.SendMessage(msg)
}

func getMyID(r *slack.RTM) (string, error) {
	users, err := r.GetUsers()
	for _, u := range users {
		if u.Name == "liebot" {
			return u.ID, nil
		}
	}
	return "", err
}

func main() {
	api := slack.New("xoxb-82948110903-Cr1MGXM72FaYNilQaZo6MM4D")
	logger := log.New(os.Stdout, "liebot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(true)

	rtm := api.NewRTM()

	myID, err := getMyID(rtm)
	if err != nil {
		panic(err)
	}
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				handleMessage(rtm, ev, myID)
			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop
			}
		}
	}
}
