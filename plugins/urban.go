//  Not worthwhile yet not worthless.

package plugins

import (
	"encoding/json"
	"fmt"
	"github.com/thoj/go-ircevent"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type APIResponse struct {
	Results []Result `json:"list"`
	Tags    []string `json:"tags"`
	Type    string   `json:"result_type"`
}

type Result struct {
	Id         int32  `json:"defid"`
	Author     string `json:"author"`
	Definition string `json:"definition"`
	Link       string `json:"permalink"`
	ThumbsDown int32  `json:"thumbs_down"`
	ThumbsUp   int32  `json:"thumbs_up"`
	Word       string `json:"word"`
}

func Urban(conn *irc.Connection, r string, event *irc.Event) {
	query := strings.TrimPrefix(event.Message(), "!urban ")
	response, err := DefineWord(query)

	if err != nil {
		fmt.Println(err)
		return
	}

	if len(response.Results) <= 0 {
		conn.Privmsg(r, "not found")
		return
	}

	for _, def := range response.Results {
		defResponse := fmt.Sprintf("%s", def.Definition)
		s := strings.Split(string(defResponse), "\n")
		lcount := 0
		tcount := 0
		for _, line := range s {
			if len(line) > 1 {
				conn.Privmsg(r, line)
				time.Sleep(300 * time.Millisecond)
				lcount += 1
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
func DefineWord(word string) (response *APIResponse, err error) {

	s := url.QueryEscape(word)
	endpoint := fmt.Sprintf("http://api.urbandictionary.com/v0/define?page=%d&term=%s", 1, s)
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&response)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
