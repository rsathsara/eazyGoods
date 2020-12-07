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

func main() {
	requestHandler()
}

func requestHandler() {
	router := mux.NewRouter()

	router.HandleFunc("/", defaultFunction)

	port := os.Getenv("HTTP_PLATFORM_PORT")
	if port == "" {
		port = "3250"
	}
	fmt.Println("Running On Port:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func defaultFunction(w http.ResponseWriter, r *http.Request) {

}
