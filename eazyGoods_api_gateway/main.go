package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("eazy123ghst60nsvRFD@12")
	store = sessions.NewCookieStore(key)
)

func main() {
	modal.Data = nil
	requestHandler()
}

func requestHandler() {
	router := mux.NewRouter()
	router.PathPrefix("/eazyGoods_api/").Handler(http.HandlerFunc(apiHandler))
	router.HandleFunc("/", redirectToHomePage)
	router.HandleFunc("/homePage", homePage)
	router.HandleFunc("/loginPage", loginPage)
	router.HandleFunc("/billingFormPage", billingFormPage)
	router.HandleFunc("/grnFormPage", grnFormPage)
	router.HandleFunc("/logout", logout)
	router.HandleFunc("/login", login).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("static_files").HTTPBox()))

	port := os.Getenv("HTTP_PLATFORM_PORT")
	if port == "" {
		port = "3240"
	}
	fmt.Println("Running On Port:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
