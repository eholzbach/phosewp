package plugins

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"golang.org/x/net/dict"
	"strings"
	"time"
)

//  Dict queries dict.org using RFC2229. The world before RESTful api's was awful.
func Dict(conn *irc.Connection, r string, event *irc.Event) {
	var db string
	var query string

	if strings.HasPrefix(event.Message(), "!dict ") {
		db = "wn" // word net
		query = strings.TrimPrefix(event.Message(), "!dict ")
	} else {
		db = "moby-thesaurus" // moby
		query = strings.TrimPrefix(event.Message(), "!acronym ")
	}

	c, err := dict.Dial("tcp", "dict.org:2628")
	if err != nil {
		fmt.Println(err)
		return
	}

	d, err := c.Define(db, query)
	c.Close()

	if len(d) <= 0 {
		conn.Privmsg(r, "not found")
	} else {
		s := strings.Split(string(d[0].Text), "\n")
		i := 0
		for _, v := range s {
			t := strings.TrimSpace(v)
			conn.Privmsg(r, t)
			time.Sleep(300 * time.Millisecond)
			i += 1
			if i == 4 || i == 8 {
				time.Sleep(1 * time.Second)
			}
			if i == 20 {
				break
			}
		}
	}
}
