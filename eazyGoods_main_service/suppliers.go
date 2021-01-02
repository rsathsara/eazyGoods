package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
)

// Supplier struct
type Supplier struct {
	ID         int    `json:"id"`
	Code       string `json:"code"`
	Desription string `json:"description"`
}

// Get All Supplier
func getSuppliers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	db, err := sql.Open(dbServer, connectString)
	errorHandler(w, err)
	defer db.Close()
	var supplier = Supplier{}
	var suppliers []Supplier
	query := `SELECT SupIdx, SupplierCode, SupplierName FROM MstrSupplier`
	result, err := db.Query(query)
	errorHandler(w, err)
	for result.Next() {
		err = result.Scan(&supplier.ID, &supplier.Code, &supplier.Desription)
		errorHandler(w, err)
		suppliers = append(suppliers, supplier)
	}
	json.NewEncoder(w).Encode(suppliers)
}

// Get Single Supplier
func getSupplier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var supplier Supplier
	json.NewEncoder(w).Encode(supplier)
}

// Create a New Supplier
func createSupplier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body, _ := ioutil.ReadAll(r.Body)
	json.NewEncoder(w).Encode(body)
}

// Update Supplier
func updateSupplier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body, _ := ioutil.ReadAll(r.Body)
	json.NewEncoder(w).Encode(body)
}
