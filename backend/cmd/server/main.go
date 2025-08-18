package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load env vars: %v", err)
	}

	config := loadConfig()

	app := &application{
		config: config,
	}

	router := app.router()

	http.ListenAndServe(":4000", router)

	fmt.Println(config)

}
