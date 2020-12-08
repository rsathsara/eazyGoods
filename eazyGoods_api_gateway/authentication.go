package main

import (
	"encoding/json"
	"net/http"
)

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)

	http.Redirect(w, r, "/loginPage", 200)
	return
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	uname := r.FormValue("username") // Data from the form
	// pwd := r.FormValue("password")   // Data from the form

	session, _ := store.Get(r, "cookie-name")

	session.Values["authenticated"] = true
	session.Values["username"] = uname
	session.Save(r, w)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	result, _ := json.Marshal(true)
	w.Write(result)
}

func sessionCheck(w http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, "cookie-name")
	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return false
	}
	return true
}
