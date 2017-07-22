package events

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"strings"
)

func Global(conn *irc.Connection, channels []string, handle string, wuapi string) {

	conn.AddCallback("001", func(event *irc.Event) {
		for _, channel := range channels {
			conn.Join(channel)
		}
	})

	conn.AddCallback("INVITE", func(event *irc.Event) {
		if len(event.Arguments) != 2 {
			return
		}
		conn.Join(event.Arguments[1])
		fmt.Println("INVITE " + strings.Join(event.Arguments, " "))
	})

	conn.AddCallback("KICK", func(event *irc.Event) {
		if event.Arguments[1] == handle {
			conn.Join(event.Arguments[0])
			s := fmt.Sprint("eat a bag of dicks, ", event.Nick, ".")
			conn.Privmsg(event.Arguments[0], s)
			fmt.Println("KICKED " + strings.Join(event.Arguments, " "))
		}
	})

	conn.AddCallback("PRIVMSG", func(event *irc.Event) {
		if strings.HasPrefix(event.Message(), "!help") == true {

			var replyto string
			var query string

			if strings.HasPrefix(event.Arguments[0], "#") {
				replyto = event.Arguments[0]
			} else {
				replyto = event.Nick
			}

			a := strings.Split(event.Message(), " ")

			if len(a) > 1 {
				query = a[1]
			} else {
				query = "empty"
			}

			answer := Help(query)

			if len(answer) > 1 {
				conn.Privmsg(replyto, answer)
			}

		}
	})
}
