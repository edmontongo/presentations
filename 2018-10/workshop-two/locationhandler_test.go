package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/edmontongo/darksky"
)

type forecasterMock struct {
	weather *darksky.Weather
	err     error
}

func (f forecasterMock) Forecast(darksky.Location) (*darksky.Weather, error) {
	return f.weather, f.err
}

func TestHTMLResponse(t *testing.T) {
	tmplCache := newTemplateCache()
	lh := newLocationHandler(
		forecasterMock{weather: &darksky.Weather{
			Currently: darksky.Currently{Summary: "hot", Temperature: 90},
		}}, tmplCache,
	)

	r, _ := http.NewRequest("GET", "/?lat=53.5&lng=-113.3", nil)
	r.Header.Add("Accept", "text/html")
	w := httptest.NewRecorder()
	lh.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Fatalf("Unexpected status code: %d != 200", w.Code)
	}

	body := w.Body.String()
	if !strings.Contains(body, "hot") {
		t.Errorf("it should be hot")
	}
}

func TestJSONResponse(t *testing.T) {
	tmplCache := newTemplateCache()
	lh := newLocationHandler(
		forecasterMock{weather: &darksky.Weather{
			Currently: darksky.Currently{Summary: "hot", Temperature: 90},
		}}, tmplCache,
	)

	r, _ := http.NewRequest("GET", "/?lat=53.5&lng=-113.3", nil)
	r.Header.Add("Accept", "application/json")
	w := httptest.NewRecorder()
	lh.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Fatalf("Unexpected status code: %d != 200", w.Code)
	}

	var weather darksky.Weather
	if err := json.Unmarshal(w.Body.Bytes(), &weather); err != nil {
		t.Fatalf("failed unmarshal: %v", err)
	}
	if weather.Currently.Temperature != 90 {
		t.Errorf("it should be 90deg, it was %f", weather.Currently.Temperature)
	}
}
