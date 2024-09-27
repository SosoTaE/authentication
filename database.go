package main

import (
	"database/sql"
	"fmt"
)

func InitDatabase(databaseConfiguration *DatabaseConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", databaseConfiguration.User, databaseConfiguration.Password, databaseConfiguration.Dbname, databaseConfiguration.Host, databaseConfiguration.Port)
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	return db, nil
}
