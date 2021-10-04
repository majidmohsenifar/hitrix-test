package graph

//go:generate go run github.com/99designs/gqlgen
import "hitrix-test/internal/product"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	productService *product.Service
}

func NewResolver(productService *product.Service) *Resolver {
	return &Resolver{
		productService: productService,
	}
}
