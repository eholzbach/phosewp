package plugins

import (
	"encoding/json"
	"fmt"
	"github.com/thoj/go-ircevent"
	"net/http"
)

type Dump struct {
	AppearedAt string   `json:"appeared_at"`
	CreatedAt  string   `json:"created_at"`
	QuoteID    string   `json:"quote_id"`
	Tags       []string `json:"tags"`
	UpdatedAt  string   `json:"updated_at"`
	Value      string   `json:"value"`
	Links      struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Embedded struct {
		Author []struct {
			CreatedAt string      `json:"created_at"`
			Bio       interface{} `json:"bio"`
			AuthorID  string      `json:"author_id"`
			Name      string      `json:"name"`
			Slug      string      `json:"slug"`
			UpdatedAt string      `json:"updated_at"`
		} `json:"author"`
		Source []struct {
			CreatedAt     string      `json:"created_at"`
			Filename      interface{} `json:"filename"`
			QuoteSourceID string      `json:"quote_source_id"`
			Remarks       interface{} `json:"remarks"`
			UpdatedAt     string      `json:"updated_at"`
			URL           string      `json:"url"`
		} `json:"source"`
	} `json:"_embedded"`
}

// Tronald pukes quotes from a garbage person
func Tronald(conn *irc.Connection, r string, event *irc.Event) {

	a, err := http.Get("https://api.tronalddump.io/random/quote")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer a.Body.Close()
	var con Dump
	json.NewDecoder(a.Body).Decode(&con)

	conn.Privmsg(r, con.Value)
}
