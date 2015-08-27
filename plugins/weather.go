package plugins

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"net/http"
	"encoding/json"
	"strconv"
	"strings"
	)

type Conditions struct {
	Response struct {
		Version string `json:"version"`
		Termsofservice string `json:"termsofService"`
		Features struct {
			Conditions int `json:"conditions"`
		} `json:"features"`
		Error struct {
			Type string `json:"type"`
			Description string `json:"description"`
		} `json:"error"`
	} `json:"response"`
	CurrentObservation struct {
		Image struct {
			URL string `json:"url"`
			Title string `json:"title"`
			Link string `json:"link"`
		} `json:"image"`
		DisplayLocation struct {
			Full string `json:"full"`
			City string `json:"city"`
			State string `json:"state"`
			StateName string `json:"state_name"`
			Country string `json:"country"`
			CountryIso3166 string `json:"country_iso3166"`
			Zip string `json:"zip"`
			Magic string `json:"magic"`
			Wmo string `json:"wmo"`
			Latitude string `json:"latitude"`
			Longitude string `json:"longitude"`
			Elevation string `json:"elevation"`
		} `json:"display_location"`
		ObservationLocation struct {
			Full string `json:"full"`
			City string `json:"city"`
			State string `json:"state"`
			Country string `json:"country"`
			CountryIso3166 string `json:"country_iso3166"`
			Latitude string `json:"latitude"`
			Longitude string `json:"longitude"`
			Elevation string `json:"elevation"`
		} `json:"observation_location"`
		Estimated struct {
		} `json:"estimated"`
		StationID string `json:"station_id"`
		ObservationTime string `json:"observation_time"`
		ObservationTimeRfc822 string `json:"observation_time_rfc822"`
		ObservationEpoch string `json:"observation_epoch"`
		LocalTimeRfc822 string `json:"local_time_rfc822"`
		LocalEpoch string `json:"local_epoch"`
		LocalTzShort string `json:"local_tz_short"`
		LocalTzLong string `json:"local_tz_long"`
		LocalTzOffset string `json:"local_tz_offset"`
		Weather string `json:"weather"`
		TemperatureString string `json:"temperature_string"`
		TempF float64 `json:"temp_f"`
		TempC float64 `json:"temp_c"`
		RelativeHumidity string `json:"relative_humidity"`
		WindString string `json:"wind_string"`
		WindDir string `json:"wind_dir"`
		WindDegrees int `json:"wind_degrees"`
		WindMph float64 `json:"wind_mph"`
		WindGustMph float64 `json:"wind_gust_mph"`
		WindKph float64 `json:"wind_kph"`
		WindGustKph string `json:"wind_gust_kph"`
		PressureMb string `json:"pressure_mb"`
		PressureIn string `json:"pressure_in"`
		PressureTrend string `json:"pressure_trend"`
		DewpointString string `json:"dewpoint_string"`
		DewpointF int `json:"dewpoint_f"`
		DewpointC int `json:"dewpoint_c"`
		HeatIndexString string `json:"heat_index_string"`
		HeatIndexF string `json:"heat_index_f"`
		HeatIndexC string `json:"heat_index_c"`
		WindchillString string `json:"windchill_string"`
		WindchillF string `json:"windchill_f"`
		WindchillC string `json:"windchill_c"`
		FeelslikeString string `json:"feelslike_string"`
		FeelslikeF string `json:"feelslike_f"`
		FeelslikeC string `json:"feelslike_c"`
		VisibilityMi string `json:"visibility_mi"`
		VisibilityKm string `json:"visibility_km"`
		Solarradiation string `json:"solarradiation"`
		Uv string `json:"UV"`
		Precip1HrString string `json:"precip_1hr_string"`
		Precip1HrIn string `json:"precip_1hr_in"`
		Precip1HrMetric string `json:"precip_1hr_metric"`
		PrecipTodayString string `json:"precip_today_string"`
		PrecipTodayIn string `json:"precip_today_in"`
		PrecipTodayMetric string `json:"precip_today_metric"`
		Icon string `json:"icon"`
		IconURL string `json:"icon_url"`
		ForecastURL string `json:"forecast_url"`
		HistoryURL string `json:"history_url"`
		ObURL string `json:"ob_url"`
		Nowcast string `json:"nowcast"`
	} `json:"current_observation"`
}
func Weather(conn *irc.Connection, resp string, dword string) {
	k := "apikey"

	i := 0
	for _, v := range dword {
		switch {
		case v >= '0' && v <= '9':
			i++
		}
	}
	if i != 5 {
		conn.Privmsgf(resp, "weather only takes 5 digit zip codes")
		return
	}

	endpoint := fmt.Sprintf("http://api.wunderground.com/api/%s/conditions/q/%s.json", k, dword)
	r, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Body.Close()
	var con Conditions
	json.NewDecoder(r.Body).Decode(&con)

	if len(con.Response.Error.Description) > 0 {
		conn.Privmsgf(resp, con.Response.Error.Description)
		return
	}

	humid := strings.Replace(con.CurrentObservation.RelativeHumidity, "%", "%%", -1)
	temp := fmt.Sprintf("%s", strconv.FormatFloat(con.CurrentObservation.TempF, 'f', 1, 64))
	w_mph := strconv.FormatFloat(con.CurrentObservation.WindMph, 'f', 1, 64)
	w_gust := strconv.FormatFloat(con.CurrentObservation.WindGustMph,'f', 1, 64)

	if w_mph == "0.0" {
		w_mph = "0"
	}

	if w_gust == "0.0" {
		w_gust = "0"
	}

	wind := fmt.Sprintf("wind from the %s at %s mph gusting to %s mph", con.CurrentObservation.WindDir, w_mph, w_gust)
	s := fmt.Sprintf("%s, %s, humidity %s, temperature %s°F\n", con.CurrentObservation.Weather, wind, humid, temp)
	conn.Privmsgf(resp, s)
}

