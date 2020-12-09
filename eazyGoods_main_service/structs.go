package main

// Modal struct
type Modal struct {
	Data []Data
}

var modal = Modal{}

//Data struct
type Data struct {
	Bill        []Bill
	Item        []Item
	Unit        []Unit
	ResponseMsg []ResponseMsg
}

// Bill struct
type Bill struct {
	ID     int    `json:"id"`
	Date   string `json:"date"`
	BillNo string `json:"billNo"`
	BillTo string `json:"billTo"`
	Items  []Item `json:"items"`
}

// Item struct
type Item struct {
	ID         int     `json:"id"`
	Code       string  `json:"code"`
	Desription string  `json:"description"`
	Unit       *Unit   `json:"unit"`
	SOH        float64 `json:"soh"`
	Qty        float64 `json:"qty"`
	Price      float64 `json:"price"`
	Value      float64 `json:"value"`
}

// Unit Struct
type Unit struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

// ResponseMsg Struct
type ResponseMsg struct {
	Status []string `json:"status"`
	Msg    []string `json:"msg"`
}
