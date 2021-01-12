package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

// OtherPage Struct
type OtherPage struct {
	ResponseMsg  ResponseMsg `json:"responseMsg"`
	UnitList     []Unit      `json:"unitList"`
	ItemCat1List []ItemCat1  `json:"itemCat1List"`
	ItemCat2List []ItemCat2  `json:"itemCat2List"`
}

// Unit Struct
type Unit struct {
	ID          int    `json:"id" tvp:"unitId"`
	Description string `json:"description" tvp:"-"`
}

// ItemCat1 Struct
type ItemCat1 struct {
	ID          int    `json:"id" tvp:"ItemCat1Idx"`
	Description string `json:"description" tvp:"ItemCategory1"`
}

// ItemCat2 Struct
type ItemCat2 struct {
	ID          int    `json:"id" tvp:"ItemCat2Idx"`
	Description string `json:"description" tvp:"ItemCategory2"`
}

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

// Get Session Details
func getSessionDetails() SessionDetails {
	response, err := http.Get("http://localhost:3240/sessionDetails")
	if err != nil {
		fmt.Print(err.Error())
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var sessionDetails SessionDetails
	json.Unmarshal(responseData, &sessionDetails)
	fmt.Println("UserName : " + sessionDetails.Username + "Name : " + sessionDetails.Name)
	return sessionDetails
}

// Get Unit List
func getUnits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	db, err := sql.Open(dbServer, connectString)
	errorHandler(w, err)
	defer db.Close()

	var otherPage OtherPage
	var unit = Unit{}
	var units []Unit

	query := `SELECT UnitIdx, UnitDescription FROM MstrStkUnit`
	result, err := db.Query(query)
	errorHandler(w, err)
	for result.Next() {
		err = result.Scan(&unit.ID, &unit.Description)
		errorHandler(w, err)
		units = append(units, unit)
	}

	otherPage.UnitList = units
	json.NewEncoder(w).Encode(otherPage)
}

// Get Item Category 1 List
func getItemCat1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	db, err := sql.Open(dbServer, connectString)
	errorHandler(w, err)
	defer db.Close()

	var otherPage OtherPage
	var obj = ItemCat1{}
	var list []ItemCat1

	query := `SELECT ItemCat1Idx, ItemCategory1 FROM MstrStkItemCat1`
	result, err := db.Query(query)
	errorHandler(w, err)
	for result.Next() {
		err = result.Scan(&obj.ID, &obj.Description)
		errorHandler(w, err)
		list = append(list, obj)
	}

	otherPage.ItemCat1List = list
	json.NewEncoder(w).Encode(otherPage)
}

// Get Item Category 2 List
func getItemCat2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	db, err := sql.Open(dbServer, connectString)
	errorHandler(w, err)
	defer db.Close()

	var otherPage OtherPage
	var obj = ItemCat2{}
	var list []ItemCat2

	query := `SELECT ItemCat2Idx, ItemCategory2 FROM MstrStkItemCat2`
	result, err := db.Query(query)
	errorHandler(w, err)
	for result.Next() {
		err = result.Scan(&obj.ID, &obj.Description)
		errorHandler(w, err)
		list = append(list, obj)
	}

	otherPage.ItemCat2List = list
	json.NewEncoder(w).Encode(otherPage)
}
