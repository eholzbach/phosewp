/*
  In lulz we trust.
  Mediawiki is a dumpster fire of early 2000's php shitware.
*/

package plugins

import (
	"cgt.name/pkg/go-mwclient"
	"cgt.name/pkg/go-mwclient/params"
	"encoding/json"
	"fmt"
	"github.com/kennygrant/sanitize"
	"github.com/thoj/go-ircevent"
	"strings"
	"time"
)

type Edsearch struct {
	Query struct {
		Search []struct {
			Ns        int       `json:"ns"`
			Title     string    `json:"title"`
			Snippet   string    `json:"snippet"`
			Size      int       `json:"size"`
			Wordcount int       `json:"wordcount"`
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

func Dramatica(conn *irc.Connection, r string, event *irc.Event) {

	w, err := mwclient.New("https://encyclopediadramatica.rs/api.php", "dongs")
	if err != nil {
		fmt.Println(err)
		return
	}

	s := params.Values{
		"action":   "query",
		"list":     "search",
		"continue": "",
		"srsearch": strings.TrimPrefix(event.Message(), "!drama "),
	}

	sresp, err := w.GetRaw(s)
	u := &Edsearch{}
	if err := json.Unmarshal([]byte(sresp), &u); err != nil {
		fmt.Println(err)
		return
	}

	if len(u.Query.Search) == 0 {
		conn.Privmsg(r, "not found")
		return
	}

	t := u.Query.Search[0].Title

	a := params.Values{
		"action":    "query",
		"format":    "json",
		"prop":      "revisions",
		"rvprop":    "content",
		"rvsection": "0",
		"titles":    t,
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
			var str []string
			var final []string

			//cut out mediawiki links
			a := strings.Split(line, " [[")
			for _, v := range a {
				if strings.Contains(v, "|") {
					s := strings.Split(v, "|")
					str = append(str, s[1])
				} else {
					str = append(str, v)
				}
			}

			for _, v := range str {
				v = strings.Replace(v, "]", "", -1)
				v = strings.Replace(v, "'''", "", -1)
				v = sanitize.HTML(v)
				final = append(final, v)
			}

			z := strings.Join(final, " ")
			q := strings.Split(z, "\n")
			lcount := 0
			tcount := 0
			for _, line := range q {
				if len(line) <= 0 {
				} else {
					if len(line) >= 430 {
						a := []rune(line)
						conn.Privmsg(r, string(a[:430]))
						time.Sleep(300 * time.Millisecond)
						conn.Privmsg(r, string(a[430:]))
						lcount += 2
					} else {
						conn.Privmsg(r, line)
						lcount += 1
					}
					if lcount == 4 {
						time.Sleep(2 * time.Second)
						lcount = 0
					}
					if tcount >= 40 {
						break
					}
				}
			}
		}
	}
}
