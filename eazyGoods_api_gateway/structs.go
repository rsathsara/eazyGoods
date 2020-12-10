package main

// Modal Struct
type Modal struct {
	Data []Data
}

// Data Struct
type Data struct {
	APIResponse []APIResponse
}

// APIResponse Struct
type APIResponse struct {
	Status int    `json:"status"`
	Body   string `json:"body"`
}

// Services Struct
type Services struct {
	ID      int
	Name    string
	URL     string
	APIName string
}

// API Struct
type API struct {
	Name string
}
