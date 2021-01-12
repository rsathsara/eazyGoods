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

// ResponseMsg Struct
type ResponseMsg struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
	Type   string `json:"type"`
}

// SessionDetails Struct
type SessionDetails struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}
