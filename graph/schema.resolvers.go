package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	generated1 "hitrix-test/graph/generated"
	"hitrix-test/graph/model"
)

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	err := r.productService.Create(input.Name, input.Price)
	if err != nil {
		return nil, err
	}
	return &model.Product{
		ID:    "someId", //TODO handle this later
		Name:  input.Name,
		Price: input.Price,
	}, nil
}

func (r *queryResolver) Products(ctx context.Context, input *model.ProductListInput) ([]*model.Product, error) {
	name := ""
	minPrice := 0.0
	maxPrice := 1000000.0
	if input != nil {
		name = *input.Name
		minPrice = *input.MinPrice
		maxPrice = *input.MaxPrice

	}
	result, err := r.productService.List(name, minPrice, maxPrice)
	if err != nil {
		return nil, err
	}
	var products []*model.Product
	for _, p := range result {
		product := &model.Product{
			ID:    "", //TODO handle this later
			Name:  p.Name,
			Price: p.Price,
		}
		products = append(products, product)
	}
	return products, nil
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
