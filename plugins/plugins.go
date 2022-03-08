// Package plugins is the main handler for all plugins available
package plugins

import (
	"net/http"
	"strings"
	"time"

	"github.com/eholzbach/phosewp/config"
	irc "github.com/thoj/go-ircevent"
)

func getURL(url string) (*http.Response, error) {
	c := http.Client{
		Timeout: 10 * time.Second,
	}

	r, err := c.Get(url)
	return r, err
}

// Plugins function handles routing to all plugins
func Plugins(conn *irc.Connection, conf config.Vars) {
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
			simple(conn, r, event)
		case "!coin":
			coins(conn, r, event, conf)
		case "!dict":
			dict(conn, r, event, conf)
		case "!fu":
			foaas(conn, r, event)
		case "!help":
			help(conn, r, event)
		case "!news":
			news(conn, r, event, conf)
		case "!quote":
			quote(conn, r, event, conf)
		case "!ron":
			ron(conn, r, event)
		case "!urban":
			urban(conn, r, event)
		case "!weather":
			weather(conn, r, event, conf)
		default:
		}

		if strings.Contains(event.Message(), "http://") || strings.Contains(event.Message(), "https://") {
			urlz(conn, r, event)
		}
	})
}
