package controllers

import (
	"app/src/app/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func FetchAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	models.GetAllUsers(&users)
	respBody, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)
}

func FetchUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var user models.User

	models.GetUser(&user, id)
	respBody, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)
}
