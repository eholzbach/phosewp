package main

import (
	"flag"
	"fmt"
	"github.com/thoj/go-ircevent"
	"phosoup/plugins"
	"strings"
)

func main() {

	// get flags and dereference pointers
	c := flag.String("c", "", "channel: channel to join, omit the #")
	h := flag.String("h", "", "handle: nick for bot to use")
	n := flag.String("n", "", "network: fqdn of server")

	flag.Parse()

	channel := fmt.Sprint("#",*c)
	network := *n
	handle := *h

	if len(channel) <= 1 {
		fmt.Println("no channel")
		return
	}
	if len(network) <= 0 {
		fmt.Println("no network")
		return
	}
	if len(handle) <= 0 {
		fmt.Println("no handle")
		return
	}

	fmt.Printf("connecting bot...\n")

	conn := irc.IRC(handle, handle)
	netw := fmt.Sprint(network,":6667")

	err := conn.Connect(netw)
	if err != nil {
		fmt.Println(err)
		return
	}

	DoThings(conn, channel, handle)
	conn.Loop()
}

func DoThings(conn *irc.Connection, channel string, handle string) {
	conn.AddCallback("001", func (event *irc.Event) {
		conn.Join(channel)
	})

	conn.AddCallback("INVITE", func(event *irc.Event) {
		if len(event.Arguments) != 2 {
			return
		}
		conn.Join(event.Arguments[1])
	})

	conn.AddCallback("KICK", func (event *irc.Event) {
		if event.Arguments[1] == handle {
			conn.Join(event.Arguments[0])
			s := fmt.Sprint("eat a bag of dicks, ", event.Nick, ".")
			conn.Privmsg(channel, s)
			fmt.Println("KICKED " + strings.Join(event.Arguments, " "))
		}
	})

	conn.AddCallback("PRIVMSG", func (event *irc.Event) {
		message := event.Message()
		if strings.HasPrefix(message, "!") == true {
			cmd := strings.Split(message, " ")
			var resp string
			if strings.HasPrefix(event.Arguments[0], "#") {
				resp = event.Arguments[0]
			} else {
				resp = event.Nick
			}
			switch {
			case strings.Contains(cmd[0], "acronym"):
				dn := "vera"
				dword := cmd[1]
				plugins.Dictword(conn, resp, dn, dword)
			case strings.Contains(cmd[0], "astronomy"):
				dword := strings.TrimPrefix(message, "!astronomy ")
				query := "astronomy"
				plugins.WeatherUnderground(conn, dword, query, resp)
			case strings.Contains(cmd[0], "dict"):
				dn := "wn"
				dword := cmd[1]
				plugins.Dictword(conn, resp, dn, dword)
			case strings.Contains(cmd[0], "drama"):
				dword := strings.TrimPrefix(message, "!drama ")
				plugins.Dramatica(conn, resp, dword)
			case strings.Contains(cmd[0], "help"):
				conn.Privmsg(event.Nick, "commands are: acronym, astronomy, drama, dict, tide, urban, weather, wiki")
			case strings.Contains(cmd[0], "tide"):
				dword := strings.TrimPrefix(message, "!tide ")
				query := "tide"
				plugins.WeatherUnderground(conn, dword, query, resp)
			case strings.Contains(cmd[0], "trivia"):
				plugins.Trivia(conn, resp)
			case strings.Contains(cmd[0], "urban"):
				dword := strings.TrimPrefix(message, "!urban ")
				plugins.Urban(conn, resp, dword)
			case strings.Contains(cmd[0], "weather"):
				dword := strings.TrimPrefix(message, "!weather ")
				query := "weather"
				plugins.WeatherUnderground(conn, dword, query, resp)
			case strings.Contains(cmd[0], "wiki"):
				dword := strings.TrimPrefix(message, "!wiki ")
				plugins.Wiki(conn, resp, dword)
			default:
			}
			fmt.Println(cmd[0])
		}

		if strings.Contains(message, "http://") || strings.Contains(message, "https://") {
			plugins.Urlresolve(conn, channel, message)
		}
	})
}
