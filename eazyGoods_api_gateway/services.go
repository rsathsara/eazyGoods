package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Modal Struct
type Modal struct {
	Data []Data
}

// Data Struct
type Data struct {
	APIResponse []APIResponse
}

// APIResponse Struct
type APIResponse struct {
	Status int    `json:"status"`
	Body   string `json:"body"`
}

// Services Struct
type Services struct {
	ID   int
	Name string
	URL  string
}

var modal = Modal{}
var apiResponse = APIResponse{}
var services = []Services{
	Services{ID: 1, Name: "main", URL: "http://localhost:3250/api/"},
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	response, err := http.Get("http://localhost:3250/api/bills")
	result := apiResposeHandler(response, err)
	w.Write(result)
}

func apiRequestHandler() {

}

func apiResposeHandler(response *http.Response, err error) []byte {
	if err != nil {
		result, _ := json.Marshal(err.Error())
		return result
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		result, _ := json.Marshal(err.Error())
		return result
	}
	responseStatus := response.StatusCode
	apiResponse.Status = responseStatus
	apiResponse.Body = string(responseBody)
	result, _ := json.Marshal(apiResponse)
	return result
}
