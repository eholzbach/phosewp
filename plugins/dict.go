package plugins

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/eholzbach/phosewp/config"
	irc "github.com/thoj/go-ircevent"
)

type webster []struct {
	Shortdef []string `json:"shortdef"`
}

//  Dict queries the Merriam-Webster Collegiate Dictionary
func dict(conn *irc.Connection, r string, event *irc.Event, conf config.Vars) {
	if len(conf.Dictionary) <= 1 {
		log.Println("dictionary api key not found")
		return
	}

	i := strings.Split(event.Message(), " ")

	if len(i) <= 1 {
		log.Println("no parameter found")
		return
	}

	word := strings.Replace(i[1], " ", "", -1)

	if len(word) <= 1 {
		log.Println("no parameter found")
	}

	a, err := getURL(fmt.Sprintf("https://www.dictionaryapi.com/api/v3/references/collegiate/json/%s?key=%s", word, conf.Dictionary))

	if err != nil {
		log.Println(err)
		return
	}

	defer a.Body.Close()

	var b webster

	if err := json.NewDecoder(a.Body).Decode(&b); err != nil {
		log.Println(err)
		return
	}

	if len(b) != 0 {
		conn.Privmsg(r, b[0].Shortdef[0])
	}
}
