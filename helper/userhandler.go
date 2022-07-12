package helper

import (
	"context"
	"database/sql"
	"fmt"
	"training/entity"
)

func (s *SQLServer) GetAllUser(ctx context.Context) []entity.User {
	rows, err := s.LocalDB.QueryContext(ctx, "GetUser")
	if err != nil {
		fmt.Println("error: ", err)
	}
	defer rows.Close()

	var result []entity.User
	for rows.Next() {
		data := entity.User{}

		err := rows.Scan(
			&data.Id,
			&data.Username,
			&data.Email,
			&data.Password,
			&data.Age,
			&data.CreatedAt,
			&data.UpdatedAt)
		if err != nil {
			fmt.Printf("[mssql] Failed reading rows: %v", err)
		}

		result = append(result, data)
	}

	return result
}

func (s *SQLServer) AddUser(ctx context.Context, data entity.User) {
	_, err := s.LocalDB.ExecContext(ctx, "AddUser",
		sql.Named("Id", data.Id),
		sql.Named("Username", data.Username),
		sql.Named("Email", data.Email),
		sql.Named("Password", data.Password),
		sql.Named("Age", data.Age),
		sql.Named("CreatedAt", data.CreatedAt),
		sql.Named("UpdatedAt", data.UpdatedAt))
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func (s *SQLServer) DeleteUser(ctx context.Context, id string) {
	_, err := s.LocalDB.ExecContext(ctx, "DeleteUser",
		sql.Named("Id", id))
	if err != nil {
		fmt.Println("error: ", err)
	}
}
