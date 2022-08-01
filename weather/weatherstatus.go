package weather

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
	"training/helper"

	"github.com/bamzi/jobrunner"
)

type WeatherStatus struct {
	Status struct {
		Water int `json:"Water"`
		Wind  int `json:"Wind"`
	} `json:"Status"`
}

func (status WeatherStatus) Run() {
	status = WeatherStatus{}
	status.Status.Water = helper.GetRandomNumber(1, 100)
	status.Status.Wind = helper.GetRandomNumber(1, 100)
	value, err := json.Marshal(status)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("./static/status.json", value, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func Status(w http.ResponseWriter, r *http.Request) {
	jobrunner.Start()
	jobrunner.Now(WeatherStatus{})
	jobrunner.Every(15*time.Second, WeatherStatus{})

	if r.Method == "GET" {
		tpl, err := template.ParseFiles("./static/template.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		var resultValue string
		value, err := ioutil.ReadFile("./static/status.json")
		if err != nil {
			fmt.Println(err)
		}
		status := WeatherStatus{}
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
}
