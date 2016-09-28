package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type page struct {
	Title string
	Body  []byte
}

func main() {
	logger, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(logger)
	log.Println("application start")
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Println("server start")
	http.ListenAndServe(":80", nil)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("call saveHandler")
	title := r.URL.Path[6:]
	body := r.FormValue("message")
	p := &page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("call editHandler")
	title := r.URL.Path[6:]
	p, err := load(title)
	if err != nil {
		p = &page{Title: title}
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("call viewHandler")
	title := r.URL.Path[6:]
	p, _ := load(title)
	t, _ := template.ParseFiles("layout.html")
	t.Execute(w, p)
}

func (p *page) save() error {
	log.Println("save " + p.Title + ".txt")
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func load(title string) (*page, error) {
	log.Println("load " + title + ".txt")
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &page{Title: title, Body: body}, nil
}
