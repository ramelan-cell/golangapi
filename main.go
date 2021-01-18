package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Post struct {
	Token string `json:"token"`
}

func getLogin(w http.ResponseWriter, r *http.Request) {
	var response Post
	username := r.FormValue("username")
	password := r.FormValue("password")
	response.Token = username + password
	w.Header().Set("Content-Type", "application/json")
	var datamap = make(map[string]interface{})

	datamap["data"] = response

	json.NewEncoder(w).Encode(datamap)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", getLogin).Methods("POST")
	http.ListenAndServe(":8000", router)
}
