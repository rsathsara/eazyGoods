package main

import (
	"encoding/json"
	"net/http"
)

// Get All Bills
func getBills(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var bills []Bill
	bills = []Bill{
		Bill{ID: 1, Date: "2020/12/08", BillNo: "INV2020/0001", BillTo: "Cash"},
		Bill{ID: 2, Date: "2020/12/08", BillNo: "INV2020/0002", BillTo: "Cash"},
		Bill{ID: 3, Date: "2020/12/08", BillNo: "INV2020/0003", BillTo: "Cash"},
		Bill{ID: 4, Date: "2020/12/08", BillNo: "INV2020/0004", BillTo: "Cash"},
	}
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
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

// Update Bill
func updatetBill(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}
