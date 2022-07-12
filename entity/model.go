package entity

import (
	"time"
)

type User struct {
	Id        int
	Username  string
	Email     string
	Password  string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Order struct {
	Order_id      int       `json:"Order_id"`
	Customer_name string    `json:"Customer_name"`
	Order_at      time.Time `json:"Order_at"`
	Items         []Item    `json:"Items"`
}

type Item struct {
	Item_id     int    `json:"Item_id"`
	Item_code   string `json:"Item_code"`
	Description string `json:"Description"`
	Quantity    int    `json:"Quantity"`
	Order_id    int    `json:"Order_id"`
}
