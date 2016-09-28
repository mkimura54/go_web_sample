package main

import (
	"net/http"
	"text/template"
)

type page struct {
	Title string
	Count int
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	pg := page{"Hello, World", 1}
	//tmpl, err := template.New("new").Parse("{{.Title}} {{.Count}} count")
	tmpl, err := template.ParseFiles("layout.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, pg)
	if err != nil {
		panic(err)
	}
}
