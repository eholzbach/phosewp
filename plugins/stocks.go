/*
  Queries markitondemand for last close price. Typical for unauthenticated api's.
*/

package plugins

import (
	"encoding/json"
	"fmt"
	"github.com/thoj/go-ircevent"
	"net/http"
	"strconv"
	"strings"
)

type Stock struct {
	Message          string  `json:"Message"`
	Status           string  `json:"Status"`
	Name             string  `json:"Name"`
	Symbol           string  `json:"Symbol"`
	Lastprice        float64 `json:"LastPrice"`
	Change           float64 `json:"Change"`
	Changepercent    float64 `json:"ChangePercent"`
	Timestamp        string  `json:"Timestamp"`
	Msdate           float64 `json:"MSDate"`
	Marketcap        int64   `json:"MarketCap"`
	Volume           int     `json:"Volume"`
	Changeytd        float64 `json:"ChangeYTD"`
	Changepercentytd float64 `json:"ChangePercentYTD"`
	High             float64 `json:"High"`
	Low              float64 `json:"Low"`
	Open             float64 `json:"Open"`
}

func Stocks(conn *irc.Connection) {
	conn.AddCallback("PRIVMSG", func(event *irc.Event) {
		if strings.HasPrefix(event.Message(), "!stock ") == true {

			var replyto string

			if strings.HasPrefix(event.Arguments[0], "#") {
				replyto = event.Arguments[0]
			} else {
				replyto = event.Nick
			}

			query := strings.TrimPrefix(event.Message(), "!stock ")

			endpoint := fmt.Sprintf("http://dev.markitondemand.com/Api/v2/Quote/json?symbol=%s", query)
			r, err := http.Get(endpoint)

			if err != nil {
				fmt.Println(err)
				return
			}

			defer r.Body.Close()
			var con Stock
			json.NewDecoder(r.Body).Decode(&con)

			if len(con.Status) <= 0 {
				conn.Privmsg(replyto, "not found")
				return
			}
			s := fmt.Sprintf("%s: $%s, change $%s %s, high $%s, low $%s, %s", con.Name, strconv.FormatFloat(con.Lastprice, 'f', 2, 64), strconv.FormatFloat(con.Change, 'f', 2, 64), strconv.FormatFloat(con.Changepercent, 'f', 2, 64), strconv.FormatFloat(con.High, 'f', 2, 64), strconv.FormatFloat(con.Low, 'f', 2, 64), con.Timestamp)
			conn.Privmsg(replyto, s)
		}

	})
}
