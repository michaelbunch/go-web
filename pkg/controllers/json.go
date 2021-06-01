package controllers

import (
	"encoding/json"
	"net/http"
)

// SendJSONResponse sends the client a JSON-API formatted payload
func SendJSONResponse(w http.ResponseWriter, data interface{}) {
	payload, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}
