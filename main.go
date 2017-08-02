package main

import (
	"fmt"
	"github.com/eholzbach/phosewp/config"
	"github.com/eholzbach/phosewp/events"
	"github.com/eholzbach/phosewp/plugins"
	"github.com/thoj/go-ircevent"
)

func main() {

	network, ssl, handle, channels, wuapi := config.Config()

	fmt.Printf("connecting bot...\n")

	conn := irc.IRC(handle, handle)
	conn.UseTLS = ssl

	err := conn.Connect(network)

	if err != nil {
		fmt.Println(err)
		return
	}

	events.Global(conn, channels, handle, wuapi)

	plugins.Dramatica(conn)
	plugins.Quote(conn)
	plugins.Stocks(conn)
	plugins.Urban(conn)
	plugins.Url(conn)
	plugins.Weather(conn, wuapi)
	plugins.Wiki(conn)
	plugins.Word(conn)

	conn.Loop()
}
