package models

type Product struct {
	Name     string
	Category string
	Supplier string
	Quantity uint64
	Price    float64
	Stock    uint64
	Unit     string
}
