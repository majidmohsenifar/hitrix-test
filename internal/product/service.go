package product

import (
	"hitrix-test/internal/entities"

	"github.com/latolukasz/beeorm"
)

type Service struct {
	ormEngine *beeorm.Engine
}

func (s *Service) Create(name string, price float64) error {
	p := &entities.Product{
		Name:  name,
		Price: price,
	}
	s.ormEngine.Flush(p)
	return nil
}

func (s *Service) List(name string, minPrice float64, maxPrice float64) ([]entities.Product, error) {
	result := []entities.Product{{
		Name:  "someName",
		Price: 1.2,
	}}

	return result, nil
}

func (s *Service) Update(ID uint64, name string, price float64) error {
	panic("implement me")
	//p := &models.Product{
	//Name:  name,
	//Price: price,
	//}
	//s.ormEngine.Flush(p)
	return nil
}
func (s *Service) Delete(name string, price float64) error {
	panic("implement me")
	//p := &models.Product{
	//Name:  name,
	//Price: price,
	//}
	//s.ormEngine.Delete(p)
	return nil
}

func New(ormEngine *beeorm.Engine) *Service {
	return &Service{
		ormEngine: ormEngine,
	}
}
