package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type handler struct {
	viewPaths []string
	helpers   template.FuncMap

	templates *templateCache
}

func (h *handler) render(w http.ResponseWriter, r *http.Request, data interface{}) {
	var contentType string
	accept := r.Header.Get("Accept")
	switch {
	case strings.Contains(accept, "text/html"):
		contentType = "text/html"
	default:
		contentType = "application/json"
	}

	w.Header().Add("Content-Type", contentType)

	switch contentType {
	case "text/html":
		if len(h.viewPaths) == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "at least one view must be passed in")
			return
		}
		tmpl, err := h.templates.get(h.viewPaths, h.helpers)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, fmt.Sprintf("template new failure: %v", err))
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, fmt.Sprintf("template execute failure: %v", err))
			return
		}
	case "application/json":
		b, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Oops something went wrong", http.StatusInternalServerError)
			return
		}
		w.Write(b)
	default:
		http.Error(w, "Only application/json and text/html are supported", http.StatusBadRequest)
	}
}
