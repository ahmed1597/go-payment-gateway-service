package models

// Gateway represents the payment gateway details stored in the database
type Gateway struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Priority   int    `json:"priority"`
	DataFormat string `json:"data_format"`
	Protocol   string `json:"protocol"`
}
