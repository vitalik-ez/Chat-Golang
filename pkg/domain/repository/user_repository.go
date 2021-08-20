package repository

import "github.com/vitalik-ez/Chat-Golang/domain/entity"

type UserRepository interface {
	SaveUser(*entity.User)
	GetUser(uint64)
	GetUsers()
	GetUserByEmailAndPassword(*entity.User)
}
