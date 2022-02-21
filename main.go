package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tyhuang06/Gallery/controllers"
)

func main() {
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	r := mux.NewRouter()

	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
