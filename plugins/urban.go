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

// urbandict json structs
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

func Urban(conn *irc.Connection, resp string, dword string) {

	response, err := DefineWord(dword)

	if err != nil {
		fmt.Println(err)
		return
	}

	if len(response.Results) <= 0 {
		conn.Privmsg(resp, "not found")
		return
	}

	for _, def := range response.Results {
		defResponse := fmt.Sprintf("%s", def.Definition)
		s := strings.Split(string(defResponse), "\n")
		lcount := 0
		for _, line := range s {
			if len(line) > 1 {
				conn.Privmsgf(resp, line)
				time.Sleep(300 * time.Millisecond)
				lcount += 1
				if lcount == 4 {
					time.Sleep(2 * time.Second)
					lcount = 0
				}
			}
		}
		break
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

