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
	user := r.authService.GetUserFromContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	params := order.AddParams{
		ID:       uint64(input.ID),
		Quantity: input.Quantity,
	}
	basket, err := r.basketService.Add(user.ID, params)
	if err != nil {
		return nil, err
	}
	var items []*model.BasketItem
	for _, item := range basket.Items {
		it := &model.BasketItem{
			ProductID:    int(item.ProductID),
			ProductTitle: item.ProductTitle,
			Quantity:     item.Quantity,
			Price:        item.Price,
		}
		items = append(items, it)

	}
	return &model.Basket{
		Items: items,
		Total: basket.Total,
	}, nil
}

func (r *mutationResolver) RemoveFromBasket(ctx context.Context, input model.RemoveFromBasketInput) (*model.Basket, error) {
	user := r.authService.GetUserFromContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	params := order.RemoveParams{
		ID: uint64(input.ID),
	}
	basket, err := r.basketService.Remove(user.ID, params)
	if err != nil {
		return nil, err
	}
	var items []*model.BasketItem
	total := float64(0)
	if basket != nil {
		for _, item := range basket.Items {
			it := &model.BasketItem{
				ProductID:    int(item.ProductID),
				ProductTitle: item.ProductTitle,
				Quantity:     item.Quantity,
				Price:        item.Price,
			}
			items = append(items, it)
			total = basket.Total
		}

	}
	return &model.Basket{
		Items: items,
		Total: total,
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
			ID:    int(p.ID),
			Name:  p.Name,
			Price: p.Price,
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *queryResolver) Basket(ctx context.Context) (*model.Basket, error) {
	user := r.authService.GetUserFromContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	basket, err := r.basketService.Get(user.ID)
	if err != nil {
		return nil, err
	}
	var items []*model.BasketItem
	total := float64(0)
	if basket != nil {
		for _, item := range basket.Items {
			it := &model.BasketItem{
				ProductID:    int(item.ProductID),
				ProductTitle: item.ProductTitle,
				Quantity:     item.Quantity,
				Price:        item.Price,
			}
			items = append(items, it)
			total = basket.Total
		}

	}
	return &model.Basket{
		Items: items,
		Total: total,
	}, nil
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
