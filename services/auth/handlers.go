package main

import (
	"encoding/json"
	"net/http"
)

func GetUserProfilesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userProfiles := GetUserProfiles()
	json.NewEncoder(w).Encode(userProfiles)
}
