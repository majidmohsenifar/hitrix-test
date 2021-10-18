package product

import (
	"fmt"
	"hitrix-test/internal/entities"

	"github.com/latolukasz/beeorm"
)

var (
	ErrProductNotFound = fmt.Errorf("product not found")
)

type Service struct {
	ormEngine *beeorm.Engine
}

func (s *Service) List(name string, minPrice float64, maxPrice float64) ([]*entities.Product, error) {
	var products []*entities.Product
	query := beeorm.NewRedisSearchQuery()
	query.FilterFloatMinMax("Price", minPrice, maxPrice)
	query.QueryFieldPrefixMatch("Name", name)
	_ = s.ormEngine.RedisSearch(&products, query, beeorm.NewPager(1, 100))
	return products, nil
}

func (s *Service) FindByID(productID uint64) (*entities.Product, error) {
	var products []*entities.Product
	query := beeorm.NewRedisSearchQuery()
	query.FilterUint("ID", productID)
	//TODO here we should use redisSearchOne but it does not work and panics handle this later
	total := s.ormEngine.RedisSearch(&products, query, beeorm.NewPager(1, 1))
	if total > 0 {
		return products[0], nil
	}
	return nil, ErrProductNotFound

}

func New(ormEngine *beeorm.Engine) *Service {
	return &Service{
		ormEngine: ormEngine,
	}
}
