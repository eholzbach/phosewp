package plugins

import (
	"encoding/json"
	"fmt"
	"github.com/thoj/go-ircevent"
	"net/http"
	"strings"
)

type gresp struct {
	Phrase string `json:"phrase"`
	Joke   string `json:"joke"`
	Quote  string `json:"quote"`
}

//
func Simple(conn *irc.Connection, r string, event *irc.Event) {

	var url string

	api := strings.Split(event.Message(), " ")

	switch api[0] {
	case "!bs":
		url = "https://corporatebs-generator.sameerkumar.website"
	case "!joke":
		url = "https://icanhazdadjoke.com/"
	case "!kanye":
		url = "https://api.kanye.rest"
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	a, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer a.Body.Close()
	var resp gresp
	json.NewDecoder(a.Body).Decode(&resp)

	switch api[0] {
	case "!bs":
		conn.Privmsg(r, resp.Phrase)
	case "!joke":
		conn.Privmsg(r, resp.Joke)
	case "!kanye":
		conn.Privmsg(r, resp.Quote)
	}
}
