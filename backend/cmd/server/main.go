package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load env vars: %v", err)
	}

	config := loadConfig()

	fmt.Println(config)

}
