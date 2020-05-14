package airtime

// Response is the response from the api
type Response struct {
	NumSent       int     `json:"numSent"`
	TotalAmount   string  `json:"totalAmount"`
	TotalDiscount string  `json:"totalDiscount"`
	ErrorMessage  string  `json:"errorMessage"`
	Responses     []Entry `json:"responses"`
}

// Entry is the entry for each airtime response
type Entry struct {
	ErrorMessage string `json:"errorMessage"`
	PhoneNumber  string `json:"phoneNumber"`
	Amount       string `json:"amount"`
	Discount     string `json:"discount"`
	Status       string `json:"status"`
	RequestID    string `json:"requestId"`
}
