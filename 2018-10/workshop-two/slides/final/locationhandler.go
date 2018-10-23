package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"

	"github.com/edmontongo/darksky"
)

type locationHandler struct {
	handler

	forecastSvc forecaster
}

type forecaster interface {
	Forecast(darksky.Location) (*darksky.Weather, error)
}

func newLocationHandler(forecaster forecaster, cache *templateCache) locationHandler {
	return locationHandler{
		forecastSvc: forecaster,
		handler: handler{
			templates: cache,
			viewPaths: []string{"./templates/app.html", "views/location.html"},
			helpers: template.FuncMap{
				"toCelsius": func(tempF float64) string {
					degC := math.Round((tempF - 32) * 5 / 9)
					return fmt.Sprintf("%.0f°C", degC)
				},
			},
		},
	}
}

func (lh locationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	q := r.URL.Query()

	// parse querystring params
	lat, err := strconv.ParseFloat(q.Get("lat"), 10)
	if err != nil {
		http.Error(w, "Invalid lat querystring param", http.StatusBadRequest)
		return
	}
	lng, err := strconv.ParseFloat(q.Get("lng"), 10)
	if err != nil {
		http.Error(w, "Invalid lng querystring param", http.StatusBadRequest)
		return
	}

	// obtain forecast
	var weather *darksky.Weather
	location := darksky.Location{Lat: lat, Long: lng}
	if weather, err = lh.forecastSvc.Forecast(location); err != nil {
		http.Error(w, "failed to fetch weather details", http.StatusInternalServerError)
		return
	}

	// handle response
	lh.render(w, r, weather)
}
