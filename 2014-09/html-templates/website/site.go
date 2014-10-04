package main

import (
	"html/template"
	"log"
	"net/http"
)

var page = template.Must(template.ParseGlob("*.html"))

func main() {
	http.HandleFunc("/style.css", func(wr http.ResponseWriter, req *http.Request) {
		http.ServeFile(wr, req, "style.css")
	})

	http.HandleFunc("/", welcome)
	http.HandleFunc("/blog", blog)
	http.ListenAndServe(":8000", nil)
}

func welcome(wr http.ResponseWriter, req *http.Request) {
	err := page.ExecuteTemplate(wr, "welcome.html", welcomeData)
	if err != nil {
		log.Println(err)
	}
}

func blog(wr http.ResponseWriter, req *http.Request) {
	err := page.ExecuteTemplate(wr, "blog.html", blogData)
	if err != nil {
		log.Println(err)
	}
}

type Page struct {
	Header, Footer interface{}
	Navigation     interface{}
	Content        interface{}
}

var welcomeData = Page{
	Header:     struct{ Title string }{"Sample blog"},
	Footer:     map[string]interface{}{"Owner": "Jane Doe", "Year": 2012},
	Navigation: []struct{ URL, Link string }{{"/", "Welcome"}, {"/blog", "The Blog"}},
}

var blogData = welcomeData
