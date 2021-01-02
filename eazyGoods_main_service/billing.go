package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

// BillPage struct
type BillPage struct {
	ResponseMsg ResponseMsg `json:"responseMsg"`
	Bill        []Bill      `json:"bill"`
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
	UnitDescription  string  `json:"unitDescription" tvp:"-"`
	RecNo            int     `json:"recNo" tvp:"recNo"`
	Qty              float64 `json:"qty" tvp:"qty"`
	Price            float64 `json:"price" tvp:"price"`
	Value            float64 `json:"value" tvp:"value"`
}

// Get All Bills
func getBills(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	db, err := sql.Open(dbServer, connectString)
	errorHandler(w, err)
	defer db.Close()

	var billPage BillPage
	var bill = Bill{}
	var bills []Bill

	query := `SELECT InvHdIdx, InvNo, CONVERT(VARCHAR,InvDate,23) AS InvDate FROM TrnInvHeader`
	result, err := db.Query(query)
	errorHandler(w, err)
	for result.Next() {
		err = result.Scan(&bill.ID, &bill.BillNo, &bill.Date)
		errorHandler(w, err)
		bills = append(bills, bill)
	}

	billPage.Bill = bills
	json.NewEncoder(w).Encode(billPage)
}

// Get Single Bill
func getBill(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	db, err := sql.Open(dbServer, connectString)
	errorHandler(w, err)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	var billPage BillPage
	var bill = Bill{}
	var billDt = BillDetail{}
	var billDtList []BillDetail

	hdQuery := `SELECT hd.InvHdIdx, CONVERT(VARCHAR,hd.InvDate,23) AS InvDate, hd.InvNo, hd.CustIdx, hd.InvTotal
		FROM TrnInvHeader AS hd WHERE hd.InvHdIdx = $1`
	dtQuery := `
		DECLARE @recNo int = 0;
		SELECT SUM (@recNo+1) OVER (ORDER BY dt.InvDtIdx) AS recNo, dt.ItemIdx, itm.ItemCode, itm.ItemDescription, itm.UnitIdx, u.UnitDescription, 
		dt.InvQty, dt.SalePrice, dt.InvValue FROM TrnInvDetail AS dt
		LEFT JOIN MstrStkItem AS itm ON itm.ItemIdx = dt.ItemIdx
		LEFT JOIN MstrStkUnit AS u ON u.UnitIdx = itm.UnitIdx WHERE dt.InvHdIdx = $1`
	result := db.QueryRow(hdQuery, id)
	err = result.Scan(&bill.BillHeader.ID, &bill.BillHeader.Date, &bill.BillHeader.BillNo,
		&bill.BillHeader.BillToID, &bill.BillHeader.BillTotal)
	if err != nil {
		if err == sql.ErrNoRows {
			responseMsg.Status = "Warning"
			responseMsg.Msg = "No Bill Details Found"
		}
	} else {
		dtResult, err := db.Query(dtQuery, id)
		errorHandler(w, err)
		for dtResult.Next() {
			err = dtResult.Scan(&billDt.RecNo, &billDt.ItemID, &billDt.ItemCode, &billDt.ItemDesctription,
				&billDt.UnitID, &billDt.UnitDescription, &billDt.Qty, &billDt.Price, &billDt.Value)
			errorHandler(w, err)
			billDtList = append(billDtList, billDt)
		}
	}

	bill.BillDetail = billDtList
	billPage.Bill = append(billPage.Bill, bill)
	billPage.ResponseMsg = responseMsg
	json.NewEncoder(w).Encode(billPage)
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

	query := `exec sp_bill @BillHeader = $1, @BillDetail = $2, @docId = NULL, @Action = $3;`
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
	db, err := sql.Open(dbServer, connectString)
	errorHandler(w, err)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]
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

	query := `exec sp_bill @BillHeader = $1, @BillDetail = $2, @docId = $3, @Action = $4;`
	result := db.QueryRow(query, billHdType, billDtType, id, "edit")
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
