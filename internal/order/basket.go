package order

type BasketService struct {
}

type AddParams struct {
	ID       int `json="id"`
	Quantity int `json="quantity"`
}

type RemoveParams struct {
	ID int `json="id"`
}

type BasketItem struct {
	ProductID    int
	ProductTitle string
	Quantity     int
	Price        float64
}
type Basket struct {
	Items []BasketItem
	Total float64
}

func (s *BasketService) Add(params AddParams) (*Basket, error) {
	return nil, nil
}

func (s *BasketService) Get() (*Basket, error) {
	return nil, nil
}

func (s *BasketService) Remove(params RemoveParams) (*Basket, error) {
	return nil, nil
}

func NewBasketService() *BasketService {
	return &BasketService{}
}
