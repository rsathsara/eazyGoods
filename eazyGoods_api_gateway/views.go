package main

import (
	"html/template"
	"net/http"
)

func loginPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static_files/templates/login.html")
	t.Execute(w, nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if sessionResponse := sessionCheck(w, r); !sessionResponse {
		redirectToLoginPage(w, r)
	}
	t, _ := template.ParseFiles("static_files/templates/main.html")
	t.Execute(w, nil)
}

func billingListPage(w http.ResponseWriter, r *http.Request) {
	if sessionResponse := sessionCheck(w, r); !sessionResponse {
		redirectToLoginPage(w, r)
	}
	t, _ := template.ParseFiles("static_files/templates/billingList.html")
	t.Execute(w, nil)
}

func billingFormPage(w http.ResponseWriter, r *http.Request) {
	if sessionResponse := sessionCheck(w, r); !sessionResponse {
		redirectToLoginPage(w, r)
	}
	t, _ := template.ParseFiles("static_files/templates/billingForm.html")
	t.Execute(w, nil)
}

func grnListPage(w http.ResponseWriter, r *http.Request) {
	if sessionResponse := sessionCheck(w, r); !sessionResponse {
		redirectToLoginPage(w, r)
	}
	t, _ := template.ParseFiles("static_files/templates/grnList.html")
	t.Execute(w, nil)
}

func grnFormPage(w http.ResponseWriter, r *http.Request) {
	if sessionResponse := sessionCheck(w, r); !sessionResponse {
		redirectToLoginPage(w, r)
	}
	t, _ := template.ParseFiles("static_files/templates/grnForm.html")
	t.Execute(w, nil)
}

func redirectToLoginPage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/loginPage", 302)
	return
}
