package main

import (
	"fmt"
	"html/template"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
)

func loginPage(w http.ResponseWriter, r *http.Request) {
	if sessionResponse := sessionCheck(w, r); sessionResponse {
		redirectToHomePage(w, r)
	}
	box, _ := rice.FindBox("static_files/templates")
	t, _ := box.String("login.html")
	tmplMessage, _ := template.New("message").Parse(t)
	tmplMessage.Execute(w, nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	if sessionResponse := sessionCheck(w, r); !sessionResponse {
		redirectToLoginPage(w, r)
	}
	box, _ := rice.FindBox("static_files/templates")
	t, _ := box.String("main.html")
	tmplMessage, _ := template.New("message").Parse(t)

	var sessionDetails = SessionDetails{}
	sessionDetails.Username = fmt.Sprintf("%v", session.Values["username"])
	sessionDetails.ID = session.Values["userId"].(int)
	sessionDetails.Name = fmt.Sprintf("%v", session.Values["name"])
	// sessionDetails := sessionDetails(w, r)
	tmplMessage.Execute(w, sessionDetails)
}

// func sessionDetails(w http.ResponseWriter, r *http.Request) SessionDetails {
// 	session, _ := store.Get(r, "cookie-name")
// 	var sessionDetails = SessionDetails{}
// 	sessionDetails.Username = fmt.Sprintf("%v", session.Values["username"])
// 	sessionDetails.ID = session.Values["userId"].(int)
// 	sessionDetails.Name = fmt.Sprintf("%v", session.Values["name"])
// 	return sessionDetails
// }

func billingFormPage(w http.ResponseWriter, r *http.Request) {
	if sessionResponse := sessionCheck(w, r); !sessionResponse {
		redirectToLoginPage(w, r)
	}
	box, _ := rice.FindBox("static_files/templates")
	t, _ := box.String("billingForm.html")
	tmplMessage, _ := template.New("message").Parse(t)
	tmplMessage.Execute(w, nil)
}

func grnFormPage(w http.ResponseWriter, r *http.Request) {
	if sessionResponse := sessionCheck(w, r); !sessionResponse {
		redirectToLoginPage(w, r)
	}
	box, _ := rice.FindBox("static_files/templates")
	t, _ := box.String("grnForm.html")
	tmplMessage, _ := template.New("message").Parse(t)
	tmplMessage.Execute(w, nil)
}

func reportPage(w http.ResponseWriter, r *http.Request) {
	if sessionResponse := sessionCheck(w, r); !sessionResponse {
		redirectToLoginPage(w, r)
	}
	box, _ := rice.FindBox("static_files/templates")
	t, _ := box.String("reportPage.html")
	tmplMessage, _ := template.New("message").Parse(t)
	reportList := []ReportList{
		{ID: 1, Description: "GRN List"},
		{ID: 2, Description: "Invoice List"},
		{ID: 3, Description: "Stock Balance Report"},
	}
	tmplMessage.Execute(w, reportList)
}

func redirectToLoginPage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/loginPage", 302)
	return
}

func redirectToHomePage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/homePage", 302)
	return
}
