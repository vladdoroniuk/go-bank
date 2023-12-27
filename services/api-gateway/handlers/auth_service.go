package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func AuthHandlers(r *mux.Router) {
	r.HandleFunc("/users", getUsers).Methods("GET")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := struct {
		Name string
		Age  int
	}{
		Name: "John",
		Age:  20,
	}
	json.NewEncoder(w).Encode(user)
}
