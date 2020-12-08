package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func mainService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	response, err := http.Get("http://localhost:3250/api/bills")
	if err != nil {
		result, _ := json.Marshal(err.Error())
		w.Write(result)
		return
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		result, _ := json.Marshal(err.Error())
		w.Write(result)
		return
	}
	w.Write(responseBody)
}

func apiResposeHandler() {

}
