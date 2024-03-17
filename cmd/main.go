package main

import (
	"github.com/ValentinLvr/test-qonto/internal"
)

func init() {
	internal.RegisterMetrics()
}

func main() {
	// we define the cities on which we will get the weather data
	cities := []string{"paris", "pekin", "tokyo", "beijing", "istanbul",
		"mumbai", "manila", "osaka", "lagos", "moscow", "kinshasa", "jakarta",
		"delhi", "london", "seoul",
	}

	// ----------  FETCH WEATHER APP DATA ----------
	// for each city we get the weather data via the API
	// we use go routine to make the calls concurrently
	for _, city := range cities {
		go internal.GetCityData(city)
	}

	// ----------  EXPOSE PROMETHEUS METRICS ----------
	internal.StartServer()
}
