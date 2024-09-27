package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"os"
)

var config *Config

type DatabaseConfig struct {
	User     string
	Password string
	Dbname   string
	Host     string
	Port     string
}

func ReadENV() (*DatabaseConfig, error) {
	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	user := os.Getenv("DBUSER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	return &DatabaseConfig{user, password, dbname, host, port}, nil
}

func ReadConfig(db *sql.DB) (*Config, error) {
	rows, err := db.Query("SELECT * FROM config")
	if err != nil {
		return nil, err
	}

	config := Config{}

	rows.Next()
	err = rows.Scan(&config.id, &config.AccessTokenTime, &config.RefreshTokenTime)

	if err != nil {
		return nil, err
	}

	return &config, nil
}
