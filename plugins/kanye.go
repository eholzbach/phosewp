package plugins

import (
	"encoding/json"
	"fmt"
	"github.com/thoj/go-ircevent"
	"net/http"
)

type KanyeMagic struct {
	Quote string `json:"quote"`
}

// Kanye provides Kanye
func Kanye(conn *irc.Connection, r string, event *irc.Event) {

	a, err := http.Get("https://api.kanye.rest")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer a.Body.Close()
	var con KanyeMagic
	json.NewDecoder(a.Body).Decode(&con)

	conn.Privmsg(r, con.Quote)
}
