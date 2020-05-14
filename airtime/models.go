package airtime

// Response is the reponse from the api
type Response struct {
	NumSent       int
	TotalAmount   string
	TotalDiscount string
	ErrorMessage  string
	Responses     []Entry
}

// Entry is the entry for each airtime response
type Entry struct {
	ErrorMessage string
	PhoneNumber  string
	Amount       string
	Discount     string
	Status       string
	RequestID    string
}
