/*
  shitpost title parser
*/

package plugins

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
)

func Url(conn *irc.Connection) {
	conn.AddCallback("PRIVMSG", func(event *irc.Event) {
		if strings.Contains(event.Message(), "http://") || strings.Contains(event.Message(), "https://") {

			var replyto string

			if strings.HasPrefix(event.Arguments[0], "#") {
				replyto = event.Arguments[0]
			} else {
				replyto = event.Nick
			}

			a := strings.Split(event.Message(), " ")
			for _, b := range a {
				if strings.HasPrefix(b, "http://") || strings.HasPrefix(b, "https://") {
					response, err := http.Get(b)
					if err != nil {
						fmt.Println(err)
						return
					}
					defer response.Body.Close()
					if title, ok := GetTitle(response.Body); ok {
						conn.Privmsg(replyto, title)
					}

				}
			}
		}
	})
}

func GetTitle(r io.Reader) (string, bool) {
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Println(err)
	}

	return traverse(doc)
}

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func traverse(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		return n.FirstChild.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := traverse(c)
		if ok {
			return result, ok
		}
	}

	return "", false
}
