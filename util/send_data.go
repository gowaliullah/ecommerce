package util

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
