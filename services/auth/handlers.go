package main

import (
	"encoding/json"
	"net/http"
)

type GetHandler struct{}

func NewGetHandler() *GetHandler {
	return &GetHandler{}
}

func (gh *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		gh.GetUserProfiles(w, r)
	}
}

func (gh *GetHandler) GetUserProfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(GetUserProfiles())
	w.Write(data)
}
