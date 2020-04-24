package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// JSONReturn returns server response in JSON format.
func JSONReturn(w http.ResponseWriter, statusCode int, jsonObject interface{}) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(jsonObject)
	if err != nil {
		log.Fatal("could not encode json return: " + err.Error())
	}
}

// JSONError returns error from server in JSON format.
func JSONError(w http.ResponseWriter, err error, httpStatus int) {
	JSONReturn(w, httpStatus, resultError{
		Result:  false,
		Code:    -1,
		Details: fmt.Sprintf("Unknown error: %v", err),
	})
}

type resultError struct {
	Result  bool   `json:"result"`
	Code    int64  `json:"code"`
	Details string `json:"message"`
}
