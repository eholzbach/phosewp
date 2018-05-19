package main

import (
	"fmt"
	"github.com/eholzbach/phosewp/config"
	"github.com/eholzbach/phosewp/events"
	"github.com/eholzbach/phosewp/plugins"
	"github.com/thoj/go-ircevent"
)

func main() {

	network, ssl, handle, channels, wuapi, newsapi := config.Config()

	fmt.Printf("connecting bot...\n")

	conn := irc.IRC(handle, handle)
	conn.UseTLS = ssl

	err := conn.Connect(network)

	if err != nil {
		fmt.Println(err)
		return
	}

	go events.Global(conn, channels, handle)
	go plugins.Plugins(conn, channels, wuapi, newsapi)

	conn.Loop()
}
