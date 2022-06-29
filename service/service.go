package service

import (
	"fmt"
	"time"
	"training/entity"
)

type UserServiceIface interface {
	Register(user *entity.User) entity.User
}

type UserSvc struct {
	ListUser map[string]entity.User
}

func NewUserService() UserServiceIface {
	list := make(map[string]entity.User)
	list["budi123"] = entity.User{
		Id:        0,
		Username:  "budi123",
		Email:     "budi123@gmail.com",
		Password:  "password123",
		Age:       9,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	list["dodi123"] = entity.User{
		Id:        1,
		Username:  "dodi123",
		Email:     "dodi@gmail.com",
		Password:  "password123",
		Age:       10,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &UserSvc{ListUser: list}
}

func (u *UserSvc) Register(user *entity.User) entity.User {
	if _, ok := u.ListUser[user.Username]; ok {
		fmt.Println("Failed register, username already exist")
	} else {
		fmt.Println("Success register user")
	}

	fmt.Println(user)
	return *user
}
