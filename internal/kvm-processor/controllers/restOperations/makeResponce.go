package restOperations

import (
	"../../models"
	"encoding/json"
	"net/http"
)

func SimpleMessage(message string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(message))
}

func ReturnVMObject(vm *models.VirtualMachine, w http.ResponseWriter) {

	jsonResponse, err := json.Marshal(vm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonResponse)
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}
