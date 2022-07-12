package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"training/entity"
	"training/service"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

var listUser = map[string]entity.User{
	"2": {
		Id:        2,
		Username:  "budi123",
		Email:     "budi123@gmail.com",
		Password:  "password123",
		Age:       9,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	"3": {
		Id:        3,
		Username:  "cantya123",
		Email:     "cantya123@gmail.com",
		Password:  "password123",
		Age:       9,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !Auth(w, r) {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte(`something went wrong`))
		return false
	}

	isValid := (username == "user") && (password == "pass")
	if !isValid {
		w.Write([]byte(`wrong username/password`))
		return false
	}

	return true
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func Register(w http.ResponseWriter, req *http.Request) {
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

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	switch r.Method {
	case http.MethodGet:
		if id != "" { // get by id
			getUsersByIDHandler(w, r, id)
		} else { // get all
			getUsersHandler(w, r)
		}
	case http.MethodPost:
		createUsersHandler(w, r)
	case http.MethodPut:
		updateUserHandler(w, r, id)
	case http.MethodDelete:
		deleteUserHandler(w, r, id)
	}
}

func getUsersByIDHandler(w http.ResponseWriter, r *http.Request, id string) {
	var result []byte
	if v, ok := listUser[id]; ok {
		w.Header().Add("Content-Type", "application/json")
		result, _ = json.Marshal(v)
		w.Write(result)
	}

	list := Context.mssql.GetAllUser(context.Background())
	for _, v := range list {
		if id == fmt.Sprint(v.Id) {
			w.Header().Add("Content-Type", "application/json")
			result, _ = json.Marshal(v)
			w.Write(result)
		}
	}
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	result, _ := json.Marshal(listUser)
	w.Write(result)

	list := Context.mssql.GetAllUser(context.Background())
	result, _ = json.Marshal(list)
	w.Write(result)
}

func createUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	if _, ok := listUser[fmt.Sprint(user.Id)]; !ok {
		listUser[fmt.Sprint(user.Id)] = user
		w.Write([]byte("success create user"))
	}

	Context.mssql.AddUser(context.Background(), user)

}

func updateUserHandler(w http.ResponseWriter, r *http.Request, id string) {
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	if _, ok := listUser[fmt.Sprint(user.Id)]; ok {
		listUser[fmt.Sprint(user.Id)] = user
		w.Write([]byte("success update user"))
	} else {
		w.Write([]byte("user not found"))
	}

	Context.mssql.AddUser(context.Background(), user)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request, id string) {
	if _, ok := listUser[id]; ok {
		delete(listUser, id)
		w.Write([]byte("user deleted"))
	} else {
		w.Write([]byte("user not found"))
	}

	Context.mssql.DeleteUser(context.Background(), id)
}

func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	orderId := params["orderId"]

	switch r.Method {
	case http.MethodGet:
		getOrdersHandler(w, r)
	case http.MethodPost:
		createOrderHandler(w, r)
	case http.MethodPut:
		updateOrdersHandler(w, r, orderId)
	case http.MethodDelete:
		deleteOrderHandler(w, r, orderId)
	}
}

func getOrdersHandler(w http.ResponseWriter, r *http.Request) {
	var result []byte
	var listOrder []entity.Order
	orders, items := Context.mssql.GetOrders(context.Background())
	for _, v := range orders {
		if item, ok := items[v.Order_id]; ok {
			v.Items = append(v.Items, item)
		}
		listOrder = append(listOrder, v)
	}
	w.Header().Add("Content-Type", "application/json")
	result, _ = json.Marshal(listOrder)
	w.Write(result)
}

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order entity.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	Context.mssql.CreateOrder(context.Background(), order)
}

func updateOrdersHandler(w http.ResponseWriter, r *http.Request, orderId string) {
	var order entity.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	Context.mssql.UpdateOrder(context.Background(), order, orderId)
}

func deleteOrderHandler(w http.ResponseWriter, r *http.Request, orderId string) {
	Context.mssql.DeleteOrder(context.Background(), orderId)
}

func GetUserUrl(w http.ResponseWriter, r *http.Request) {
	url := "https://random-data-api.com/api/users/random_user?size=10"
	var result []map[string]interface{}
	res, _ := requestHttp(url, "GET", []byte(""))
	_ = json.Unmarshal(res, &result)

	type data struct {
		Id         float64 `json:"id"`
		Uid        string  `json:"uid"`
		First_name string  `json:"first_name"`
		Last_name  string  `json:"last_name"`
		Username   string  `json:"username"`
		Address    struct {
			City           string `json:"city"`
			Street_name    string `json:"street_name"`
			Street_address string `json:"street_address"`
			Zip_code       string `json:"zip_code"`
			State          string `json:"state"`
			Country        string `json:"country"`
			Coordinates    struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"coordinates"`
		} `json:"address"`
	}
	var listdata []data

	for _, v := range result {
		data := data{}
		data.Id = v["id"].(float64)
		data.Uid = v["uid"].(string)
		data.First_name = v["first_name"].(string)
		data.Last_name = v["last_name"].(string)
		data.Username = v["username"].(string)
		address := v["address"].(map[string]interface{})
		coordinates := address["coordinates"].(map[string]interface{})
		data.Address.City = address["city"].(string)
		data.Address.Street_name = address["street_name"].(string)
		data.Address.Street_address = address["street_address"].(string)
		data.Address.Zip_code = address["zip_code"].(string)
		data.Address.State = address["state"].(string)
		data.Address.Country = address["country"].(string)
		data.Address.Coordinates.Lat = coordinates["lat"].(float64)
		data.Address.Coordinates.Lng = coordinates["lng"].(float64)
		listdata = append(listdata, data)
	}
	userData, err := json.Marshal(listdata)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(userData)
}
