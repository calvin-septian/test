package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
	"time"

	"training/helper"

	"github.com/bamzi/jobrunner"
	"github.com/gorilla/mux"
)

var list sync.Map

type Status struct {
	Water int `json:"Water"`
	Wind  int `json:"Wind"`
}

func (status Status) Run() {
	status = Status{}
	status.Water = helper.GetRandomNumber(1, 100)
	status.Wind = helper.GetRandomNumber(1, 100)
	list.Store("Status", status)
}

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

	jobrunner.Start()
	jobrunner.Now(Status{})
	jobrunner.Every(15*time.Second, Status{})

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
	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			tpl, err := template.ParseFiles("template.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			var resultValue string
			list.Range(func(key, value interface{}) bool {
				isValid := false
				if key.(string) == "Status" {
					status := value.(Status)
					var StatusWater, StatusWind string
					if status.Water < 5 {
						StatusWater = "aman"
					} else if status.Water > 5 && status.Water < 9 {
						StatusWater = "siaga"
					} else {
						StatusWater = "bahaya"
					}

					if status.Wind < 6 {
						StatusWind = "aman"
					} else if status.Wind > 6 && status.Wind < 16 {
						StatusWind = "siaga"
					} else {
						StatusWind = "bahaya"
					}

					resultValue = fmt.Sprintf("StatusWater : %s (%dm)\n StatusWind : %s (%dm/s)", StatusWater, status.Water, StatusWind, status.Wind)
					isValid = true
				}
				return isValid
			})

			tpl.Execute(w, resultValue)
			return
		}
		http.Error(w, "Invalid Method", http.StatusBadRequest)
	})
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8090",
	}
	srv.ListenAndServe()
}
