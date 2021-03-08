package plugins

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	irc "github.com/thoj/go-ircevent"
)

type gresp struct {
	Insult string `json:"insult"`
	Joke   string `json:"joke"`
	Phrase string `json:"phrase"`
	Quote  string `json:"quote"`
	That   string `json:"that"`
	This   string `json:"this"`
}

// simple makes calls to simple api's
func simple(conn *irc.Connection, r string, event *irc.Event) {

	var url string

	api := strings.Split(event.Message(), " ")

	switch api[0] {
	case "!bs":
		url = "https://corporatebs-generator.sameerkumar.website"
	case "!insult":
		url = "https://evilinsult.com/generate_insult.php?lang=en&type=json"
	case "!joke":
		url = "https://icanhazdadjoke.com/"
	case "!kanye":
		url = "https://api.kanye.rest"
	case "!startup":
		url = "http://itsthisforthat.com/api.php?json"
	}

	// had to set the header for dad jokes
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	var resp gresp
	json.NewDecoder(res.Body).Decode(&resp)

	switch api[0] {
	case "!bs":
		conn.Privmsg(r, resp.Phrase)
	case "!insult":
		conn.Privmsg(r, resp.Insult)
	case "!joke":
		conn.Privmsg(r, resp.Joke)
	case "!kanye":
		conn.Privmsg(r, resp.Quote)
	case "!startup":
		conn.Privmsg(r, strings.ToLower(fmt.Sprintf("so, basically, it's like a %s for %s.", resp.This, resp.That)))
	}
}
