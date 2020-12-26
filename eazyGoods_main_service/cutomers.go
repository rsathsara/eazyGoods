package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
)

// Customer struct
type Customer struct {
	ID         int    `json:"id"`
	Code       string `json:"code"`
	Desription string `json:"description"`
}

// Get All Customers
func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	db, err := sql.Open(dbServer, connectString)
	errorHandler(w, err)
	defer db.Close()
	var customer = Customer{}
	var customers []Customer
	query := `SELECT CustIdx, CustomerCode, CustomerName FROM MstrCustomer`
	result, err := db.Query(query)
	errorHandler(w, err)
	for result.Next() {
		err = result.Scan(&customer.ID, &customer.Code, &customer.Desription)
		errorHandler(w, err)
		customers = append(customers, customer)
	}
	json.NewEncoder(w).Encode(customers)
}

// Get Single Customer
func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var customer Customer
	json.NewEncoder(w).Encode(customer)
}

// Create a New Customer
func createCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body, _ := ioutil.ReadAll(r.Body)
	json.NewEncoder(w).Encode(body)
}

// Update Customer
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body, _ := ioutil.ReadAll(r.Body)
	json.NewEncoder(w).Encode(body)
}

// New Customer No
func newCustomerNo() {

}
