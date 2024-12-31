package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/todos-api/jovi345/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := router.RegisterRoute()

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	url := fmt.Sprintf("%v:%v", host, port)

	r.Run(url)
}
