package main

import (
	"fmt"
	"net/http"
	"time"

	"training/helper"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello World!")

	helper.ConnectDB()
	defer helper.CloseConnectionDB()

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

	// helper.GetBiodata()

	r := mux.NewRouter()
	r.HandleFunc("/hello", helper.Hello)
	r.HandleFunc("/register", helper.Register)
	r.HandleFunc("/user", helper.UsersHandler)
	r.HandleFunc("/user/{id}", helper.UsersHandler)
	r.HandleFunc("/orders", helper.OrdersHandler)
	r.HandleFunc("/orders/{orderId}", helper.OrdersHandler)
	r.Handle("/userurl", helper.Middleware(http.HandlerFunc(helper.GetUserUrl)))
	r.HandleFunc("/newregister", helper.UserRegisterHandler)
	r.HandleFunc("/login", helper.UserLoginHandler)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8090",
	}
	srv.ListenAndServe()
}
