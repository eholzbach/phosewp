// Package plugins is the main handler for all plugins available
package plugins

import (
	"github.com/eholzbach/phosewp/config"
	"github.com/thoj/go-ircevent"
	"strings"
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
		case "!bs", "!joke", "!kanye":
			Simple(conn, r, event)
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
		case "!wiki":
			Wiki(conn, r, event)
		default:
		}

		if strings.Contains(event.Message(), "http://") || strings.Contains(event.Message(), "https://") {
			Url(conn, r, event)
		}
	})
}
