package plugins

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"net/http"
	"encoding/json"
	"strconv"
	"strings"
)

type WeatherU struct {
	Response struct {
		Version string `json:"version"`
		Termsofservice string `json:"termsofService"`
		Features struct {
			Conditions int `json:"conditions"`
			Tide int `json:"tide"`
			Astronomy int `json:"astronomy"`
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
		WindGustMph string `json:"wind_gust_mph"`
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
	Tide struct {
		Tideinfo []struct {
			Tidesite string `json:"tideSite"`
			Lat string `json:"lat"`
			Lon string `json:"lon"`
			Units string `json:"units"`
			Type string `json:"type"`
			Tzname string `json:"tzname"`
		} `json:"tideInfo"`
		Tidesummary []struct {
			Date struct {
				Pretty string `json:"pretty"`
				Year string `json:"year"`
				Mon string `json:"mon"`
				Mday string `json:"mday"`
				Hour string `json:"hour"`
				Min string `json:"min"`
				Tzname string `json:"tzname"`
				Epoch string `json:"epoch"`
			} `json:"date"`
			Utcdate struct {
				Pretty string `json:"pretty"`
				Year string `json:"year"`
				Mon string `json:"mon"`
				Mday string `json:"mday"`
				Hour string `json:"hour"`
				Min string `json:"min"`
				Tzname string `json:"tzname"`
				Epoch string `json:"epoch"`
			} `json:"utcdate"`
			Data struct {
				Height string `json:"height"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"tideSummary"`
		Tidesummarystats []struct {
			Maxheight float64 `json:"maxheight"`
			Minheight float64 `json:"minheight"`
		} `json:"tideSummaryStats"`
	} `json:"tide"`
	MoonPhase struct {
		Percentilluminated string `json:"percentIlluminated"`
		Ageofmoon string `json:"ageOfMoon"`
		Phaseofmoon string `json:"phaseofMoon"`
		Hemisphere string `json:"hemisphere"`
		CurrentTime struct {
			Hour string `json:"hour"`
			Minute string `json:"minute"`
		} `json:"current_time"`
		Sunrise struct {
			Hour string `json:"hour"`
			Minute string `json:"minute"`
		} `json:"sunrise"`
		Sunset struct {
			Hour string `json:"hour"`
			Minute string `json:"minute"`
		} `json:"sunset"`
	} `json:"moon_phase"`
	SunPhase struct {
		Sunrise struct {
			Hour string `json:"hour"`
			Minute string `json:"minute"`
		} `json:"sunrise"`
		Sunset struct {
			Hour string `json:"hour"`
			Minute string `json:"minute"`
		} `json:"sunset"`
	} `json:"sun_phase"`
}
func WeatherUnderground(conn *irc.Connection, dword string, query string, resp string) {
	k := "apikey"

	i := 0
	for _, v := range dword {
		switch {
		case v >= '0' && v <= '9':
			i++
		}
	}
	if i != 5 {
		s := fmt.Sprintf("%s only takes 5 digit zip codes", query)
		conn.Privmsgf(resp, s)
		return
	}

	if query == "weather" {
		query = "conditions"
	}

	endpoint := fmt.Sprintf("http://api.wunderground.com/api/%s/%s/q/%s.json", k, query, dword)
	r, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Body.Close()
	var con WeatherU
	json.NewDecoder(r.Body).Decode(&con)

	if len(con.Response.Error.Description) > 0 {
		conn.Privmsgf(resp, con.Response.Error.Description)
		return
	}

	if query == "conditions" {
		humid := strings.Replace(con.CurrentObservation.RelativeHumidity, "%", "%%", -1)
		temp := fmt.Sprintf("%s", strconv.FormatFloat(con.CurrentObservation.TempF, 'f', 1, 64))
		wmph := strconv.FormatFloat(con.CurrentObservation.WindMph, 'f', 1, 64)
		wind := fmt.Sprintf("wind from the %s at %s mph gusting to %s mph", con.CurrentObservation.WindDir, wmph, con.CurrentObservation.WindGustMph)
		s := fmt.Sprintf("%s, %s, humidity %s, temperature %s°F\n", con.CurrentObservation.Weather, wind, humid, temp)
		conn.Privmsgf(resp, s)
	} else if query == "tide" {
		l := strings.Split(con.Tide.Tidesummary[0].Date.Pretty, "on")
		low := fmt.Sprintf("low of %s at %s", con.Tide.Tidesummary[0].Data.Height, strings.TrimSpace(l[0]))
		h := strings.Split(con.Tide.Tidesummary[3].Date.Pretty, "on")
		high := fmt.Sprintf("high of %s at %s", con.Tide.Tidesummary[3].Data.Height, strings.TrimSpace(h[0]))
		lt := strings.Split(con.Tide.Tidesummary[4].Date.Pretty, "on")
		lowt := fmt.Sprintf("low of %s at %s", con.Tide.Tidesummary[0].Data.Height, strings.TrimSpace(lt[0]))
		ht := strings.Split(con.Tide.Tidesummary[7].Date.Pretty, "on")
		hight := fmt.Sprintf("high of %s at %s", con.Tide.Tidesummary[3].Data.Height, strings.TrimSpace(ht[0]))
		s := fmt.Sprintf("Today: %s, %s. Tomorrow: %s, %s.", low, high, lowt, hight)
		conn.Privmsgf(resp, s)
	} else if query == "astronomy" {
		sr := fmt.Sprintf("%s:%s", con.SunPhase.Sunrise.Hour, con.SunPhase.Sunrise.Minute)
		ss := fmt.Sprintf("%s:%s", con.SunPhase.Sunset.Hour, con.SunPhase.Sunset.Minute)
		s := fmt.Sprintf("Sunrise %s, sunset %s.", sr, ss)
		m := fmt.Sprintf("Moonphase %s%s illuminated, age %s, in a %s phase.", con.MoonPhase.Percentilluminated, "%%", con.MoonPhase.Ageofmoon, strings.ToLower(con.MoonPhase.Phaseofmoon))
		f := fmt.Sprintf("%s %s",s, m)
		conn.Privmsgf(resp, f)
	}
}
