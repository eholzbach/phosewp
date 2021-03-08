package plugins

import (
	"io"
	"log"
	"net/http"
	"strings"

	irc "github.com/thoj/go-ircevent"
	"golang.org/x/net/html"
)

// url resolves website titles
func urlz(conn *irc.Connection, r string, event *irc.Event) {
	a := strings.Split(event.Message(), " ")
	for _, b := range a {
		if strings.HasPrefix(b, "http://") || strings.HasPrefix(b, "https://") {
			response, err := http.Get(b)
			if err != nil {
				log.Println(err)
				return
			}
			defer response.Body.Close()
			if title, ok := GetTitle(response.Body); ok {
				title = strings.Replace(title, "\n", "", -1)
				title = strings.TrimSpace(title)
				conn.Privmsg(r, title)
			}

		}
	}
}

// GetTitle reads the url and returns the title
func GetTitle(r io.Reader) (string, bool) {
	doc, err := html.Parse(r)
	if err != nil {
		log.Println(err)
		return "", false
	}

	return traverse(doc)
}

// isTitleElement tests for the title tag
func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func traverse(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		if n.FirstChild != nil {
			return n.FirstChild.Data, true
		}
		return "", false
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := traverse(c)
		if ok {
			return result, ok
		}
	}

	return "", false
}
