package main

import (
	"encoding/json"
	"net/http"
)

func emptyString(data string) bool {
	if len(data) == 0 {
		return true
	}
	return false
}

func errorHandler(w http.ResponseWriter, err error) {
	if err != nil {
		responseMsg.Status = "Error"
		responseMsg.Msg = string(err.Error())
		json.NewEncoder(w).Encode(responseMsg)
	}
}
