package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type ProductRepository struct {
	p Products
}

func NewProductRepository() (ProductRepository, error) {
	f, err := os.Open("data/products_1.json")
	if err != nil {
		return ProductRepository{}, err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return ProductRepository{}, err
	}

	var products Products
	err = json.Unmarshal(data, &products)
	if err != nil {
		return ProductRepository{}, err
	}

	return ProductRepository{products}, nil
}

func (pr ProductRepository) getByEAN(ean string) (Product, error) {
	for _, p := range pr.p.Products {
		if p.EAN == ean {
			return p, nil
		}
	}

	return Product{}, errors.New("no product found")
}

func (pr ProductRepository) findProduct(filter *ProductFilter) []Product {
	p := []Product{}
	for _, prod := range pr.p.Products {
		if filter == nil {
			p = append(p, prod)
			continue
		}
		if filter.Tagged != nil && *filter.Tagged != prod.Tagged {
			continue
		}

		if filter.PriceLessThan != nil && prod.Price >= *filter.PriceLessThan {
			continue
		}

		p = append(p, prod)
	}

	return p
}
