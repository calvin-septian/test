package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"

	"training/helper"

	"github.com/bamzi/jobrunner"
	"github.com/gorilla/mux"
)

type Status struct {
	Status struct {
		Water int `json:"Water"`
		Wind  int `json:"Wind"`
	} `json:"Status"`
}

func (status Status) Run() {
	status = Status{}
	status.Status.Water = helper.GetRandomNumber(1, 100)
	status.Status.Wind = helper.GetRandomNumber(1, 100)
	value, err := json.Marshal(status)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("status.json", value, 0644)
	if err != nil {
		fmt.Println(err)
	}
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
			value, err := ioutil.ReadFile("status.json")
			if err != nil {
				fmt.Println(err)
			}
			status := Status{}
			err = json.Unmarshal(value, &status)
			if err != nil {
				fmt.Println(err)
			}

			var StatusWater, StatusWind string
			if status.Status.Water < 5 {
				StatusWater = "aman"
			} else if status.Status.Water > 5 && status.Status.Water < 9 {
				StatusWater = "siaga"
			} else {
				StatusWater = "bahaya"
			}

			if status.Status.Wind < 6 {
				StatusWind = "aman"
			} else if status.Status.Wind > 6 && status.Status.Wind < 16 {
				StatusWind = "siaga"
			} else {
				StatusWind = "bahaya"
			}

			resultValue = fmt.Sprintf("StatusWater : %s (%dm)\n StatusWind : %s (%dm/s)", StatusWater, status.Status.Water, StatusWind, status.Status.Wind)

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
