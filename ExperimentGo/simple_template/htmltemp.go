package main

import (
	"net/http"
	"html/template"
)

type webPage struct {
	Title string 
	Info  string
}

func webPageHandler(w http.ResponseWriter, r *http.Request) {
	p := webPage{Title: "Welcome!", Info: "Page created with Go"}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func main()  {
	http.HandleFunc("/", webPageHandler)
	http.ListenAndServe(":8000", nil)

}