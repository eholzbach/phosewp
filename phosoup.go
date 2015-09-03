package main

import (
	"flag"
	"fmt"
	"github.com/thoj/go-ircevent"
	"io/ioutil"
	"os"
	"os/user"
	"phosoup/plugins"
	"runtime"
	"strings"
)

func main() {

	var c *string
	var e string
	var h *string
	var n *string
	var w *string
	var channel string
	var handle string
	var network string
	var wuapi string

	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	home := user.HomeDir
	const conf string = "/.phosoup.conf"
	d := home + conf

	if _, err := os.Stat(d); os.IsNotExist(err) {
		if runtime.GOOS == "freebsd" {
			e = "/usr/local/etc/phosoup.conf"
		} else {
			e = "/etc/phosoup.conf"
		}
	} else {
		e = d
	}

	if _, err := os.Stat(e); err == nil {

		f, err := ioutil.ReadFile(e)

		if err != nil {
			panic(err)
		}

		g := strings.Split(string(f), "\n")

		for _, v := range g {
			t := strings.TrimSpace(v)
			if strings.HasPrefix(t, "channel ") {
				s := strings.Split(t, " ")
				channel = s[1]
			}
			if strings.HasPrefix(t, "handle ") {
				s := strings.Split(t, " ")
				handle = s[1]
			}
			if strings.HasPrefix(t, "network ") {
				s := strings.Split(t, " ")
				network = s[1]
			}
			if strings.HasPrefix(t, "wuapi ") {
				s := strings.Split(t, " ")
				wuapi = s[1]
			}
		}
	} else {
		// get flags and dereference pointers
		c = flag.String("c", "", "channel: channel to join")
		h = flag.String("h", "", "handle: nick for bot to use")
		n = flag.String("n", "", "network: fqdn of server")
		w = flag.String("w", "", "wuapi: api key for weather underground")

		flag.Parse()
		channel = *c
		network = *n
		handle = *h
		wuapi = *w
	}

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

	err = conn.Connect(netw)
	if err != nil {
		fmt.Println(err)
		return
	}

	DoThings(conn, channel, handle, wuapi)
	conn.Loop()
}

func DoThings(conn *irc.Connection, channel string, handle string, wuapi string) {
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
				plugins.WeatherUnderground(conn, dword, query, resp, wuapi)
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
				plugins.WeatherUnderground(conn, dword, query, resp, wuapi)
			case strings.Contains(cmd[0], "urban"):
				dword := strings.TrimPrefix(message, "!urban ")
				plugins.Urban(conn, resp, dword)
			case strings.Contains(cmd[0], "weather"):
				dword := strings.TrimPrefix(message, "!weather ")
				query := "weather"
				plugins.WeatherUnderground(conn, dword, query, resp, wuapi)
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
