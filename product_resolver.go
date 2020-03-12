package main

import "github.com/graph-gophers/graphql-go"

type ProductResolver struct {
	p Product
}

func (pr *ProductResolver) Ean() graphql.ID {
	return graphql.ID(pr.p.EAN)
}

func (pr *ProductResolver) Name() string {
	return pr.p.Name
}

func (pr *ProductResolver) Price() float64 {
	return pr.p.Price
}

func (pr *ProductResolver) Image() string {
	return pr.p.Image
}

func (pr *ProductResolver) Department() string {
	return pr.p.Department
}

func (pr *ProductResolver) AgeRestriction() int32 {
	return pr.p.AgeRestriction
}

func (pr *ProductResolver) Tagged() bool {
	return pr.p.Tagged
}

func (pr *ProductResolver) QuantityRestriction() int32 {
	return pr.p.QuantityRestriction
}

func (pr *ProductResolver) Aisle() int32 {
	return pr.p.Aisle
}

func (pr *ProductResolver) Row() int32 {
	return pr.p.Row
}

func (pr *ProductResolver) Bay() int32 {
	return pr.p.Bay
}

func (pr *ProductResolver) Shelf() int32 {
	return pr.p.Shelf
}
