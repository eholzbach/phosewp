package plugins

import (
	"github.com/thoj/go-ircevent"
	"strings"
)

func Plugins(conn *irc.Connection, channels []string, darksky string, newsapi string) {
	conn.AddCallback("PRIVMSG", func(event *irc.Event) {
		query := strings.Split(event.Message(), " ")
		switch query[0] {

		case "!acronym":
			Dict(conn, event)
		case "!drama":
			Dramatica(conn, event)
		case "!dict":
			Dict(conn, event)
		case "!fu":
			FoaaS(conn, event)
		case "!help":
			Help(conn, event)
		case "!news":
			News(conn, event, newsapi)
		case "!quote":
			Quote(conn, event)
		case "!stock":
			Stocks(conn, event)
		case "!trump":
			Tronald(conn, event)
		case "!urban":
			Urban(conn, event)
		case "!weather":
			Weather(conn, event, darksky)
		case "!wiki":
			Wiki(conn, event)
		default:
		}

		if strings.Contains(event.Message(), "http://") || strings.Contains(event.Message(), "https://") {
			Url(conn, event)
		}
	})
}
