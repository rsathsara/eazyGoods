package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

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
var responseMsg = ResponseMsg{}
var modal = Modal{}

func main() {
	// Clear memory
	modal.Data = nil
	requestHandler()
	// mainDb()
}

func requestHandler() {
	router := mux.NewRouter()

	router.HandleFunc("/", defaultFunction)
	// Billing
	router.HandleFunc("/api/bills", getBills).Methods("GET")        // Done
	router.HandleFunc("/api/bills/{id}", getBill).Methods("GET")    // Done
	router.HandleFunc("/api/bills", createBill).Methods("POST")     // Done
	router.HandleFunc("/api/bills/{id}", updateBill).Methods("PUT") // Done
	// GRN
	router.HandleFunc("/api/grns", getGrns).Methods("GET")        // Done
	router.HandleFunc("/api/grns/{id}", getGrn).Methods("GET")    // Done
	router.HandleFunc("/api/grns", createGrn).Methods("POST")     // Done
	router.HandleFunc("/api/grns/{id}", updateGrn).Methods("PUT") // Done
	// User
	router.HandleFunc("/api/users", getUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	// Items
	router.HandleFunc("/api/items", getItems).Methods("GET")        // Done
	router.HandleFunc("/api/items/{id}", getItem).Methods("GET")    // Done
	router.HandleFunc("/api/items", createItem).Methods("POST")     // Done
	router.HandleFunc("/api/items/{id}", updateItem).Methods("PUT") // Done
	// Customers
	router.HandleFunc("/api/customers", getCustomers).Methods("GET") // Done
	router.HandleFunc("/api/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/api/customers", createCustomer).Methods("POST")
	router.HandleFunc("/api/customers/{id}", updateCustomer).Methods("PUT")
	// Suppliers
	router.HandleFunc("/api/suppliers", getSuppliers).Methods("GET") // Done
	router.HandleFunc("/api/suppliers/{id}", getSupplier).Methods("GET")
	router.HandleFunc("/api/suppliers", createSupplier).Methods("POST")
	router.HandleFunc("/api/suppliers/{id}", updateSupplier).Methods("PUT")
	// Reports
	router.HandleFunc("/api/reports", getReports).Methods("GET")            // Done
	router.HandleFunc("/api/reports/{id}", reportGenerator).Methods("POST") // Done
	// Other
	router.HandleFunc("/api/newNumbers/{docType}", getNewNumber).Methods("GET") // Done
	router.HandleFunc("/api/units", getUnits).Methods("GET")                    // Done
	router.HandleFunc("/api/itemCat1", getItemCat1).Methods("GET")              // Done
	router.HandleFunc("/api/itemCat2", getItemCat2).Methods("GET")              // Done

	port := os.Getenv("HTTP_PLATFORM_PORT")
	if port == "" {
		port = "3250"
	}
	fmt.Println("Running On Port:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// func mainDb() {
// 	db, err = sql.Open(dbServer, connectString)
// 	errorHandler(w, err)
// 	defer db.Close()
// }

func defaultFunction(w http.ResponseWriter, r *http.Request) {

}
