// FoaaS

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

type Foff struct {
	Message  string `json:"message"`
	Subtitle string `json:"subtitle"`
}

type operators []struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Fields []struct {
		Name  string `json:"name"`
		Field string `json:"field"`
	} `json:"fields"`
}

func FoaaS(conn *irc.Connection) {
	conn.AddCallback("PRIVMSG", func(event *irc.Event) {
		if strings.HasPrefix(event.Message(), "!fu") == true {

			var replyto string
			var reply string

			if strings.HasPrefix(event.Arguments[0], "#") {
				replyto = event.Arguments[0]
			} else {
				replyto = event.Nick
			}

			name := strings.Trim(strings.TrimPrefix(event.Message(), "!fu"), " ")
			if len(name) == 0 {
				name = "0"
			}

			getOps, err := http.Get("https://www.foaas.com/operations")

			if err != nil {
				fmt.Println(err)
				return
			}

			defer getOps.Body.Close()
			con := operators{}
			json.NewDecoder(getOps.Body).Decode(&con)

			endpoint := randFoaas(con, name)
			reply = getFoaas(endpoint, name)

			conn.Privmsg(replyto, reply)
		}
	})
}

func getRand(count int) int {
	rand.Seed(time.Now().Unix())
	a := rand.Intn(count-1) + 1
	return a
}

func getFoaas(endpoint string, name string) string {
	var b string
	var url string

	a := strings.Split(endpoint, "/")
	if len(a) == 4 {
		url = fmt.Sprintf("https://www.foaas.com/%s/%s/%s", a[1], name, name)
	} else {
		url = fmt.Sprintf("https://www.foaas.com/%s/%s", a[1], name)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return "foaas fucked off"
	}

	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "foaas fucked off"
	}

	defer resp.Body.Close()
	var fo Foff
	json.NewDecoder(resp.Body).Decode(&fo)

	if len(fo.Message) <= 0 {
		fmt.Println(fo.Message)
		b = "short"
	} else {
		b = fo.Message
	}
	return b
}

func randFoaas(con operators, name string) string {
	var a int
	var b []string
	count := 0

	for {
		if count > 30 {
			fmt.Println("failed to FoaaS")
			break
		}
		a = getRand(len(con))
		b = strings.Split(con[a].URL, "/")
		if len(b) == 2 || len(b) == 5 {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		if strings.HasPrefix(b[2], ":from") && name == "0" {
			break
		}
		if strings.HasPrefix(b[2], ":name") && name != "0" {
			break
		}
		time.Sleep(500 * time.Millisecond)
		count += 1
		continue
	}

	return con[a].URL
}
