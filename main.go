package main

import (
	"fmt"
	"github.com/eholzbach/phosewp/config"
	"github.com/eholzbach/phosewp/events"
	"github.com/eholzbach/phosewp/plugins"
	"github.com/thoj/go-ircevent"
)

func main() {
	// read configuration
	conf := config.Config()
	fmt.Printf("connecting bot...\n")

	// set up connection parameters
	conn := irc.IRC(conf.Handle, conf.Handle)
	conn.UseTLS = conf.Ssl
	conn.UseSASL = conf.Sasl
	conn.SASLLogin = conf.Handle
	conn.SASLPassword = conf.Password

	// connnect to server
	err := conn.Connect(conf.Network)
	if err != nil {
		fmt.Printf("Error creating connection: %s\n", err)
		return
	}

	// start up event handlers
	go events.Global(conn, conf)
	go plugins.Plugins(conn, conf)

	conn.Loop()
}
