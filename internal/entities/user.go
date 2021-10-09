package entities

import "github.com/latolukasz/beeorm"

type User struct {
	beeorm.ORM `orm:"redisSearch=search"`
	ID         uint64
	Email      string `orm:"searchable"`
	Password   string
}

func (user *User) GetUniqueFieldName() string {
	return "Email"
}

func (user *User) GetPassword() string {
	return user.Password
}
