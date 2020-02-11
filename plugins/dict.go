package plugins

import (
	"encoding/json"
	"fmt"
	"github.com/eholzbach/phosewp/config"
	"github.com/thoj/go-ircevent"
	"net/http"
	"strings"
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
func Dict(conn *irc.Connection, r string, event *irc.Event, conf *config.ConfigVars) {
	if len(conf.Dictionary) <= 1 {
		fmt.Println("dictionary api key not found")
		return
	}

	word := strings.Split(event.Message(), " ")
	url := fmt.Sprintf("https://www.dictionaryapi.com/api/v3/references/collegiate/json/%s?key=%s", word[1], conf.Dictionary)

	a, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer a.Body.Close()
	var b webster
	err = json.NewDecoder(a.Body).Decode(&b)
	if err != nil {
		return
	}

	if len(b) != 0 {
		conn.Privmsg(r, b[0].Shortdef[0])
	}
}
