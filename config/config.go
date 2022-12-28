package config

import (
	"fmt"
	"os"
)

func GetPostgresDSN() string {
	var HOST = os.Getenv("POSTGRES_HOST")
	var USER = os.Getenv("POSTGRES_USER")
	var PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	var DATABASE_NAME = os.Getenv("POSTGRES_DATABASE_NAME")
	var PORT = os.Getenv("POSGRES_PORT")
	var SSL_MODE = os.Getenv("SSL_MODE")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", HOST, USER, PASSWORD, DATABASE_NAME, PORT, SSL_MODE)
}
