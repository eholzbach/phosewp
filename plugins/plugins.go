package plugins

import (
	"github.com/thoj/go-ircevent"
	"strings"
)

func Plugins(conn *irc.Connection, channels []string, wuapi string) {
	conn.AddCallback("PRIVMSG", func(event *irc.Event) {
		query := strings.Split(event.Message(), " ")
		switch query[0] {
		case "!help":
			Help(conn, event)
		}

		if strings.Contains(event.Message(), "http://") || strings.Contains(event.Message(), "https://") {
			Url(conn, event)
		}
	})
}
