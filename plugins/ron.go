package plugins

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"io/ioutil"
	"net/http"
	"strings"
)

// Ron would like you to shut up and idle
func Ron(conn *irc.Connection, r string, event *irc.Event) {
	a, err := http.Get("https://ron-swanson-quotes.herokuapp.com/v2/quotes")

	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := ioutil.ReadAll(a.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	q := strings.TrimLeft(string(b), "[\"")
	q = strings.TrimRight(q, "\"]")

	conn.Privmsg(r, q)
}
