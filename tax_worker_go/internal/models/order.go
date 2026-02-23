package models

type Order struct {
	OrderID   string  `json:"order_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Subtotal  float64 `json:"subtotal"`
	Timestamp string  `json:"timestamp"`
}

type ProcessedOrder struct {
	OrderID          string
	Latitude         float64
	Longitude        float64
	Subtotal         float64
	Timestamp        string
	CompositeTaxRate float64
	TaxAmount        float64
	TotalAmount      float64
	Jurisdictions    []string
	Breakdown        TaxBreakdown
}

type TaxBreakdown struct {
	StateRate    float64
	CountyRate   float64
	CityRate     float64
	SpecialRates float64
}
