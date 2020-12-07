package main

import (
	"html/template"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
)

func loginPage(w http.ResponseWriter, r *http.Request) {
	// use this code when serving template without rice box
	// t, _ := template.ParseFiles("static_files/templates/login.html")
	// t.Execute(w, nil)
	box, _ := rice.FindBox("static_files/templates")
	t, _ := box.String("login.html")
	tmplMessage, _ := template.New("message").Parse(t)
	tmplMessage.Execute(w, nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if sessionResponse := sessionCheck(w, r); !sessionResponse {
		redirectToLoginPage(w, r)
	}
	// t, _ := template.ParseFiles("static_files/templates/main.html")
	// t.Execute(w, nil)
	box, _ := rice.FindBox("static_files/templates")
	t, _ := box.String("main.html")
	tmplMessage, _ := template.New("message").Parse(t)
	tmplMessage.Execute(w, nil)
}

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

func redirectToLoginPage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/loginPage", 302)
	return
}

func redirectToHomePage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/homePage", 302)
	return
}
