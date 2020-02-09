package events

import (
	"fmt"
	"github.com/eholzbach/phosewp/config"
	"github.com/thoj/go-ircevent"
	"strings"
)

func Global(conn *irc.Connection, conf *config.ConfigVars) {

	conn.AddCallback("001", func(event *irc.Event) {
		for _, channel := range conf.Channels {
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
		if event.Arguments[1] == conf.Handle {
			conn.Join(event.Arguments[0])
			s := fmt.Sprint("eat a bag of dicks, ", event.Nick, ".")
			conn.Privmsg(event.Arguments[0], s)
			fmt.Println("KICKED " + strings.Join(event.Arguments, " "))
		}
	})
}
