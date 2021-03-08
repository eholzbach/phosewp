// Package plugins is the main handler for all plugins available
package plugins

import (
	"strings"

	"github.com/eholzbach/phosewp/config"
	irc "github.com/thoj/go-ircevent"
)

// Plugins function handles routing to all plugins
func Plugins(conn *irc.Connection, conf *config.ConfigVars) {
	conn.AddCallback("PRIVMSG", func(event *irc.Event) {
		// reply target
		var r string

		if strings.HasPrefix(event.Arguments[0], "#") {
			// channel
			r = event.Arguments[0]
		} else {
			// user
			r = event.Nick
		}

		// get event prefix
		query := strings.Split(event.Message(), " ")

		switch query[0] {
		case "!bs", "!joke", "!insult", "!kanye", "!startup":
			Simple(conn, r, event)
		case "!coin":
			Coins(conn, r, event, conf)
		case "!dict":
			Dict(conn, r, event, conf)
		case "!fu":
			FoaaS(conn, r, event)
		case "!help":
			Help(conn, r, event)
		case "!news":
			News(conn, r, event, conf)
		case "!quote":
			Quote(conn, r, event, conf)
		case "!ron":
			Ron(conn, r, event)
		case "!urban":
			Urban(conn, r, event)
		case "!weather":
			Weather(conn, r, event, conf)
		default:
		}

		if strings.Contains(event.Message(), "http://") || strings.Contains(event.Message(), "https://") {
			Url(conn, r, event)
		}
	})
}
