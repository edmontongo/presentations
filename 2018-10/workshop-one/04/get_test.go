package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestForecast(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{}")
	})

	ts := httptest.NewServer(handler)
	defer ts.Close()

	// BEGIN OMIT
	ds := NewDarkSky(ts.URL, "secret")
	ds.Forecast(Location{})
	// END OMIT
}
