package main

import (
	"encoding/json"
	"net/http"

	"github.com/graysonchao/liebot"
	"github.com/zenazn/goji"
)

func comicHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	v := r.Form

	slackRequest := &struct {
		Token       string
		TeamID      string
		TeamDomain  string
		ChannelID   string
		ChannelName string
		UserID      string
		UserName    string
		Command     string
		Text        string
		ResponseURL string
	}{
		Token:       v.Get("token"),
		TeamID:      v.Get("team_id"),
		TeamDomain:  v.Get("team_domain"),
		ChannelID:   v.Get("channel_id"),
		ChannelName: v.Get("channel_name"),
		UserID:      v.Get("user_id"),
		UserName:    v.Get("user_name"),
		Command:     v.Get("command"),
		Text:        v.Get("text"),
		ResponseURL: v.Get("response_url"),
	}

	comicLink := liebot.Search(slackRequest.Text)

	res := &struct {
		ResponseType string   `json:"response_type"`
		Text         string   `json:"text"`
		Attachments  []string `json:"attachments"`
	}{
		ResponseType: "in_channel",
		Text:         comicLink,
		Attachments:  []string{},
	}

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(res)
}

func main() {
	goji.Post("/comic", comicHandler)
	goji.Serve()
}
