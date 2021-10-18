package order

import (
	"encoding/json"
	"fmt"
	"hitrix-test/internal/product"

	"github.com/latolukasz/beeorm"
)

type BasketService struct {
	ormEngine      *beeorm.Engine
	productService *product.Service
}

type AddParams struct {
	ID       uint64 `json="id"`
	Quantity int    `json="quantity"`
}

type RemoveParams struct {
	ID uint64 `json="id"`
}

type BasketItem struct {
	ProductID    uint64  `json:"productId"`
	ProductTitle string  `json:"productTitle"`
	Quantity     int     `json:"quantity"`
	Price        float64 `json:"price"`
}
type Basket struct {
	Items []BasketItem `json:"items"`
	Total float64      `json:"total"`
}

func (s *BasketService) Add(userID uint64, params AddParams) (*Basket, error) {
	if params.Quantity <= 0 {
		return nil, fmt.Errorf("quantity must be positive")
	}
	product, err := s.productService.FindByID(params.ID)
	if err != nil {
		return nil, err
	}

	basket, err := s.get(userID)
	if err != nil {
		return nil, err
	}
	if basket != nil {
		found := false
		items := basket.Items
		for i, item := range items {
			if item.ProductID == params.ID {
				found = true
				basket.Items[i].Quantity = params.Quantity
			}
		}
		if !found {
			item := BasketItem{
				ProductID:    params.ID,
				ProductTitle: product.Name,
				Quantity:     params.Quantity,
				Price:        product.Price,
			}
			basket.Items = append(basket.Items, item)
		}
		s.calculateTotal(basket)
		err := s.set(userID, basket)
		if err != nil {
			return nil, err
		}
		return s.Get(userID)
	}
	item := BasketItem{
		ProductID:    params.ID,
		ProductTitle: product.Name,
		Quantity:     params.Quantity,
		Price:        product.Price,
	}
	basket = &Basket{
		Items: []BasketItem{item},
		Total: 0.0,
	}
	s.calculateTotal(basket)
	err = s.set(userID, basket)
	if err != nil {
		return nil, err
	}
	return s.Get(userID)
}

func (s *BasketService) Get(userID uint64) (*Basket, error) {
	basket, err := s.get(userID)
	if err != nil {
		return nil, err
	}
	return basket, nil
}

func (s *BasketService) get(userID uint64) (*Basket, error) {
	basket := &Basket{}
	key := s.getKey(userID)
	redis := s.ormEngine.GetRedis()
	data, ok := redis.Get(key)
	if !ok {
		return nil, nil
	}
	err := json.Unmarshal([]byte(data), basket)
	return basket, err
}

func (s *BasketService) set(userID uint64, basket *Basket) error {
	key := s.getKey(userID)
	redis := s.ormEngine.GetRedis()
	data, err := json.Marshal(basket)
	if err != nil {
		return err
	}
	redis.Set(key, string(data), 0)

	return nil
}

func (s *BasketService) Remove(userID uint64, params RemoveParams) (*Basket, error) {
	basket, err := s.get(userID)
	if err != nil {
		return nil, err
	}
	if basket != nil {
		found := false
		items := basket.Items
		for i, item := range items {
			if item.ProductID == params.ID {
				found = true
				basket.Items = append(basket.Items[:i], basket.Items[i+1:]...)
				break
			}
		}
		if found {
			s.calculateTotal(basket)
			err := s.set(userID, basket)
			if err != nil {
				return nil, err
			}
		}
	}
	return s.Get(userID)
}

func (s *BasketService) getKey(userID uint64) string {
	return fmt.Sprintf("basket:%d", userID)
}

func (s *BasketService) calculateTotal(basket *Basket) {
	var total float64
	for _, item := range basket.Items {
		total += float64(item.Quantity) * item.Price

	}
	basket.Total = total
}
func NewBasketService(ormEngine *beeorm.Engine, productService *product.Service) *BasketService {
	return &BasketService{
		ormEngine:      ormEngine,
		productService: productService,
	}
}
