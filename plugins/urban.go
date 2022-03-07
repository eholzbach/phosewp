package plugins

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"

	irc "github.com/thoj/go-ircevent"
)

type apiResponse struct {
	Results []result `json:"list"`
	Tags    []string `json:"tags"`
	Type    string   `json:"result_type"`
}

type result struct {
	ID         int32  `json:"defid"`
	Author     string `json:"author"`
	Definition string `json:"definition"`
	Link       string `json:"permalink"`
	ThumbsDown int32  `json:"thumbs_down"`
	ThumbsUp   int32  `json:"thumbs_up"`
	Word       string `json:"word"`
}

// Urban provides garbage from urban dict
func urban(conn *irc.Connection, r string, event *irc.Event) {
	query := strings.TrimPrefix(event.Message(), "!urban ")
	response, err := defineWord(query)

	if err != nil {
		log.Println(err)
		return
	}

	if len(response.Results) <= 0 {
		conn.Privmsg(r, "not found")
		return
	}

	s := strings.Replace(response.Results[0].Definition, "[", "", -1)
	s = strings.Replace(s, "]", "", -1)

	conn.Privmsg(r, s)
}

// DefineWord looks up a word on urban dict
func defineWord(word string) (*apiResponse, error) {
	var r *apiResponse

	url := fmt.Sprintf("http://api.urbandictionary.com/v0/define?page=1&term=%s", url.QueryEscape(word))
	resp, err := getURL(url)

	if err != nil {
		return r, err
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&r); err != nil {
		return r, err
	}

	return r, nil
}
