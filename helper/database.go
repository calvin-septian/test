package helper

import (
	"database/sql"
	"fmt"
)

type SQLServer struct {
	LocalDB *sql.DB
}

func NewSQLConnection() *SQLServer {
	s := SQLServer{}
	db, err := sql.Open("sqlserver", "server=.\\SQLEXPRESS;database=training;trusted_connection=yes")
	if err != nil {
		fmt.Println(err)
	}
	s.LocalDB = db

	return &s
}
