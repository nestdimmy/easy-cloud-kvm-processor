package controllers

import (
	"encoding/json"
	"net/http"
)

func Message(fieldName string, message string) map[string]interface{} {
	return map[string]interface{}{fieldName: message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
