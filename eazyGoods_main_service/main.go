package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	requestHandler()
}

func requestHandler() {
	router := mux.NewRouter()

	router.HandleFunc("/", defaultFunction)

	port := os.Getenv("HTTP_PLATFORM_PORT")
	if port == "" {
		port = "8282"
	}
	fmt.Println("Running On Port:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func defaultFunction(w http.ResponseWriter, r *http.Request) {

}
