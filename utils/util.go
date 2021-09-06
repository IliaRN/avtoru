package utils

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Unable to marshall to json", http.StatusInternalServerError)
		return
	}
}
