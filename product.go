package main

type Products struct {
	Products []Product `json:"products"`
}

type Product struct {
	EAN                 string  `json:"ean"`
	Name                string  `json:"name"`
	Price               float64 `json:"price"`
	Image               string  `json:"image"`
	Department          string  `json:"department"`
	AgeRestriction      int32   `json:"age_restriction"`
	Tagged              bool    `json:"tagged"`
	QuantityRestriction int32   `json:"quantity_restriction"`
	Aisle               int32   `json:"aisle"`
	Row                 int32   `json:"row"`
	Bay                 int32   `json:"bay"`
	Shelf               int32   `json:"shelf"`
}
