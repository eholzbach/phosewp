package plugins

import (
	"encoding/xml"
	"fmt"
	"github.com/thoj/go-ircevent"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Item struct {
        XMLName xml.Name `xml:"item"`
        Question string `xml:"question"`
        A string `xml:"a"`
        B string `xml:"b"`
        C string `xml:"c"`
        D string `xml:"d"`
        Answer string `xml:"answer"`
        Explanation string `xml:"explanation"`
        Figure_flag string `xml:"figure_flag"`
        Type string `xml:"type"`
}
type Cissp struct {
        XMLName xml.Name `xml:"cissp"`
        Items []Item `xml:"item"`
}

func Trivia (conn *irc.Connection, resp string) {
	x, err := os.Open("plugindata/cissp.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer x.Close()
	d, err := ioutil.ReadAll(x)
	if err != nil {
		fmt.Println(err)
	}
	var q Cissp
	xml.Unmarshal(d, &q)

	max := (len(q.Items) + 1)
	qnum := random(0, max)

	conn.Privmsgf(resp, q.Items[qnum].Question)
	time.Sleep(5 * time.Second)
	conn.Privmsgf(resp, "A: %s\n", q.Items[qnum].A)
	conn.Privmsgf(resp, "B: %s\n", q.Items[qnum].B)
	conn.Privmsgf(resp, "C: %s\n", q.Items[qnum].C)
	conn.Privmsgf(resp, "D: %s\n", q.Items[qnum].D)
	i := 0

	conn.AddCallback("PRIVMSG", func (event *irc.Event) {
		message := event.Message()
		var resp string
		if strings.HasPrefix(event.Arguments[0], "#") {
			resp = event.Arguments[0]
		} else {
		resp = event.Nick
		}

		if strings.EqualFold(message, "A") || strings.EqualFold(message, "B") || strings.EqualFold(message, "C") || strings.EqualFold(message, "D") {
			if strings.EqualFold(message, q.Items[qnum].Answer) {
				conn.Privmsgf(resp, "Correct! Answer is %s", q.Items[qnum].Answer)
				conn.Privmsgf(resp, q.Items[qnum].Explanation)
			}
		}

		i += 1
		fmt.Println(i)
		if i == 3 {
			conn.Privmsgf(resp, "Fail. Answer was %s", q.Items[qnum].Answer)
			conn.Privmsgf(resp, q.Items[qnum].Explanation)
		}

	})
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}
