package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (api *API) apiHandler(w http.ResponseWriter, r *http.Request) {
	if sessionResponse := sessionCheck(w, r); !sessionResponse {
		redirectToLoginPage(w, r)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var result []byte
	URL := getURL(r.URL.Path, api.Name)
	body, _ := ioutil.ReadAll(r.Body)
	if len(URL) > 0 {
		switch r.Method {
		case "GET":
			response, err := http.Get(URL)
			result = apiResposeHandler(response, err)
		case "POST":
			response, err := http.Post(URL, "application/json", bytes.NewBuffer(body))
			result = apiResposeHandler(response, err)
		case "PUT":
			client := &http.Client{}
			req, err := http.NewRequest("PUT", URL, bytes.NewBuffer(body))
			if err != nil {
				apiResponse.Status = 500
				apiResponse.Body = string(err.Error())
				result, _ = json.Marshal(apiResponse)
			}
			response, err := client.Do(req)
			result = apiResposeHandler(response, err)
		case "Delete":
			client := &http.Client{}
			req, err := http.NewRequest("DELETE", URL, bytes.NewBuffer(body))
			if err != nil {
				apiResponse.Status = 500
				apiResponse.Body = string(err.Error())
				result, _ = json.Marshal(apiResponse)
			}
			response, err := client.Do(req)
			result = apiResposeHandler(response, err)
		default:
			apiResponse.Status = 405
			apiResponse.Body = "405 Method Not Allowed"
			result, _ = json.Marshal(apiResponse)
		}
	} else {
		apiResponse.Status = 404
		apiResponse.Body = "404 page not found"
		result, _ = json.Marshal(apiResponse)
	}
	w.Write(result)
}

func getURL(urlPath string, apiName string) string {
	urlPart := strings.Split(urlPath, "/")
	var serviceURL string
	var serviceRequest string
	for _, v := range services {
		if v.APIName == apiName && v.Name == urlPart[2] {
			serviceURL = v.URL
			break
		}
	}
	if len(serviceURL) > 0 {
		serviceRequest = url.PathEscape(strings.TrimPrefix(urlPath, apiName+urlPart[2]+"/"))
		return serviceURL + serviceRequest
	}
	return ""
}

func apiResposeHandler(response *http.Response, err error) []byte {
	if err != nil {
		apiResponse.Status = 500
		apiResponse.Body = string(err.Error())
		result, _ := json.Marshal(apiResponse)
		return result
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		apiResponse.Status = 500
		apiResponse.Body = string(err.Error())
		result, _ := json.Marshal(apiResponse)
		return result
	}
	apiResponse.Status = response.StatusCode
	apiResponse.Body = string(responseBody)
	result, _ := json.Marshal(apiResponse)
	return result
}
