package plugins

import (
	"github.com/thoj/go-ircevent"
	"strings"
)

// Help provides basic usage instructions
func Help(conn *irc.Connection, r string, event *irc.Event) {

	var query string
	var response string

	a := strings.Split(event.Message(), " ")

	if len(a) > 1 {
		query = a[1]
	}

	switch query {
	case "bs":
		response = "nil ; Corporate bullshit generator"
	case "dict":
		response = "string ; Queries a dictionary"
	case "fu":
		response = "nil or string ; FoaaS"
	case "insult":
		response = "nil ; Hurts in the feels"
	case "joke":
		response = "nil ; Dad jokes"
	case "kanye":
		response = "nil ; Kanye West"
	case "news":
		response = "nil or string ; Prints a recent article title from random garbage news source"
	case "quote":
		response = "add string to save ; get [id] to fetch quote"
	case "ron":
		response = "nil ; Ron Swanson"
	case "startup":
		response = "nil ; Startup idea generator"
	case "urban":
		response = "string ; Urban Dictionary"
	case "weather":
		response = "zip code ; Returns the current temperature, weather condition, humidity, wind, 'feels like' temperature, barometric pressure, and visibility"
	case "wiki":
		response = "string ; Wikipedia"
	default:
		response = "Commands are: bs, dict, fu, insult, joke, kanye, news, quote, ron, startup, urban, weather, wiki"
	}

	conn.Privmsg(r, response)
}
