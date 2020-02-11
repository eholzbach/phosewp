package plugins

import (
	"encoding/json"
	"fmt"
	"github.com/thoj/go-ircevent"
	"net/http"
)

type bsPhrase struct {
	Phrase string `json:"phrase"`
}

// Bs corporate uullshit/buzzword generator
func Bs(conn *irc.Connection, r string, event *irc.Event) {
	a, err := http.Get("https://corporatebs-generator.sameerkumar.website/")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer a.Body.Close()
	var b bsPhrase
	err = json.NewDecoder(a.Body).Decode(&b)
	if err != nil {
		return
	}

	conn.Privmsg(r, b.Phrase)
}
