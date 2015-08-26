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

// ed structs
type Edsearch struct {
	Query struct {
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

type Edresult struct {
	Query struct {
		Pages map[string]struct {
			Pageid    int    `json:"pageid"`
			Ns        int    `json:"ns"`
			Title     string `json:"title"`
			Revisions []struct {
				Star string `json:"*"`
			} `json:"revisions"`
		} `json:"pages"`
	} `json:"query"`
}

func Dramatica(conn *irc.Connection, resp string, dword string) {

	w, err := mwclient.New("https://encyclopediadramatica.se/api.php", "golang_api")
	if err != nil {
		fmt.Println(err)
		return
	}

	s := params.Values {
		"action":       "query",
		"list":         "search",
		"continue":     "",
		"srsearch":     dword,
	}

	sresp, err := w.GetRaw(s)
	u := &Edsearch{}
	if err := json.Unmarshal([]byte(sresp), &u); err != nil {
		fmt.Println(err)
		return
	}

	if len(u.Query.Search) == 0 {
		conn.Privmsg(resp, "not found")
		return
	}

	t := u.Query.Search[0].Title

	a := params.Values {
		"action":       "query",
		"format":       "json",
		"prop":         "revisions",
		"rvprop":       "content",
		"rvsection":    "0",
		"titles":       t,
	}

	rresp, err := w.GetRaw(a)
	b := &Edresult{}
	if err := json.Unmarshal([]byte(rresp), &b); err != nil {
		fmt.Println(err)
		return
	}

	for _, p := range b.Query.Pages {
		x := strings.Split(string(p.Revisions[0].Star), "\n")
		for _, line := range x {
			if strings.HasPrefix(line, "[[Image") ||
				strings.HasPrefix(line, "[[File") ||
				strings.HasPrefix(line, "{{") ||
				len(line) <= 0 {
			} else {
				l := strings.Replace(line, "'", "", -1)
				o := strings.Replace(l, "[", "", -1)
				z := strings.Replace(o, "]", "", -1)

				if len(z) >= 430 {
					a := []rune(z)
					conn.Privmsgf(resp, string(a[:430]))
					conn.Privmsgf(resp, string(a[430:]))
				} else {
					conn.Privmsgf(resp, z)
				}

			}
		}
	}
}

