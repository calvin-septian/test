package helper

import (
	"context"
	"fmt"
	"testing"
	"time"
	"training/entity"
)

func TestGetOrders(t *testing.T) {
	ConnectDB()
	defer CloseConnectionDB()
	orders, items := Context.mssql.GetOrders(context.Background())

	fmt.Println(orders)
	fmt.Println(items)
}

func TestCreateOrder(t *testing.T) {
	ConnectDB()
	defer CloseConnectionDB()

	obj := entity.Order{
		Order_id:      1001,
		Customer_name: "rangga",
		Order_at:      time.Now(),
		Items: []entity.Item{{Item_id: 1234,
			Item_code:   "a1a",
			Description: "ssd",
			Quantity:    10,
			Order_id:    1001}}}

	Context.mssql.CreateOrder(context.Background(), obj)
}

func TestUpdateOrder(t *testing.T) {
	ConnectDB()
	defer CloseConnectionDB()

	obj := entity.Order{
		Customer_name: "rangga",
		Order_at:      time.Now(),
		Items: []entity.Item{
			{
				Item_id:     1234,
				Item_code:   "aa1",
				Description: "benang",
				Quantity:    15,
			},
		},
	}

	Context.mssql.UpdateOrder(context.Background(), obj, "16")
}

func TestDeleteOrder(t *testing.T) {
	ConnectDB()
	defer CloseConnectionDB()

	Context.mssql.DeleteOrder(context.Background(), "16")
}
