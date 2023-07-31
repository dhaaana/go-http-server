package utils

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	resp := make(map[string]interface{})
	resp["message"] = message
	resp["status"] = statusCode
	resp["data"] = data
	jsonResp, err := json.Marshal(resp)

	if err != nil {
		LogError("Error happened in JSON marshal. Err:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResp)

}
