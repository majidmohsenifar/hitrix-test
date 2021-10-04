package entities

import "github.com/latolukasz/beeorm"

type Product struct {
	beeorm.ORM `orm:"redisCache"`
	ID         uint64
	Name       string
	Price      float64
}
