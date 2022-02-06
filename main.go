package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tyhuang06/Gallery/views"
)

var (
	homeView    *views.View
	contactView *views.View
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")

	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/contact", contactHandler)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
