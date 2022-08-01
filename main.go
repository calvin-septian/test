package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"training/helper"
	"training/weather"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"

	flag "github.com/spf13/pflag"
)

type options struct {
	configFilename string
}

type configuration struct {
	Username string `yaml:"username"`
}

var config = configuration{}

func main() {
	fmt.Println("Hello World!")

	args := loadOptions()
	err := parseConfig(args.configFilename)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(config.Username)

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
	r.HandleFunc("/status", weather.Status)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8090",
	}
	srv.ListenAndServe()
}

func loadOptions() options {
	o := options{}

	flag.StringVar(&o.configFilename, "config", "", "Path to the config files")

	flag.CommandLine.SortFlags = false

	flag.Parse()

	return o
}

func parseConfig(configFile string) error {

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
		return err
	}

	return nil
}
