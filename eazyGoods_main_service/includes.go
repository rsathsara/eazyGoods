package main

import (
	"encoding/json"
	"net/http"
)

func errorHandler(w http.ResponseWriter, err error) {
	if err != nil {
		responseMsg.Status = "Error"
		responseMsg.Msg = string(err.Error())
		json.NewEncoder(w).Encode(responseMsg)
	}
}
