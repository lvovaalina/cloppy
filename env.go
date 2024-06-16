package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_PATH string
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err)
		return
	}

	DB_PATH = os.Getenv("DB_PATH")
}
