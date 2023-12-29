package main

import (
	"encoding/json"
	"net/http"
)

func GetUserProfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userProfiles, _ := GetUserProfiles(w, r)
	json.NewEncoder(w).Encode(userProfiles)
}
