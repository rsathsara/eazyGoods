package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	rice "github.com/GeertJohan/go.rice"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var connectString = "sqlserver://developer:max@123@149.28.138.109?database=MaxPOS_EazyGoods"
var dbServer = "mssql"
var err error
var db *sql.DB
var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("eazy123ghst60nsvRFD@12")
	store = sessions.NewCookieStore(key)
)

var modal = Modal{}
var apiResponse = APIResponse{}
var responseMsg = ResponseMsg{}
var api *API
var services = []Services{
	Services{ID: 1, Name: "main", URL: "http://localhost:3250/api/", APIName: "/eazyGoods_api/"},
	// Services{ID: 2, Name: "main", URL: "http://localhost:3250/api/", APIName: "/test_api/"},
}

func main() {
	// Clear Memory
	modal.Data = nil
	requestHandler()
}

func requestHandler() {
	router := mux.NewRouter()
	// Route for eazyGoods_api
	api = &API{Name: "/eazyGoods_api/"}
	router.PathPrefix("/eazyGoods_api/").Handler(http.HandlerFunc(api.apiHandler))

	// Route for test_api
	// api = &API{Name: "/test_api/"}
	// router.PathPrefix("/test_api/").Handler(http.HandlerFunc(api.apiHandler))

	// Route for eazyGoods_web templates
	router.HandleFunc("/", homePage)
	router.HandleFunc("/homePage", homePage)
	router.HandleFunc("/loginPage", loginPage)
	router.HandleFunc("/billingFormPage", billingFormPage)
	router.HandleFunc("/grnFormPage", grnFormPage)
	router.HandleFunc("/itemFormPage", itemFormPage)
	router.HandleFunc("/reportPage", reportPage)
	router.HandleFunc("/logout", logout)
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/sessionDetails", getSessionDetails).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("static_files").HTTPBox()))

	port := os.Getenv("HTTP_PLATFORM_PORT")
	if port == "" {
		port = "3240"
	}
	fmt.Println("Running On Port:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
