package plugins

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/eholzbach/phosewp/config"
	irc "github.com/thoj/go-ircevent"
)

type zipcodes struct {
	Data []struct {
		Zipcode   string `json:"zipcode"`
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"data"`
}

type forcast struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Currently struct {
		Time                 int     `json:"time"`
		Summary              string  `json:"summary"`
		Icon                 string  `json:"icon"`
		NearestStormDistance int     `json:"nearestStormDistance"`
		NearestStormBearing  int     `json:"nearestStormBearing"`
		PrecipIntensity      int     `json:"precipIntensity"`
		PrecipProbability    int     `json:"precipProbability"`
		Temperature          float64 `json:"temperature"`
		ApparentTemperature  float64 `json:"apparentTemperature"`
		DewPoint             float64 `json:"dewPoint"`
		Humidity             float64 `json:"humidity"`
		Pressure             float64 `json:"pressure"`
		WindSpeed            float64 `json:"windSpeed"`
		WindGust             float64 `json:"windGust"`
		WindBearing          int     `json:"windBearing"`
		CloudCover           float64 `json:"cloudCover"`
		UvIndex              int     `json:"uvIndex"`
		Visibility           float64 `json:"visibility"`
		Ozone                float64 `json:"ozone"`
	} `json:"currently"`
	Minutely struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time              int `json:"time"`
			PrecipIntensity   int `json:"precipIntensity"`
			PrecipProbability int `json:"precipProbability"`
		} `json:"data"`
	} `json:"minutely"`
	Hourly struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time                int     `json:"time"`
			Summary             string  `json:"summary"`
			Icon                string  `json:"icon"`
			PrecipIntensity     int     `json:"precipIntensity"`
			PrecipProbability   int     `json:"precipProbability"`
			Temperature         float64 `json:"temperature"`
			ApparentTemperature float64 `json:"apparentTemperature"`
			DewPoint            float64 `json:"dewPoint"`
			Humidity            float64 `json:"humidity"`
			Pressure            float64 `json:"pressure"`
			WindSpeed           float64 `json:"windSpeed"`
			WindGust            float64 `json:"windGust"`
			WindBearing         int     `json:"windBearing"`
			CloudCover          int     `json:"cloudCover"`
			UvIndex             int     `json:"uvIndex"`
			Visibility          float64 `json:"visibility"`
			Ozone               float64 `json:"ozone"`
			PrecipType          string  `json:"precipType,omitempty"`
			PrecipAccumulation  float64 `json:"precipAccumulation,omitempty"`
		} `json:"data"`
	} `json:"hourly"`
	Daily struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time                        int     `json:"time"`
			Summary                     string  `json:"summary"`
			Icon                        string  `json:"icon"`
			SunriseTime                 int     `json:"sunriseTime"`
			SunsetTime                  int     `json:"sunsetTime"`
			MoonPhase                   float64 `json:"moonPhase"`
			PrecipIntensity             float64 `json:"precipIntensity"`
			PrecipIntensityMax          float64 `json:"precipIntensityMax"`
			PrecipIntensityMaxTime      int     `json:"precipIntensityMaxTime"`
			PrecipProbability           float64 `json:"precipProbability"`
			PrecipType                  string  `json:"precipType"`
			TemperatureHigh             float64 `json:"temperatureHigh"`
			TemperatureHighTime         int     `json:"temperatureHighTime"`
			TemperatureLow              float64 `json:"temperatureLow"`
			TemperatureLowTime          int     `json:"temperatureLowTime"`
			ApparentTemperatureHigh     float64 `json:"apparentTemperatureHigh"`
			ApparentTemperatureHighTime int     `json:"apparentTemperatureHighTime"`
			ApparentTemperatureLow      float64 `json:"apparentTemperatureLow"`
			ApparentTemperatureLowTime  int     `json:"apparentTemperatureLowTime"`
			DewPoint                    float64 `json:"dewPoint"`
			Humidity                    float64 `json:"humidity"`
			Pressure                    float64 `json:"pressure"`
			WindSpeed                   float64 `json:"windSpeed"`
			WindGust                    float64 `json:"windGust"`
			WindGustTime                int     `json:"windGustTime"`
			WindBearing                 int     `json:"windBearing"`
			CloudCover                  float64 `json:"cloudCover"`
			UvIndex                     int     `json:"uvIndex"`
			UvIndexTime                 int     `json:"uvIndexTime"`
			Visibility                  float64 `json:"visibility"`
			Ozone                       float64 `json:"ozone"`
			TemperatureMin              float64 `json:"temperatureMin"`
			TemperatureMinTime          int     `json:"temperatureMinTime"`
			TemperatureMax              float64 `json:"temperatureMax"`
			TemperatureMaxTime          int     `json:"temperatureMaxTime"`
			ApparentTemperatureMin      float64 `json:"apparentTemperatureMin"`
			ApparentTemperatureMinTime  int     `json:"apparentTemperatureMinTime"`
			ApparentTemperatureMax      float64 `json:"apparentTemperatureMax"`
			ApparentTemperatureMaxTime  int     `json:"apparentTemperatureMaxTime"`
		} `json:"data"`
	} `json:"daily"`
	Flags struct {
		Sources        []string `json:"sources"`
		NearestStation float64  `json:"nearest-station"`
		Units          string   `json:"units"`
	} `json:"flags"`
	Offset int `json:"offset"`
}

// weather returns a forcast summary from Darksky
func weather(conn *irc.Connection, r string, event *irc.Event, conf config.Vars) {
	if len(conf.Darksky) <= 1 {
		log.Println("dark sky api key not found")
		return
	}

	a := strings.Split(event.Message(), " ")

	if !validInput(a) {
		conn.Privmsg(r, fmt.Sprintf("weather only accepts 5 digit zip codes"))
		return
	}

	file, err := os.Open(conf.Zipcodes)

	if err != nil {
		conn.Privmsg(r, fmt.Sprintf("zipcode data not found"))
		return
	}

	latitude, longitude := getCoordinates(a[1], file)

	endpoint := fmt.Sprintf("https://api.darksky.net/forecast/%s/%s,%s", conf.Darksky, latitude, longitude)

	resp, err := http.Get(endpoint)

	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	var con forcast

	json.NewDecoder(resp.Body).Decode(&con)

	humidity := strconv.FormatFloat(con.Currently.Humidity, 'f', 2, 64)[2:]
	b := fmt.Sprintf("%s, Wind %.0f mph, Humidity %s%%, Temperature %.0f°", con.Currently.Summary, con.Currently.WindSpeed, humidity, con.Currently.Temperature)
	conn.Privmsg(r, b)

	return
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

	if i != 5 {
		return false
	}

	return true
}

// getCoordinates resolves estimated gps coorinates from a zipcode
func getCoordinates(query string, file *os.File) (string, string) {
	var z *zipcodes
	err := json.NewDecoder(file).Decode(&z)

	if err != nil {
		return "21.343331", "-157.941721"
	}

	for _, v := range z.Data {
		if v.Zipcode == query {
			return v.Latitude, v.Longitude
		}
	}

	return "21.343331", "-157.941721"
}
