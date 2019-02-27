package main

import (
	"fmt"
	"html"
	"html/template"
	"log"
	"net/http"
)

// Page 表示形式
type Page struct {
	Title string
	Count int
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello,world")
	log.Printf("%v\n", r)
}

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi %q\n", html.EscapeString(r.URL.Path))
	log.Printf("%v\n", r)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{"wdkkou", 1}
	tmpl, err := template.ParseFiles("../html/index.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/hi", hi)
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
