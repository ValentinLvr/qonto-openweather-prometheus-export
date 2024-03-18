package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	WeatherAPIUrl = "https://api.openweathermap.org/data/3.0/onecall"
)

var (
	appID = os.Getenv("WEATHER_APP_ID")
)

type weatherData struct {
	Lat      float64      `json:"lat"`
	Lon      float64      `json:"lon"`
	Timezone string       `json:"timezone"`
	Current  currentData  `json:"current"`
	Hourly   []hourlyData `json:"hourly"`
}

type currentData struct {
	Timestamp   int64    `json:"dt"`
	Temperature float64  `json:"temp"`
	Rain        rainData `json:"rain"`
}

type rainData struct {
	Precipitation float64 `json:"1h"`
}

type hourlyData struct {
	Timestamp   int64    `json:"dt"`
	Temperature float64  `json:"temp"`
	Rain        rainData `json:"rain"`
}

// Retrieve city coordinates & make a simple request with these coordinates to get current & forecasted weather data
func GetCityData(cityName string) {
	Lat, Lon := GetCoordinatesByCityName(cityName)
	cityWeatherData := weatherData{}

	for {
		// ----- Get Current & forecast weather Data -----
		resp, err := http.Get(
			WeatherAPIUrl +
				"?lat=" + strconv.FormatFloat(Lat, 'f', -1, 64) +
				"&lon=" + strconv.FormatFloat(Lon, 'f', -1, 64) +
				"&appid=" + appID +
				"&exclude=minutely,daily,alerts" + "&units=metric")
		if err != nil {
			log.Fatalf("error while getting weatherApp API data: %s", err.Error())
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("error reading response body: %s", err.Error())
		}

		err = json.Unmarshal(body, &cityWeatherData)
		if err != nil {
			log.Fatalf("error unmarshaling response body: %s", err.Error())
		}

		err = resp.Body.Close()
		if err != nil {
			log.Fatalf("error closing response body: %s", err.Error())
		}

		// ----- Populate prometheus metrics -----
		currentTemperature.With(prometheus.Labels{"city": cityName}).Set(cityWeatherData.Current.Temperature)
		currentPrecipitation.With(prometheus.Labels{"city": cityName}).Set(cityWeatherData.Current.Rain.Precipitation)
		// Hourly[3] to get 3-hour forecasted data
		forecastTemperature.With(prometheus.Labels{"city": cityName}).Set(cityWeatherData.Hourly[3].Temperature)
		forecastPrecipitation.With(prometheus.Labels{"city": cityName}).Set(cityWeatherData.Hourly[3].Rain.Precipitation)

		// wait 1mn before making another call
		fmt.Printf("Fetch new weather data for city: %s", cityName)
		time.Sleep(60 * time.Second)
	}
}
