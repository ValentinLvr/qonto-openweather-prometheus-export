package internal

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type CityGeoCoding struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}
type CityGeoCodingResponse []CityGeoCoding

const (
	geoCodingAPIUrl = "http://api.openweathermap.org/geo/1.0/direct"
)

// Retrieve Latitude & Longitude coordinates giving the city name
func GetCoordinatesByCityName(cityName string) (latitude float64, longitude float64) {
	resp, err := http.Get(geoCodingAPIUrl + "?q=" + cityName + "&appid=" + appID)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	city := CityGeoCodingResponse{}
	err = json.Unmarshal(body, &city)
	if err != nil {
		log.Fatal(err.Error())
	}
	return city[0].Latitude, city[0].Longitude
}
