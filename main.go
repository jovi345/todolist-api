package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/todos-api/jovi345/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := router.RegisterRoute()

	r.Run(":8080")
}
