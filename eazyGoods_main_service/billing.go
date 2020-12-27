package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	mssql "github.com/denisenkom/go-mssqldb"
)

// BillPage struct
type BillPage struct {
	ResponseMsg ResponseMsg `json:"responseMsg"`
	Bill        Bill        `json:"bill"`
}

// Bill struct
type Bill struct {
	BillHeader
	BillDetail []BillDetail `json:"billDetails"`
}

// BillHeader struct
type BillHeader struct {
	ID         int     `json:"id,omitempty" tvp:"id"`
	Date       string  `json:"date" tvp:"date"`
	BillNo     string  `json:"billNo" tvp:"-"`
	BillToID   int     `json:"billToId" tvp:"billToId"`
	BillTotal  float64 `json:"billTotal" tvp:"billTotal"`
	CashPaid   float64 `json:"cashPaid,omitempty" tvp:"-"`
	CreditPaid float64 `json:"creditPaid,omitempty" tvp:"-"`
}

// BillDetail struct
type BillDetail struct {
	ItemID           int     `json:"id" tvp:"itemId"`
	ItemCode         string  `json:"code" tvp:"-"`
	ItemDesctription string  `json:"description" tvp:"-"`
	UnitID           int     `json:"unitId" tvp:"-"`
	RecNo            int     `json:"recNo" tvp:"recNo"`
	Qty              float64 `json:"qty" tvp:"qty"`
	Price            float64 `json:"price" tvp:"price"`
	Value            float64 `json:"value" tvp:"value"`
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
	db, err := sql.Open(dbServer, connectString)
	errorHandler(w, err)
	defer db.Close()

	body, _ := ioutil.ReadAll(r.Body)
	var bill Bill
	json.Unmarshal(body, &bill)

	var billPage BillPage
	var billHdSlice []BillHeader
	billHdSlice = append(billHdSlice, bill.BillHeader)
	billHdType := mssql.TVP{
		TypeName: "BillHdTableType",
		Value:    billHdSlice,
	}
	billDtType := mssql.TVP{
		TypeName: "BillDtTableType",
		Value:    bill.BillDetail,
	}

	query := `exec sp_bill @BillHeader = $1, @BillDetail = $2, @Action = $3;`
	result := db.QueryRow(query, billHdType, billDtType, "save")
	err = result.Scan(&responseMsg.Status, &responseMsg.Msg)
	if err != nil {
		responseMsg.Status = "Error"
		responseMsg.Msg = string(err.Error())
	} else {
		billPage.ResponseMsg = responseMsg
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(billPage)
}

// Update Bill
func updateBill(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body, _ := ioutil.ReadAll(r.Body)
	json.NewEncoder(w).Encode(body)
}
