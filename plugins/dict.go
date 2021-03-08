package plugins

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/eholzbach/phosewp/config"
	irc "github.com/thoj/go-ircevent"
)

type webster []struct {
	Meta struct {
		ID        string   `json:"id"`
		UUID      string   `json:"uuid"`
		Sort      string   `json:"sort"`
		Src       string   `json:"src"`
		Section   string   `json:"section"`
		Stems     []string `json:"stems"`
		Offensive bool     `json:"offensive"`
	} `json:"meta"`
	Hwi struct {
		Hw  string `json:"hw"`
		Prs []struct {
			Mw    string `json:"mw"`
			Sound struct {
				Audio string `json:"audio"`
				Ref   string `json:"ref"`
				Stat  string `json:"stat"`
			} `json:"sound"`
		} `json:"prs"`
	} `json:"hwi"`
	Fl  string `json:"fl"`
	Ins []struct {
		Il  string `json:"il"`
		If  string `json:"if"`
		Prs []struct {
			Mw    string `json:"mw"`
			Sound struct {
				Audio string `json:"audio"`
				Ref   string `json:"ref"`
				Stat  string `json:"stat"`
			} `json:"sound"`
		} `json:"prs"`
	} `json:"ins"`
	Def []struct {
		Sseq [][][]interface{} `json:"sseq"`
	} `json:"def"`
	Et       [][]string `json:"et"`
	Date     string     `json:"date"`
	Shortdef []string   `json:"shortdef"`
}

//  Dict queries the Merriam-Webster Collegiate Dictionary
func dict(conn *irc.Connection, r string, event *irc.Event, conf *config.ConfigVars) {
	if len(conf.Dictionary) <= 1 {
		log.Println("dictionary api key not found")
		return
	}

	i := strings.Split(event.Message(), " ")
	if len(i) <= 1 {
		log.Println("no parameter found")
		return
	}

	word := strings.Replace(i[1], " ", "", -1)
	if len(word) <= 1 {
		log.Println("no parameter found")
	}

	url := fmt.Sprintf("https://www.dictionaryapi.com/api/v3/references/collegiate/json/%s?key=%s", word, conf.Dictionary)

	a, err := http.Get(url)

	if err != nil {
		log.Println(err)
		return
	}

	defer a.Body.Close()
	var b webster
	err = json.NewDecoder(a.Body).Decode(&b)
	if err != nil {
		log.Println(err)
		return
	}

	if len(b) != 0 {
		conn.Privmsg(r, b[0].Shortdef[0])
	}
}
