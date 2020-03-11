package plugins

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
)

// Url resolves website titles
func Url(conn *irc.Connection, r string, event *irc.Event) {
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
		fmt.Println(err)
	}

	return traverse(doc)
}

// isTitleElement tests for the title tag
func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func traverse(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		return n.FirstChild.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.NextSibling == nil {
			break
		}

		result, ok := traverse(c)
		if ok {
			return result, ok
		}
	}

	return "", false
}
