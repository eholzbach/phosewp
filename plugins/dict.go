package plugins

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"golang.org/x/net/dict"
	"strings"
	"time"
)

//need to strip newlines within defs, server supplied format is stupid 
func Dictword(conn *irc.Connection, resp string, dn string, dword string) {
	c, err := dict.Dial("tcp", "dict.org:2628")
	if err != nil {
		fmt.Println(err)
		return
	}

	d, err := c.Define(dn, dword)
	c.Close()

	if len(d) <= 0 {
		conn.Privmsg(resp, "not found")
	} else {
		s := strings.Split(string(d[0].Text), "\n")
		i := 0
		for _, v := range s {
			t := strings.TrimSpace(v)
			conn.Privmsgf(resp, t)
			time.Sleep(300 * time.Millisecond)
			i += 1
			if i == 4 || i == 8 {
				time.Sleep(1 * time.Second)
			}
			if i == 10 {
				break
			}
		}
	}
}
