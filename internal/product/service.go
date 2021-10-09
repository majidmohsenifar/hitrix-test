package product

import (
	"hitrix-test/internal/entities"

	"github.com/latolukasz/beeorm"
)

type Service struct {
	ormEngine *beeorm.Engine
}

func (s *Service) List(name string, minPrice float64, maxPrice float64) ([]entities.Product, error) {
	result := []entities.Product{{
		Name:  "someName",
		Price: 1.2,
	}}

	return result, nil
}

func New(ormEngine *beeorm.Engine) *Service {
	return &Service{
		ormEngine: ormEngine,
	}
}
