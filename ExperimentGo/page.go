package main

import (
	"fmt"
	"net/http"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hi there, i love %s! ", r.URL.Path[1:])
}

func main()  {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}