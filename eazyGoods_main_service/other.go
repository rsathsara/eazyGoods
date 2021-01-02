package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

// Get New Number
func getNewNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	db, err := sql.Open(dbServer, connectString)
	errorHandler(w, err)
	defer db.Close()

	vars := mux.Vars(r)
	docType := vars["docType"]

	var newNumber string
	query := `SELECT PrifixText+RIGHT(REPLICATE('0', RoundLength) + CONVERT(varchar(100),
		(ISNULL(DefaultNumber, 0)+1)), RoundLength) AS num FROM SysDefaultNumber WHERE DocType = $1`
	result := db.QueryRow(query, docType)
	err = result.Scan(&newNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			responseMsg.Status = "Warning"
			responseMsg.Msg = "New Number is Not Define Yet"
			json.NewEncoder(w).Encode(responseMsg)
			return
		}
	}
	responseMsg.Status = "Success"
	responseMsg.Msg = newNumber
	json.NewEncoder(w).Encode(responseMsg)
}

// Generate Report
func reportGenerator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	db, err := sql.Open(dbServer, connectString)
	errorHandler(w, err)
	defer db.Close()

	vars := mux.Vars(r)
	reportType := vars["id"]

	body, _ := ioutil.ReadAll(r.Body)
	var reportFilter ReportFilter
	json.Unmarshal(body, &reportFilter)

	var reportOutput = ReportOutput{}
	var report []ReportOutput

	dtResult, err := db.Query(`exec sp_reports @ReportId = $1, @ReportStartDate = $2, @ReportEnd = $3;`,
		reportType, reportFilter.StartDate, reportFilter.EndDate)
	errorHandler(w, err)
	for dtResult.Next() {
		err = dtResult.Scan()
		errorHandler(w, err)
		report = append(report, reportOutput)
	}
	json.NewEncoder(w).Encode(report)
}
