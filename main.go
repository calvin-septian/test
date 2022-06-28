package main

import (
	"fmt"
	"time"

	"training/entity"
	"training/helper"
	"training/service"
)

func main() {
	fmt.Println("Hello World!")

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			fmt.Printf("angka %d ganjil\n", i)
		} else {
			fmt.Printf("angka %d genap\n", i)
		}
	}

	name := []string{"andi", "budi", "cacing"}
	for _, v := range name {
		fmt.Printf("nama saya %s\n", v)
	}

	helper.GetBiodata()

	userSvc := service.NewUserService()
	userSvc.Register(&entity.User{
		Id:        0,
		Username:  "adi123",
		Email:     "adi123@gmail.com",
		Password:  "password123",
		Age:       9,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

}
