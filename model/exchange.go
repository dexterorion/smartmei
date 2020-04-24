package model

// ExchangeResponse represents the exchange response model
type ExchangeResponse struct {
	Base  string             `json:"base"`
	Data  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}
