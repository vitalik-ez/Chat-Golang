package persistence

import (
	"fmt"

	"github.com/vitalik-ez/Chat-Golang/domain/entity"
)

type UserRepo struct {
	Db string
}

func (r *UserRepo) SaveUser(user *entity.User) {
	fmt.Println("Save User")

}

func (r *UserRepo) GetUser(id uint64) {
	fmt.Println("GetUser")
}

func (r *UserRepo) GetUsers() {
	fmt.Println("GetUsers")
}

func (r *UserRepo) GetUserByEmailAndPassword(u *entity.User) {
	fmt.Println("GetUserByEmailAndPassword")
}
