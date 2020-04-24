package model

// DescriptionValueData represents the description and value data coming from smartmei
type DescriptionValueData struct {
	Description string  `json:"description"`
	Value       float64 `json:"value"`
}
