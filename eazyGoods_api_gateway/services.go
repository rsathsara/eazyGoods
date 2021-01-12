package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (api *API) apiHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	var addData = map[string]string{
		"SessionUserAuthenticated": fmt.Sprintf("%v", session.Values["authenticated"]),
		"SessionUserID":            fmt.Sprintf("%v", session.Values["userId"]),
		"SessionUsername":          fmt.Sprintf("%v", session.Values["username"]),
		"SessionName":              fmt.Sprintf("%v", session.Values["name"]),
	}
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
			cli := &http.Client{}
			req, _ := customRequestHeader("GET", URL, nil, addData)
			response, err := cli.Do(req)
			result = apiResposeHandler(response, err)
		case "POST":
			cli := &http.Client{}
			req, _ := customRequestHeader("POST", URL, bytes.NewBuffer(body), addData)
			response, err := cli.Do(req)
			result = apiResposeHandler(response, err)
			// response, err := http.Post(URL, "application/json", bytes.NewBuffer(body))
			// result = apiResposeHandler(response, err)
		case "PUT":
			client := &http.Client{}
			req, _ := customRequestHeader("PUT", URL, bytes.NewBuffer(body), addData)
			response, err := client.Do(req)
			result = apiResposeHandler(response, err)
		case "Delete":
			client := &http.Client{}
			req, _ := customRequestHeader("DELETE", URL, bytes.NewBuffer(body), addData)
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

func customRequestHeader(method, path string, body io.Reader, addData map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Session-User-Authenticated", addData["SessionUserAuthenticated"])
	req.Header.Add("Session-Username", addData["SessionUsername"])
	req.Header.Add("Session-Name", addData["SessionName"])
	req.Header.Add("Session-User-ID", addData["SessionUserID"])
	return req, nil
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
