package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

// ItemPage Struct
type ItemPage struct {
	ResponseMsg ResponseMsg `json:"responseMsg"`
	ItemList    []Item      `json:"itemList"`
	ItemDetails Item        `json:"itemDetails"`
}

// Item struct
type Item struct {
	ID         int     `json:"id"`
	Code       string  `json:"code"`
	Desription string  `json:"description"`
	Unit       Unit    `json:"unit"`
	SOH        float64 `json:"soh"`
}

// Get All Items
func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	db, err := sql.Open(dbServer, connectString)
	errorHandler(w, err)
	defer db.Close()

	var itemPage ItemPage
	var item = Item{}
	var items []Item

	query := `SELECT ItemIdx, ItemCode, ItemDescription FROM MstrStkItem`
	result, err := db.Query(query)
	errorHandler(w, err)
	for result.Next() {
		err = result.Scan(&item.ID, &item.Code, &item.Desription)
		errorHandler(w, err)
		items = append(items, item)
	}

	itemPage.ItemList = items
	json.NewEncoder(w).Encode(itemPage)
}

// Get Single Item
func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	db, err := sql.Open(dbServer, connectString)
	errorHandler(w, err)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	var itemPage ItemPage
	var item Item
	var unit Unit

	query := `SELECT itm.ItemIdx, itm.ItemCode, itm.ItemDescription, u.UnitIdx, u.UnitDescription FROM MstrStkItem AS itm
		LEFT JOIN MstrStkUnit AS u ON u.UnitIdx = itm.UnitIdx WHERE ItemIdx = $1`
	result := db.QueryRow(query, id)
	err = result.Scan(&item.ID, &item.Code, &item.Desription, &unit.ID, &unit.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			responseMsg.Status = "Warning"
			responseMsg.Msg = "No Item Details Found"
		}
	} else {
		responseMsg.Status = "Success"
		responseMsg.Msg = ""
	}
	item.Unit = unit
	itemPage.ItemDetails = item
	itemPage.ResponseMsg = responseMsg
	json.NewEncoder(w).Encode(itemPage)
}

// Create a New Item
func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body, _ := ioutil.ReadAll(r.Body)
	json.NewEncoder(w).Encode(body)
}

// Update Item
func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body, _ := ioutil.ReadAll(r.Body)
	json.NewEncoder(w).Encode(body)
}

// New Item No
func newItemNo() {

}
