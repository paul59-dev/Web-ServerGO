package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//http.HandleFunc("/", homeHandler) // Response with all route

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Router
	r := mux.NewRouter()
	r.HandleFunc("/book/{title}/page/{page}", bookHandler)

	http.Handle("/", r)

	http.ListenAndServe(":80", nil)
}

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
// }

func bookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var title string = vars["title"]
	var page string = vars["page"]

	fmt.Fprintf(w, "You've requested the book %s on page %s\n", title, page)
}
