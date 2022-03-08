package plugins

import (
	"io/ioutil"
	"log"
	"strings"

	irc "github.com/thoj/go-ircevent"
)

// Ron would like you to shut up and idle
func ron(conn *irc.Connection, r string, event *irc.Event) {
	a, err := getURL("https://ron-swanson-quotes.herokuapp.com/v2/quotes")

	if err != nil {
		log.Println(err)
		return
	}

	b, err := ioutil.ReadAll(a.Body)

	if err != nil {
		log.Println(err)
		return
	}

	q := strings.TrimLeft(string(b), "[\"")
	q = strings.TrimRight(q, "\"]")

	conn.Privmsg(r, q)
}
