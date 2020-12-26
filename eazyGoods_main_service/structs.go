package main

// Modal struct
type Modal struct {
	Data []Data
}

//Data struct
type Data struct {
	Bill        []Bill
	Item        []Item
	Unit        []Unit
	ResponseMsg []ResponseMsg
	Customer    []Customer
}

// Unit Struct
type Unit struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

// ResponseMsg Struct
type ResponseMsg struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
	Type   string `json:"type"`
}
