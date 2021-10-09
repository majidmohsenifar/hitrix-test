package graph

//go:generate go run github.com/99designs/gqlgen
import (
	"hitrix-test/internal/auth"
	"hitrix-test/internal/order"
	"hitrix-test/internal/product"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	productService *product.Service
	authService    *auth.Service
	basketService  *order.BasketService
}

func NewResolver(productService *product.Service, authService *auth.Service, basketService *order.BasketService) *Resolver {
	return &Resolver{
		productService: productService,
		authService:    authService,
		basketService:  basketService,
	}
}
