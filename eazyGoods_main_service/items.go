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

// ItemPage Struct
type ItemPage struct {
	ResponseMsg ResponseMsg `json:"responseMsg"`
	ItemList    []Item      `json:"itemList"`
	ItemDetails Item        `json:"itemDetails"`
}

// Item struct
type Item struct {
	ID          int     `json:"id" tvp:"id"`
	Code        string  `json:"code" tvp:"code"`
	Description string  `json:"description" tvp:"description"`
	ItemCat1Id  int     `json:"itemCat1Id" tvp:"itemCat1Id"`
	ItemCat2Id  int     `json:"itemCat2Id" tvp:"itemCat2Id"`
	Unit        Unit    `json:"unit"`
	SOH         float64 `json:"soh"`
	SalePrice   float64 `json:"salePrice" tvp:"salePrice"`
}

// InsertItem struct
type InsertItem struct {
	ID          int     `json:"id" tvp:"id"`
	Code        string  `json:"code" tvp:"code"`
	Description string  `json:"description" tvp:"description"`
	ItemCat1Id  int     `json:"itemCat1Id" tvp:"itemCat1Id"`
	ItemCat2Id  int     `json:"itemCat2Id" tvp:"itemCat2Id"`
	UnitID      int     `json:"unitId" tvp:"unitId"`
	SalePrice   float64 `json:"salePrice" tvp:"salePrice"`
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
		err = result.Scan(&item.ID, &item.Code, &item.Description)
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

	query := `SELECT itm.ItemIdx, itm.ItemCode, itm.ItemDescription, ISNULL(itm.SalesPrice, 0.000) AS SalesPrice, 
		u.UnitIdx, u.UnitDescription, itm.ItemCat1Idx, itm.ItemCat2Idx FROM MstrStkItem AS itm
		LEFT JOIN MstrStkUnit AS u ON u.UnitIdx = itm.UnitIdx WHERE ItemIdx = $1`
	result := db.QueryRow(query, id)
	err = result.Scan(&item.ID, &item.Code, &item.Description, &item.SalePrice, &unit.ID, &unit.Description, &item.ItemCat1Id, &item.ItemCat2Id)
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
	db, err := sql.Open(dbServer, connectString)
	errorHandler(w, err)
	defer db.Close()

	body, _ := ioutil.ReadAll(r.Body)
	var item InsertItem
	json.Unmarshal(body, &item)

	var itemPage ItemPage

	var itemSlice []InsertItem
	itemSlice = append(itemSlice, item)
	itemType := mssql.TVP{
		TypeName: "ItemTableType",
		Value:    itemSlice,
	}

	query := `exec sp_item @ItemDetails = $1, @docId = NULL, @Action = $2;`
	result := db.QueryRow(query, itemType, "save")
	err = result.Scan(&responseMsg.Status, &responseMsg.Msg)
	if err != nil {
		responseMsg.Status = "Error"
		responseMsg.Msg = string(err.Error())
	}

	itemPage.ResponseMsg = responseMsg
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(itemPage)
}

// Update Item
func updateItem(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbServer, connectString)
	errorHandler(w, err)
	defer db.Close()

	body, _ := ioutil.ReadAll(r.Body)
	var item InsertItem
	json.Unmarshal(body, &item)

	var itemPage ItemPage

	var itemSlice []InsertItem
	itemSlice = append(itemSlice, item)
	itemType := mssql.TVP{
		TypeName: "ItemTableType",
		Value:    itemSlice,
	}

	query := `exec sp_item @ItemDetails = $1, @docId = NULL, @Action = $2;`
	result := db.QueryRow(query, itemType, "edit")
	err = result.Scan(&responseMsg.Status, &responseMsg.Msg)
	if err != nil {
		responseMsg.Status = "Error"
		responseMsg.Msg = string(err.Error())
	}

	itemPage.ResponseMsg = responseMsg
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(itemPage)
}
