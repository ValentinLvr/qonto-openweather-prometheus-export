package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// We set Temperature and rain precipitation as gauge metrics because its values can only go up or down
// following the https://prometheus.io/docs/concepts/metric_types/#gauge recommendations
// we add a `city` label to have a filtering capibility
var (
	currentTemperature = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "current_temperature",
			Help: "Current temperature observed",
		},
		[]string{"city"},
	)
	forecastTemperature = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "forecast_temperature",
			Help: "Forecast temperature",
		},
		[]string{"city"},
	)
	currentPrecipitation = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "current_precipitation",
			Help: "Current rain precipitation observed",
		},
		[]string{"city"},
	)
	forecastPrecipitation = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "forecast_precipitation",
			Help: "Forecast rain precipitation",
		},
		[]string{"city"},
	)
)

func StartServer() {
	fmt.Println("Starting prometheus server, listening on port 2112")
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":2112", nil)
	if err != nil {
		log.Fatalf("error starting prometheus server: %s", err.Error())
	}
}

func RegisterMetrics() {
	prometheus.Register(currentTemperature)
	prometheus.Register(forecastTemperature)
	prometheus.Register(currentPrecipitation)
	prometheus.Register(forecastPrecipitation)
}
