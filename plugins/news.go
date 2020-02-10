package plugins

import (
	"encoding/json"
	"fmt"
	"github.com/thoj/go-ircevent"
	"math/rand"
	"net/http"
	"strings"
	"time"
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

	sources := []string{
		"breitbart-news",
		"fox-news",
		"national-review",
	}

	url := "https://newsapi.org/v2"

	// try top headlines from a random source
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	msg := strings.Split(event.Message(), " ")
	var endpoint string

	for _, i := range r.Perm(len(sources)) {
		if len(msg) > 1 {
			query := strings.Join(msg[1:], "%20")
			endpoint = fmt.Sprintf("%s/top-headlines?apiKey=%s&sources=%s&pageSize=1&q=%s", url, token, sources[i], query)
		} else {
			endpoint = fmt.Sprintf("%s/top-headlines?apiKey=%s&sources=%s&pageSize=1", url, token, sources[i])
		}

		a, err := http.Get(endpoint)

		if err != nil {
			return "failed"
		}

		defer a.Body.Close()
		var con fakeNews
		json.NewDecoder(a.Body).Decode(&con)

		if con.TotalResults != 0 {
			return con.Articles[0].Title
		}
	}

	// try everything, limited to a month by api
	all := strings.Join(sources, ",")
	if len(msg) > 1 {
		query := strings.Join(msg[1:], "%20")
		endpoint = fmt.Sprintf("%s/everything?apiKey=%s&sources=%s&pageSize=1&q=%s", url, token, all, query)
	} else {
		endpoint = fmt.Sprintf("%s/everything?apiKey=%s&sources=%s&pageSize=1", url, token, all)
	}

	a, err := http.Get(endpoint)

	if err != nil {
		return "failed"
	}

	defer a.Body.Close()
	var con fakeNews
	json.NewDecoder(a.Body).Decode(&con)

	if con.TotalResults != 0 {
		return con.Articles[0].Title
	}

	return "no articles found"
}

func News(conn *irc.Connection, r string, event *irc.Event, token string) {

	if len(token) <= 1 {
		fmt.Println("newsapi key not found")
		return
	}

	line := breitButt(event, token)
	conn.Privmsg(r, line)
}
