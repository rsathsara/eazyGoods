package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("eazy123ghst60nsvRFD@12")
	store = sessions.NewCookieStore(key)
)

var modal = Modal{}

func main() {
	// Clear memory
	modal.Data = nil
	requestHandler()
}

func requestHandler() {
	router := mux.NewRouter()

	router.HandleFunc("/", defaultFunction)
	// Billing
	router.HandleFunc("/api/bills", getBills).Methods("GET")
	router.HandleFunc("/api/bills/{id}", getBill).Methods("GET")
	router.HandleFunc("/api/bills", createBill).Methods("POST")
	router.HandleFunc("/api/bills/{id}", updateBill).Methods("PUT")
	// GRN
	router.HandleFunc("/api/grns", getGrns).Methods("GET")
	router.HandleFunc("/api/grns/{id}", getGrn).Methods("GET")
	router.HandleFunc("/api/grns", createGrn).Methods("POST")
	router.HandleFunc("/api/grns/{id}", updateGrn).Methods("PUT")
	// User
	router.HandleFunc("/api/users", getUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", getUser).Methods("GET")

	port := os.Getenv("HTTP_PLATFORM_PORT")
	if port == "" {
		port = "3250"
	}
	fmt.Println("Running On Port:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func defaultFunction(w http.ResponseWriter, r *http.Request) {

}
