package plugins

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/eholzbach/phosewp/config"
	irc "github.com/thoj/go-ircevent"
)

type forecast []struct {
	WeatherText string `json:"WeatherText"`
	Temperature struct {
		Imperial struct {
			Value float64 `json:"Value"`
		} `json:"Imperial"`
	} `json:"Temperature"`
	RelativeHumidity int `json:"RelativeHumidity"`
	Wind             struct {
		Speed struct {
			Imperial struct {
				Value float64 `json:"Value"`
			} `json:"Imperial"`
		} `json:"Speed"`
	} `json:"Wind"`
}

type location []struct {
	Key string `json:"Key"`
}

// weather returns a forcast summary from Darksky
func weather(conn *irc.Connection, r string, event *irc.Event, conf config.Vars) {
	if len(conf.AccuWeather) <= 1 {
		log.Println("AccuWeather api key not found")
		return
	}

	zip := strings.Split(event.Message(), " ")

	if !validInput(zip) {
		conn.Privmsg(r, "weather only accepts 5 digit zip codes")
		return
	}

	var loc location

	resp, err := getURL(fmt.Sprintf("http://dataservice.accuweather.com/locations/v1/postalcodes/US/search?apikey=%s&q=%s", conf.AccuWeather, zip))

	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&loc)
	resp.Body.Close()

	if len(loc) == 0 {
		log.Println("location not found")
		return
	}

	var forecast forecast

	resp, err = getURL(fmt.Sprintf("http://dataservice.accuweather.com/currentconditions/v1/%s?apikey=%s&details=true", loc[0].Key, conf.AccuWeather))

	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&forecast)

	if len(forecast) <= 0 {
		return
	}

	b := fmt.Sprintf("%s, Wind %.0f mph, Humidity %d%%, Temperature %.0f°", forecast[0].WeatherText, forecast[0].Wind.Speed.Imperial.Value,
		forecast[0].RelativeHumidity, forecast[0].Temperature.Imperial.Value)

	conn.Privmsg(r, b)
}

// validInput validates the entry is a zip code
func validInput(a []string) bool {
	if len(a) != 2 {
		return false
	}

	i := 0
	for _, v := range a[1] {
		switch {
		case v >= '0' && v <= '9':
			i++
		}
	}

	return i == 5
}
