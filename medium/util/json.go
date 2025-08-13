package util

import (
	"encoding/json"
	"net/http"
)

func ReadJson(r *http.Request, data any) error {
	read := json.NewDecoder(r.Body)
	return read.Decode(data)
}

func JsonResponse(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
