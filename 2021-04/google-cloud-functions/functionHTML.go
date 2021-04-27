package google_cloud_functions

import (
	"fmt"
	"html/template"
	"net/http"
)

// START OMIT
var page = template.Must(template.New("examplePage").Parse(exampleHTML))
var mux http.ServeMux

func init() {
	mux.Handle("/", handle("home"))
	mux.Handle("/A", handle("pageA"))
	mux.Handle("/B", handle("pageB"))
}

func HTML(w http.ResponseWriter, req *http.Request) {
	mux.ServeHTTP(w, req)
}

func handle(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		err := page.ExecuteTemplate(w, name, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "could not write page", err.Error())
		}
	}
}

// END OMIT
