package product

import (
	"hitrix-test/internal/entities"

	"github.com/latolukasz/beeorm"
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

func New(ormEngine *beeorm.Engine) *Service {
	return &Service{
		ormEngine: ormEngine,
	}
}
