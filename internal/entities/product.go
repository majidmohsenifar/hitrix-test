package entities

import "github.com/latolukasz/beeorm"

type Product struct {
	beeorm.ORM `orm:"redisCache;redisSearch=search"`
	ID         uint64  `orm:"searchable"`
	Name       string  `orm:"searchable"`
	Price      float64 `orm:"searchable"`
}
