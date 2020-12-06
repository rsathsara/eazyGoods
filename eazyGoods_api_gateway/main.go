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
	requestHandler()
}

func requestHandler() {
	var dir string
	router := mux.NewRouter()

	router.HandleFunc("/", homePage)
	router.HandleFunc("/homePage", homePage)
	router.HandleFunc("/loginPage", loginPage)
	router.HandleFunc("/billingListPage", billingListPage)
	router.HandleFunc("/billingFormPage", billingFormPage)
	router.HandleFunc("/grnListPage", grnListPage)
	router.HandleFunc("/grnFormPage", grnFormPage)
	router.HandleFunc("/logout", logout)
	router.HandleFunc("/login", login).Methods("POST")
	// router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("assets").HTTPBox()))
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("static_files").HTTPBox()))
	router.PathPrefix("/static_files/templates/").Handler(http.StripPrefix("/static_files/templates/", http.FileServer(http.Dir(dir))))

	port := os.Getenv("HTTP_PLATFORM_PORT")
	if port == "" {
		port = "8181"
	}
	fmt.Println("Running On Port:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
