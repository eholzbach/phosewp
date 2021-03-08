package plugins

import (
	"strings"

	irc "github.com/thoj/go-ircevent"
)

// Help provides basic usage instructions
func help(conn *irc.Connection, r string, event *irc.Event) {

	var query string
	var response string

	a := strings.Split(event.Message(), " ")

	if len(a) > 1 {
		query = a[1]
	}

	switch query {
	case "bs":
		response = "Corporate bullshit generator"
	case "coin":
		response = "Cryptocurrency data, accepts coinmarketcap codes"
	case "dict":
		response = "Queries a dictionary"
	case "fu":
		response = "FoaaS"
	case "insult":
		response = "Hurts in the feels"
	case "joke":
		response = "Dad jokes"
	case "kanye":
		response = "Kanye West"
	case "news":
		response = "Prints a recent article title from random garbage news source"
	case "quote":
		response = "Stores and fetches quotes ; <add string> to save ; <get> to fetch"
	case "ron":
		response = "Ron Swanson"
	case "startup":
		response = "Startup idea generator"
	case "urban":
		response = "Urban Dictionary"
	case "weather":
		response = "Returns weather conditials by zip code"
	default:
		response = "Commands are: bs, coin, dict, fu, insult, joke, kanye, news, quote, ron, startup, urban, weather"
	}

	conn.Privmsg(r, response)
}
