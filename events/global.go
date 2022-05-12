// Package events runs global event routines
package events

import (
	"fmt"
	"strings"

	"github.com/eholzbach/phosewp/config"
	irc "github.com/thoj/go-ircevent"
)

// Global event watcher
func Global(conn *irc.Connection, conf config.Vars) {
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
			conn.Privmsg(event.Arguments[0], fmt.Sprint("eat a pinecone, ", event.Nick, "."))
			fmt.Println("KICKED " + strings.Join(event.Arguments, " "))
		}
	})
}
