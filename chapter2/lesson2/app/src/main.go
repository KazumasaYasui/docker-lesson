package main

import (
	"app/src/app/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello golang from docker!")
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/users", controllers.FetchAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.FetchUser).Methods("GET")

	log.Println("starting server on 8080 port...")
	http.ListenAndServe(":8080", r)
}
