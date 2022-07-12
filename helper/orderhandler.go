package helper

import (
	"context"
	"database/sql"
	"fmt"
	"training/entity"
)

func (s *SQLServer) GetOrders(ctx context.Context) ([]entity.Order, map[int]entity.Item) {
	rows, err := s.LocalDB.Query("select * from [order]")
	if err != nil {
		fmt.Println("error: ", err)
	}
	defer rows.Close()

	var orders []entity.Order
	for rows.Next() {
		data := entity.Order{}

		err := rows.Scan(
			&data.Order_id,
			&data.Customer_name,
			&data.Order_at)
		if err != nil {
			fmt.Printf("[mssql] Failed reading rows: %v", err)
		}

		orders = append(orders, data)
	}

	rows1, err := s.LocalDB.Query("select * from [item]")
	if err != nil {
		fmt.Println("error: ", err)
	}
	defer rows1.Close()

	items := make(map[int]entity.Item)
	for rows1.Next() {
		data := entity.Item{}

		err := rows1.Scan(
			&data.Item_id,
			&data.Item_code,
			&data.Description,
			&data.Quantity,
			&data.Order_id)
		if err != nil {
			fmt.Printf("[mssql] Failed reading rows: %v", err)
		}

		if _, ok := items[data.Order_id]; !ok {
			items[data.Order_id] = data
		}
	}

	return orders, items
}

func (s *SQLServer) CreateOrder(ctx context.Context, data entity.Order) {
	var id int
	err := s.LocalDB.QueryRow("AddOrder",
		sql.Named("Customer_name", data.Customer_name),
		sql.Named("Order_at", data.Order_at)).Scan(&id)
	if err != nil {
		fmt.Println("error: ", err)
	}

	for _, v := range data.Items {
		_, err = s.LocalDB.ExecContext(ctx, "AddItem",
			sql.Named("Item_id", v.Item_id),
			sql.Named("Item_code", v.Item_code),
			sql.Named("Description", v.Description),
			sql.Named("Quantity", v.Quantity),
			sql.Named("Order_id", id))
		if err != nil {
			fmt.Println("error: ", err)
		}
	}
}

func (s *SQLServer) UpdateOrder(ctx context.Context, data entity.Order, orderId string) {
	_, err := s.LocalDB.Exec("update [order] set Order_id = @Order_id, Customer_name = @Customer_name, Order_at = @Order_at where Order_id = @Order_id",
		sql.Named("Order_id", orderId),
		sql.Named("Customer_name", data.Customer_name),
		sql.Named("Order_at", data.Order_at))
	if err != nil {
		fmt.Println("error: ", err)
	}

	for _, v := range data.Items {
		_, err = s.LocalDB.Exec("update [item] set Item_id = @Item_id, Item_code = @Item_code, Description = @Description, Quantity = @Quantity, Order_id = @Order_id where Order_id = @Order_id",
			sql.Named("Item_id", v.Item_id),
			sql.Named("Item_code", v.Item_code),
			sql.Named("Description", v.Description),
			sql.Named("Quantity", v.Quantity),
			sql.Named("Order_id", orderId))
		if err != nil {
			fmt.Println("error: ", err)
		}
	}
}

func (s *SQLServer) DeleteOrder(ctx context.Context, orderId string) {
	_, err := s.LocalDB.Exec("delete from [item] where Order_id = @Order_id",
		sql.Named("Order_id", orderId))
	if err != nil {
		fmt.Println("error: ", err)
	}

	_, err = s.LocalDB.Exec("delete from [order] where Order_id = @Order_id",
		sql.Named("Order_id", orderId))
	if err != nil {
		fmt.Println("error: ", err)
	}
}
