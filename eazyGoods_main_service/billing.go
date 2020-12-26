package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
)

// BillPage struct
type BillPage struct {
	ResponseMsg ResponseMsg `json:"responseMsg"`
	Bill        Bill        `json:"bill"`
}

// Bill struct
type Bill struct {
	ID         int          `json:"id,omitempty"`
	Date       string       `json:"date"`
	BillNo     string       `json:"billNo"`
	BillTo     string       `json:"billTo"`
	BillDetail []BillDetail `json:"billDetails"`
	BillTotal  float64      `json:"billTotal"`
	CashPaid   float64      `json:"cashPaid,omitempty"`
	CreditPaid float64      `json:"creditPaid,omitempty"`
}

// BillDetail struct
type BillDetail struct {
	Item
	RecNo int     `json:"recNo"`
	Qty   float64 `json:"qty"`
	Price float64 `json:"price"`
	Value float64 `json:"value"`
}

// Get All Bills
func getBills(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var bills []Bill
	json.NewEncoder(w).Encode(bills)
}

// Get Single Bill
func getBill(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var bill Bill
	json.NewEncoder(w).Encode(bill)
}

// Create a New Bill
func createBill(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json; charset=utf-8")
	db, err := sql.Open(dbServer, connectString)
	errorHandler(w, err)
	defer db.Close()

	body, _ := ioutil.ReadAll(r.Body)
	var bill Bill
	json.Unmarshal(body, &bill)

	dataSet, _ := json.Marshal(bill)
	row, err := db.Query(`EXEC [sp_bill] @BillData = ?, @Action = ?;`, dataSet, "save")
	if err != nil {
		fmt.Println(err)
	} else {
		for row.Next() {
			var firstname string
			row.Scan(&firstname)
			fmt.Println(firstname)
		}
	}

	// fmt.Printf("data : %+v", bill)
}

// Update Bill
func updateBill(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body, _ := ioutil.ReadAll(r.Body)
	json.NewEncoder(w).Encode(body)
}
