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

// ReportFilter Struct
type ReportFilter struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

// ReportOutput Struct
type ReportOutput struct {
	Col01 string `json:"col01"`
	Col02 string `json:"col02"`
	Col03 string `json:"col03"`
	Col04 string `json:"col04"`
	Col05 string `json:"col05"`
	Col06 string `json:"col06"`
	Col07 string `json:"col07"`
	Col08 string `json:"col08"`
	Col09 string `json:"col09"`
	Col10 string `json:"col10"`
}
