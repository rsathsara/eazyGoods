package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

// GrnPage struct
type GrnPage struct {
	ResponseMsg ResponseMsg `json:"responseMsg"`
	Grn         []Grn       `json:"grn"`
}

// Grn struct
type Grn struct {
	GrnHeader
	GrnDetail []GrnDetail `json:"grnDetails"`
}

// GrnHeader struct
type GrnHeader struct {
	ID        int     `json:"id,omitempty" tvp:"id"`
	Date      string  `json:"date" tvp:"date"`
	GrnNo     string  `json:"grnNo" tvp:"-"`
	GrnFromID int     `json:"grnFromId" tvp:"SupIdx"`
	GrnTotal  float64 `json:"grnTotal" tvp:"GrnGTot"`
}

// GrnDetail struct
type GrnDetail struct {
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

// Get All GRNs
func getGrns(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	db, err := sql.Open(dbServer, connectString)
	errorHandler(w, err)
	defer db.Close()

	var grnPage GrnPage
	var grn = Grn{}
	var grns []Grn

	query := `SELECT GrnHdIdx, GrnNo, CONVERT(VARCHAR,GrnDate,23) AS GrnDate FROM TrnGrnHeader`
	result, err := db.Query(query)
	errorHandler(w, err)
	for result.Next() {
		err = result.Scan(&grn.ID, &grn.GrnNo, &grn.Date)
		errorHandler(w, err)
		grns = append(grns, grn)
	}

	grnPage.Grn = grns
	json.NewEncoder(w).Encode(grnPage)
}

// Get Single GRN
func getGrn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	db, err := sql.Open(dbServer, connectString)
	errorHandler(w, err)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	var grnPage GrnPage
	var grn = Grn{}
	var grnDt = GrnDetail{}
	var grnDtList []GrnDetail

	hdQuery := `SELECT hd.GrnHdIdx, CONVERT(VARCHAR,hd.GrnDate,23) AS GrnDate, hd.GrnNo, hd.SupIdx, hd.GrnGTot
	FROM TrnGrnHeader AS hd WHERE hd.GrnHdIdx = $1`
	dtQuery := `
		DECLARE @recNo int = 0;
		SELECT SUM (@recNo+1) OVER (ORDER BY dt.GrnDtIdx) AS recNo, dt.ItemIdx, itm.ItemCode, itm.ItemDescription, itm.UnitIdx, 
		u.UnitDescription, dt.GrnQty, dt.GrnPrice, dt.GrnValue FROM TrnGrnDetail AS dt
		LEFT JOIN MstrStkItem AS itm ON itm.ItemIdx = dt.ItemIdx
		LEFT JOIN MstrStkUnit AS u ON u.UnitIdx = itm.UnitIdx WHERE dt.GrnHdIdx = $1`
	result := db.QueryRow(hdQuery, id)
	err = result.Scan(&grn.ID, &grn.Date, &grn.GrnNo, &grn.GrnFromID, &grn.GrnTotal)

	if err != nil {
		if err == sql.ErrNoRows {
			responseMsg.Status = "Warning"
			responseMsg.Msg = "No GRN Details Found"
		}
	} else {
		dtResult, err := db.Query(dtQuery, id)
		errorHandler(w, err)
		for dtResult.Next() {
			err = dtResult.Scan(&grnDt.RecNo, &grnDt.ItemID, &grnDt.ItemCode, &grnDt.ItemDesctription,
				&grnDt.UnitID, &grnDt.UnitDescription, &grnDt.Qty, &grnDt.Price, &grnDt.Value)
			errorHandler(w, err)
			grnDtList = append(grnDtList, grnDt)
		}
	}

	grn.GrnDetail = grnDtList
	grnPage.Grn = append(grnPage.Grn, grn)
	grnPage.ResponseMsg = responseMsg
	json.NewEncoder(w).Encode(grnPage)
}

// Create a New GRN
func createGrn(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbServer, connectString)
	errorHandler(w, err)
	defer db.Close()

	body, _ := ioutil.ReadAll(r.Body)
	var grn Grn
	json.Unmarshal(body, &grn)

	var grnPage GrnPage
	var grnHdSlice []GrnHeader
	grnHdSlice = append(grnHdSlice, grn.GrnHeader)
	grnHdType := mssql.TVP{
		TypeName: "GrnHdTableType",
		Value:    grnHdSlice,
	}
	grnDtType := mssql.TVP{
		TypeName: "GrnDtTableType",
		Value:    grn.GrnDetail,
	}

	query := `exec sp_grn @GrnHeader = $1, @GrnDetail = $2, @docId = NULL, @Action = $3;`
	result := db.QueryRow(query, grnHdType, grnDtType, "save")
	err = result.Scan(&responseMsg.Status, &responseMsg.Msg)
	if err != nil {
		responseMsg.Status = "Error"
		responseMsg.Msg = string(err.Error())
	}

	grnPage.ResponseMsg = responseMsg

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(grnPage)
}

// Update GRN
func updateGrn(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbServer, connectString)
	errorHandler(w, err)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	body, _ := ioutil.ReadAll(r.Body)
	var grn Grn
	json.Unmarshal(body, &grn)

	var grnPage GrnPage
	var grnHdSlice []GrnHeader
	grnHdSlice = append(grnHdSlice, grn.GrnHeader)
	grnHdType := mssql.TVP{
		TypeName: "GrnHdTableType",
		Value:    grnHdSlice,
	}
	grnDtType := mssql.TVP{
		TypeName: "GrnDtTableType",
		Value:    grn.GrnDetail,
	}

	query := `exec sp_grn @GrnHeader = $1, @GrnDetail = $2, @docId = $3, @Action = $4;`
	result := db.QueryRow(query, grnHdType, grnDtType, id, "edit")
	err = result.Scan(&responseMsg.Status, &responseMsg.Msg)
	if err != nil {
		responseMsg.Status = "Error"
		responseMsg.Msg = string(err.Error())
	}

	grnPage.ResponseMsg = responseMsg
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(grnPage)
}
