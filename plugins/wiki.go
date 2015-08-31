package plugins

import (
	"cgt.name/pkg/go-mwclient"
	"cgt.name/pkg/go-mwclient/params"
	"encoding/json"
	"fmt"
	"github.com/thoj/go-ircevent"
	"strings"
	"time"
)

// wikipedia json structs
type Wresult struct {
	Batchcomplete string
	Query         struct {
		Normalized []struct {
			From string
			To   string
		}
		Pages map[string]struct {
			Pageid  int
			Ns      int
			Title   string
			Extract string
		}
	}
}

type Qsearch struct {
	Batchcomplete string `json:"batchcomplete"`
	Continue struct {
		Sroffset int `json:"sroffset"`
		Continue string `json:"continue"`
	} `json:"continue"`
	Query struct {
		Searchinfo struct {
			Totalhits int `json:"totalhits"`
		} `json:"searchinfo"`
		Search []struct {
			Ns int `json:"ns"`
			Title string `json:"title"`
			Snippet string `json:"snippet"`
			Size int `json:"size"`
			Wordcount int `json:"wordcount"`
			Timestamp time.Time `json:"timestamp"`
		} `json:"search"`
	} `json:"query"`
}

func Wiki(conn *irc.Connection, resp string, dword string) {
	w, err := mwclient.New("http://en.wikipedia.org/w/api.php", "golang_api")
	if err != nil {
		fmt.Println(err)
		return
	}

	s := params.Values{
		"action":       "query",
		"list":         "search",
		"continue":     "",
		"srsearch":     dword,
	}

	sresp, err := w.GetRaw(s)
	u := &Qsearch{}
	if err := json.Unmarshal([]byte(sresp), &u); err != nil {
		fmt.Println(err)
		return
	}

	if u.Query.Searchinfo.Totalhits == 0 {
		conn.Privmsg(resp, "not found")
		return
	}

	t := u.Query.Search[0].Title
	q := params.Values{
		"action":       "query",
		"format":       "json",
		"redirects":    "",
		"prop":         "extracts",
		"exintro":      "",
		"explaintext":  "",
		"titles":       t,
	}

	v, err := w.GetRaw(q)
	if err != nil {
		fmt.Println(err)
		return
	}

	r := &Wresult{}
	if err := json.Unmarshal([]byte(v), &r); err != nil {
		fmt.Println(err)
		return
	}

	for _, p := range r.Query.Pages {
		m := strings.Split(p.Extract, "\n")
		conn.Privmsgf(resp, m[0])
	}
}
