package plugins

import (
	"bytes"
	"fmt"
	"github.com/thoj/go-ircevent"
	"net/http"
	"strings"
)

func Urlresolve(conn *irc.Connection, resp string, message string) {
	a := strings.Split(message, " ")
	for _, b := range a {
		if strings.HasPrefix(b, "http://") || strings.HasPrefix(b, "https://") {
			response, err := http.Get(b)
			if err != nil {
				fmt.Println(err)
				return
			} else {
				defer response.Body.Close()
				c := http.MaxBytesReader(nil, response.Body, 10000)
				d := new(bytes.Buffer)
				d.ReadFrom(c)
				e := d.String()
				f := strings.Split(string(e), "<")
				for _, g := range f {
					if strings.Contains(g, "title>") ||
					   strings.Contains(g, "Title>") ||
					   strings.Contains(g, "TITLE>") {
						h := strings.TrimSpace(g)
						i := strings.Split(h, ">")
						j := i[1]
						k := strings.Replace(j, "\n", "", -1)
						l := strings.TrimSpace(k)
						conn.Privmsgf(resp, l)
					}
				}
			}
		}
	}
}

