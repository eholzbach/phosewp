package main

import (
	"fmt"
	"github.com/eholzbach/phosewp/config"
	"github.com/eholzbach/phosewp/events"
	"github.com/eholzbach/phosewp/plugins"
	"github.com/thoj/go-ircevent"
)

func main() {

	conf := config.Config()
	fmt.Println(conf.Channels)
	fmt.Printf("connecting bot...\n")

	conn := irc.IRC(conf.Handle, conf.Handle)
	conn.UseTLS = conf.Ssl

	err := conn.Connect(conf.Network)

	if err != nil {
		fmt.Printf("Error creating connection: %s\n", err)
		return
	}

	go events.Global(conn, conf)
	go plugins.Plugins(conn, conf)

	conn.Loop()
}
