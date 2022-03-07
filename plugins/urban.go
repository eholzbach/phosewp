package plugins

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

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

	for _, def := range response.Results {
		s := strings.Split(string(def.Definition), "\n")

		var lcount, tcount int

		for _, line := range s {
			if len(line) > 1 {
				conn.Privmsg(r, line)
				time.Sleep(300 * time.Millisecond)
				lcount++
				if lcount == 4 {
					time.Sleep(2 * time.Second)
					tcount += lcount
					lcount = 0
				}
				if tcount >= 40 {
					break
				}
			}
		}
	}
}

// DefineWord looks up a word on urban dict
func defineWord(word string) (response *apiResponse, err error) {
	s := url.QueryEscape(word)
	endpoint := fmt.Sprintf("http://api.urbandictionary.com/v0/define?page=%d&term=%s", 1, s)

	resp, err := http.Get(endpoint)

	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&response); err != nil {
		log.Println(err)
	}

	return
}
