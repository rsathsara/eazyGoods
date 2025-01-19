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
	if sessionResponse := sessionCheck(w, r); !sessionResponse {
		http.Redirect(w, r, "/loginPage", 302)
		return
	}
	box, _ := rice.FindBox("static_files/templates")
	t, _ := box.String("main.html")
	tmplMessage, _ := template.New("message").Parse(t)

	session, _ := store.Get(r, "cookie-name")
	var sessionDetails = SessionDetails{
		ID:       fmt.Sprintf("%v", session.Values["userId"]),
		Username: fmt.Sprintf("%v", session.Values["username"]),
		Name:     fmt.Sprintf("%v", session.Values["name"]),
	}
	tmplMessage.Execute(w, sessionDetails)
}

func billingFormPage(w http.ResponseWriter, r *http.Request) {
	if sessionResponse := sessionCheck(w, r); !sessionResponse {
		http.Redirect(w, r, "/loginPage", 302)
		return
	}
	box, _ := rice.FindBox("static_files/templates")
	t, _ := box.String("billingForm.html")
	tmplMessage, _ := template.New("message").Parse(t)
	tmplMessage.Execute(w, nil)
}

func grnFormPage(w http.ResponseWriter, r *http.Request) {
	if sessionResponse := sessionCheck(w, r); !sessionResponse {
		http.Redirect(w, r, "/loginPage", 302)
		return
	}
	box, _ := rice.FindBox("static_files/templates")
	t, _ := box.String("grnForm.html")
	tmplMessage, _ := template.New("message").Parse(t)
	tmplMessage.Execute(w, nil)
}

func itemFormPage(w http.ResponseWriter, r *http.Request) {
	if sessionResponse := sessionCheck(w, r); !sessionResponse {
		http.Redirect(w, r, "/loginPage", 302)
		return
	}
	box, _ := rice.FindBox("static_files/templates")
	t, _ := box.String("itemForm.html")
	tmplMessage, _ := template.New("message").Parse(t)
	tmplMessage.Execute(w, nil)
}

func reportPage(w http.ResponseWriter, r *http.Request) {
	if sessionResponse := sessionCheck(w, r); !sessionResponse {
		http.Redirect(w, r, "/loginPage", 302)
		return
	}
	box, _ := rice.FindBox("static_files/templates")
	t, _ := box.String("reportPage.html")
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
