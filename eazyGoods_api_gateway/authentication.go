package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)

	http.Redirect(w, r, "/loginPage", 302)
	return
}

func login(w http.ResponseWriter, r *http.Request) {
	var dbUname string
	var dbPswd string
	var name string
	var uID int

	r.ParseForm()
	uname := r.FormValue("username") // Data from the form
	pwd := r.FormValue("password")   // Data from the form

	emptyUname := emptyString(uname)
	emptyPswd := emptyString(pwd)

	if emptyUname || emptyPswd {
		responseMsg.Status = "Warning"
		responseMsg.Msg = "There is Empty Field"
	} else {
		db, err := sql.Open(dbServer, connectString)
		errorHandler(w, err)
		defer db.Close()

		rows := db.QueryRow("SELECT UserIdx, LoginName, UserName, UserPassword FROM SysUser where LoginName= $1", uname)
		err = rows.Scan(&uID, &dbUname, &name, &dbPswd)
		if err != nil {
			if err == sql.ErrNoRows {
				responseMsg.Status = "Warning"
				responseMsg.Msg = "Invalid Username"
			}
		}

		if err := bcrypt.CompareHashAndPassword([]byte(dbPswd), []byte(pwd)); err != nil {
			responseMsg.Status = "Warning"
			responseMsg.Msg = "Invalid Password"
		} else {
			session, _ := store.Get(r, "cookie-name")
			session.Values["authenticated"] = true
			session.Values["username"] = uname
			session.Values["userId"] = uID
			session.Values["name"] = name
			session.Save(r, w)
			responseMsg.Status = "Success"
			responseMsg.Msg = ""
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	response, _ := json.Marshal(responseMsg)
	apiResponse.Status = 200
	apiResponse.Body = string(response)
	result, _ := json.Marshal(apiResponse)
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

func getSessionDetails(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	var sessionDetails = SessionDetails{
		ID:       fmt.Sprintf("%v", session.Values["userId"]),
		Username: fmt.Sprintf("%v", session.Values["username"]),
		Name:     fmt.Sprintf("%v", session.Values["name"]),
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(sessionDetails)
}
