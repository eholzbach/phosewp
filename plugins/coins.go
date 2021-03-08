package plugins

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/eholzbach/phosewp/config"
	irc "github.com/thoj/go-ircevent"
)

type data struct {
	Data map[string]*json.RawMessage
}

type cdata struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Quote  struct {
		USD struct {
			Price            float64 `json:"price"`
			Volume24H        float64 `json:"volume_24h"`
			PercentChange1H  float64 `json:"percent_change_1h"`
			PercentChange24H float64 `json:"percent_change_24h"`
			PercentChange7D  float64 `json:"percent_change_7d"`
			PercentChange30D float64 `json:"percent_change_30d"`
		} `json:"USD"`
	} `json:"quote"`
}

func coins(conn *irc.Connection, r string, event *irc.Event, conf *config.ConfigVars) {
	if len(conf.Coinmarketcap) <= 1 {
		log.Println("coinmarketcap api key not found")
		return
	}

	coin := strings.Trim(strings.TrimPrefix(event.Message(), "!coin"), " ")
	if len(coin) == 0 {
		return
	}

	url := fmt.Sprintf("https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?symbol=%s", coin)

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", conf.Coinmarketcap)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	var d data

	json.NewDecoder(resp.Body).Decode(&d)

	for _, v := range d.Data {
		var c cdata
		if err := json.Unmarshal(*v, &c); err != nil {
			log.Println(err)
			return
		}
		a := fmt.Sprintf("%s: %s  -  Price: $%.2f  -  Change: 1h %.2f%%,  1d %.2f%%,  30d, %.2f%%", c.Symbol, c.Name, c.Quote.USD.Price, c.Quote.USD.PercentChange1H, c.Quote.USD.PercentChange7D, c.Quote.USD.PercentChange30D)
		conn.Privmsg(r, a)
	}

}
