package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/edmontongo/darksky"
)

// runme start OMIT
// runme
// runme end OMIT

func main() {
	var secret string
	flag.StringVar(&secret, "secret", "", "DarkSky API key (required)")
	port := flag.String("port", "3000", "Port service will listen on")
	flag.Parse()
	if len(secret) == 0 {
		secret = os.Getenv("secret")
	}
	if len(secret) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// interface svc start OMIT
	darkskySvc := darksky.New(secret)
	http.Handle("/", locationHandler{forecastSvc: darkskySvc})
	// interface svc end OMIT

	fmt.Printf("Listening on :%s\n", *port)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		fmt.Printf("failed to start server: %v", err)
		os.Exit(1)
	}
}

// interface start OMIT
type forecaster interface {
	Forecast(darksky.Location) (*darksky.Weather, error)
}

type locationHandler struct {
	forecastSvc forecaster
}

// interface end OMIT

func (lh locationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// latstart OMIT
	q := r.URL.Query()

	lat, err := strconv.ParseFloat(q.Get("lat"), 10)
	if err != nil {
		http.Error(w, "Invalid lat querystring param", http.StatusBadRequest)
		return
	}
	// latend OMIT

	lng, err := strconv.ParseFloat(q.Get("lng"), 10)
	if err != nil {
		http.Error(w, "Invalid lng querystring param", http.StatusBadRequest)
		return
	}

	// forecaststart OMIT
	weather, err := lh.forecastSvc.Forecast(darksky.Location{Lat: lat, Long: lng})
	if err != nil {
		http.Error(w, "failed to fetch weather details", http.StatusInternalServerError)
		return
	}
	// forecastend OMIT

	// contenttypestart OMIT
	accept := r.Header.Get("Accept")
	contentType := "application/json"
	if strings.Contains(accept, "text/html") {
		contentType = "text/html"
	}
	w.Header().Add("Content-Type", contentType)

	switch contentType {
	// htmlstart OMIT
	case "text/html":
		t, err := template.New("app.html").
			Funcs(template.FuncMap{
				"toCelsius": func(tempF float64) string {
					degC := math.Round((tempF - 32) * 5 / 9)
					return fmt.Sprintf("%.0f°C", degC)
				},
			}).
			ParseFiles(
				"templates/app.html",
				"views/location.html",
			)
		if err != nil {
			http.Error(w, "template gen failed", http.StatusInternalServerError)
			return
		}
		if err = t.Execute(w, weather); err != nil {
			http.Error(w, "render failed", http.StatusInternalServerError)
			return
		}
	// htmlend OMIT
	// jsonmarshalstart OMIT
	case "application/json":
		b, err := json.Marshal(weather)
		if err != nil {
			http.Error(w, "Oops something went wrong", http.StatusInternalServerError)
			return
		}
		w.Write(b)
	}
	// jsonmarshalend OMIT
}
