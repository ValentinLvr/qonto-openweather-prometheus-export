# OpenWeather prometheus exporter

Little Application that convert [openweathermap](https://openweathermap.org) temperature and rain data to prometheus metrics.

1. We retrieve the coordinates of the top 15 biggest cities in the world. (`internal/geocoding.go`)
2. We define what kind of prometheus metrics we want to share. (`internal/prometheus.go`)
    - Current/forecasted temperatures as gauge metrics
    - Current/forecasted rain precipitation as gauge metrics
2. We make calls to the openweathermap API to get the current & forecasted temperature/rain precipitation each X seconds for these cities.
We, then, update the above metrics for each requests (`internal/client.go`)

## Run

- add your openweathermap appID on the `docker-compose.yml` file (`WEATHER_APP_ID` environement variable)
- run grafana/prometheus/app services
```shell
docker-compose build && docker-compose up
```

## Grafana dashboards
go to `grafana` folder to get the json dashboards