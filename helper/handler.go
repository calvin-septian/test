package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"training/entity"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

// Global app context
type applicationContext struct {
	mssql *SQLServer
}

//Context type helper
var Context applicationContext

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

func ConnectDB() {
	sql := NewSQLConnection()
	Context.mssql = sql
}
