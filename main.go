package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var botToken string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botToken = os.Getenv("TOKEN")
	if botToken == "" {
		panic("TOKEN environment variable is empty")
	}
	fmt.Println(botToken)
}
