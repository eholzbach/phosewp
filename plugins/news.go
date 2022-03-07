package plugins

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/eholzbach/phosewp/config"
	irc "github.com/thoj/go-ircevent"
)

type fakeNews struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"source"`
		Author      string    `json:"author"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		URL         string    `json:"url"`
		URLToImage  string    `json:"urlToImage"`
		PublishedAt time.Time `json:"publishedAt"`
	} `json:"articles"`
}

func breitButt(event *irc.Event, token string) string {

	url := "https://newsapi.org/v2"

	// garbage sources only for maximum lolwat
	sources := []string{
		"breitbart-news",
		"the-american-conservative",
	}

	msg := strings.Split(event.Message(), " ")
	var endpoint string

	// randomly select a source
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, i := range r.Perm(len(sources)) {
		if len(msg) > 1 {
			query := strings.Join(msg[1:], "%20")
			endpoint = fmt.Sprintf("%s/top-headlines?apiKey=%s&sources=%s&pageSize=20&q=%s", url, token, sources[i], query)
		} else {
			endpoint = fmt.Sprintf("%s/top-headlines?apiKey=%s&sources=%s&pageSize=20", url, token, sources[i])
		}

		a, err := http.Get(endpoint)

		if err != nil {
			return "failed"
		}

		defer a.Body.Close()
		var con fakeNews
		json.NewDecoder(a.Body).Decode(&con)

		if con.TotalResults != 0 {
			a := r.Intn(con.TotalResults - 1 + 1)
			return con.Articles[a].Title
		}
	}

	// try everything
	all := strings.Join(sources, ",")
	if len(msg) > 1 {
		query := strings.Join(msg[1:], "%20")
		endpoint = fmt.Sprintf("%s/everything?apiKey=%s&sources=%s&pageSize=20&q=%s", url, token, all, query)
	} else {
		endpoint = fmt.Sprintf("%s/everything?apiKey=%s&sources=%s&pageSize=20", url, token, all)
	}

	a, err := http.Get(endpoint)

	if err != nil {
		return "failed"
	}

	defer a.Body.Close()
	var con fakeNews
	json.NewDecoder(a.Body).Decode(&con)

	if con.TotalResults != 0 {
		a := r.Intn(con.TotalResults - 1 + 1)
		return con.Articles[a].Title
	}

	return "no articles found"
}

// News provides garbage titles from garbage sources
func news(conn *irc.Connection, r string, event *irc.Event, conf config.Vars) {
	if len(conf.Newsapi) <= 1 {
		log.Println("newsapi key not found")
		return
	}

	line := breitButt(event, conf.Newsapi)
	conn.Privmsg(r, line)
}
