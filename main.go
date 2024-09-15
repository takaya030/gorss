package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/takaya030/gorss/controller"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Cannot load env file: %v", err)
	}

	router := controller.GetRouter()
	router.Run(":8080")
}
