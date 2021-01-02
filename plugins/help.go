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
		response = "string ; Corporate bullshit generator"
	case "dict":
		response = "string ; Queries a dictionary"
	case "fu":
		response = "nil or string ; FoaaS"
	case "joke":
		response = "string ; Dad jokes"
	case "kanye":
		response = "string ; Kanye West"
	case "news":
		response = "nil or string ; Prints a recent article title from random garbage news source"
	case "quote":
		response = "add string to save ; get [id] to fetch quote"
	case "ron":
		response = "string ; Ron Swanson"
	case "urban":
		response = "string ; Urban Dictionary"
	case "weather":
		response = "zip code ; Returns the current temperature, weather condition, humidity, wind, 'feels like' temperature, barometric pressure, and visibility"
	case "wiki":
		response = "string ; Wikipedia"
	default:
		response = "Commands are: bs, dict, fu, kanye, news, quote, ron, urban, weather, wiki"
	}

	conn.Privmsg(r, response)
}
