package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"training/entity"
	"training/helper"
	"training/service"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello World!")

	helper.ConnectDB()

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			fmt.Printf("angka %d ganjil\n", i)
		} else {
			fmt.Printf("angka %d genap\n", i)
		}
	}

	name := []string{"andi", "budi", "cacing"}
	for _, v := range name {
		go func(nama string) {
			fmt.Printf("nama saya %s\n", nama)
		}(v)
	}
	time.Sleep(1 * time.Second)

	helper.GetBiodata()

	r := mux.NewRouter()
	r.HandleFunc("/hello", hello)
	r.HandleFunc("/register", register)
	r.HandleFunc("/user", helper.UsersHandler)
	r.HandleFunc("/user/{id}", helper.UsersHandler)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8090",
	}
	srv.ListenAndServe()
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func register(w http.ResponseWriter, req *http.Request) {
	userSvc := service.NewUserService()
	user := userSvc.Register(&entity.User{
		Id:        0,
		Username:  "adi123",
		Email:     "adi123@gmail.com",
		Password:  "password123",
		Age:       9,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	userData, _ := json.Marshal(user)
	w.Header().Add("Content-Type", "application/json")
	w.Write(userData)
}
