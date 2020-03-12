package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct {
	ProductRepo ProductRepository
}

type ProductFilter struct {
	PriceLessThan *float64
	Tagged        *bool
}

func (q *query) Products(args struct{ Filter *ProductFilter }) ([]*ProductResolver, error) {
	pr := []*ProductResolver{}

	for _, p := range q.ProductRepo.findProduct(args.Filter) {
		pr = append(pr, &ProductResolver{p})
	}

	return pr, nil
}

func (q *query) Product(args struct{ EAN string }) (*ProductResolver, error) {
	p, err := q.ProductRepo.getByEAN(args.EAN)

	if err != nil {
		return nil, err
	}

	return &ProductResolver{p}, nil
}

var products Products

func main() {

	repo, err := NewProductRepository()

	if err != nil {
		log.Fatal(err)
	}

	s, _ := os.Open("schema.graphql")
	b, _ := ioutil.ReadAll(s)

	schema := graphql.MustParseSchema(string(b), &query{repo})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Println("Serving GQL Service")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
