package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	generated1 "hitrix-test/graph/generated"
	"hitrix-test/graph/model"
	"hitrix-test/internal/auth"
	"hitrix-test/internal/order"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, input model.RegisterInput) (*model.User, error) {
	params := auth.RegisterParams{
		Email:    input.Email,
		Password: input.Password,
	}
	err := r.authService.Register(params)
	if err != nil {
		return nil, err
	}
	return &model.User{
		Email: input.Email,
	}, nil
}

func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginInput) (*model.Login, error) {
	params := auth.LoginParams{
		Email:    input.Email,
		Password: input.Password,
	}
	accessToken, refreshToken, err := r.authService.Login(params)
	if err != nil {
		return nil, err

	}
	return &model.Login{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (r *mutationResolver) AddToBasket(ctx context.Context, input model.AddToBasketInput) (*model.Basket, error) {
	if user := r.authService.GetUserFromContext(ctx); user == nil {
		return nil, fmt.Errorf("access denied")
	}
	params := order.AddParams{
		ID:       input.ID,
		Quantity: input.Quantity,
	}
	r.basketService.Add(params)
	//TODO handle this later
	return &model.Basket{}, nil
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

func (r *queryResolver) Basket(ctx context.Context) (*model.Basket, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
