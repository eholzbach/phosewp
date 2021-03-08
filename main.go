package main

import (
	"log"
	"os"

	"github.com/eholzbach/phosewp/config"
	"github.com/eholzbach/phosewp/events"
	"github.com/eholzbach/phosewp/plugins"
	irc "github.com/thoj/go-ircevent"
)

func main() {
	// read configuration
	log.Print("reading configuration...")
	conf, err := config.Config()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// set up connection parameters
	conn := irc.IRC(conf.Handle, conf.Handle)
	conn.UseTLS = conf.Ssl
	conn.UseSASL = conf.Sasl
	conn.SASLLogin = conf.Handle
	conn.SASLPassword = conf.Password

	log.Println("connecting bot...")

	// connnect to server
	err = conn.Connect(conf.Network)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// start up event handlers
	go events.Global(conn, conf)
	go plugins.Plugins(conn, conf)

	conn.Loop()
}
