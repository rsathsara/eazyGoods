package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

// DbOutput Map
type DbOutput map[string]interface{}

// ReportPage Struct
type ReportPage struct {
	ResponseMsg  ResponseMsg  `json:"responseMsg"`
	ReportList   []ReportList `json:"reportList"`
	ReportOutput ReportOutput `json:"reportOutput"`
	ReportTitle  string       `json:"reportTitle"`
}

// ReportList Struct
type ReportList struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

// ReportFilter Struct
type ReportFilter struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

// ReportOutput Struct
type ReportOutput struct {
	List    []DbOutput `json:"list"`
	Columns []string   `json:"columns"`
}

var reportList = []ReportList{
	{ID: 1, Description: "GRN List"},
	{ID: 2, Description: "Invoice List"},
	{ID: 3, Description: "Stock Balance Report"},
}

// Generate Report
func reportGenerator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	db, err := sql.Open(dbServer, connectString)
	errorHandler(w, err)
	defer db.Close()

	vars := mux.Vars(r)
	reportType := vars["id"]

	var reportPage ReportPage
	var reportFilter ReportFilter
	var reportOutput ReportOutput

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &reportFilter)

	result, err := db.Query(`exec sp_reports @ReportId = $1, @ReportStartDate = $2, @ReportEndDate = $3;`,
		reportType, reportFilter.StartDate, reportFilter.EndDate)
	errorHandler(w, err)

	columns, _ := result.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	var outputList []DbOutput

	for result.Next() {
		var output = DbOutput{}

		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		result.Scan(valuePtrs...)

		for i, col := range columns {
			val := values[i]

			b, ok := val.([]byte)
			var v interface{}
			if ok {
				v = string(b)
			} else {
				v = val
			}
			output[col] = v
		}

		outputList = append(outputList, output)
	}

	reportTypeID, _ := strconv.Atoi(reportType)
	for _, v := range reportList {
		if v.ID == reportTypeID {
			reportPage.ReportTitle = v.Description
			break
		}
	}

	reportOutput.Columns = columns
	reportOutput.List = outputList
	reportPage.ReportOutput = reportOutput
	json.NewEncoder(w).Encode(reportPage)
}

// Get ReportList
func getReports(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var reportPage ReportPage
	reportPage.ReportList = reportList
	json.NewEncoder(w).Encode(reportPage)
}
