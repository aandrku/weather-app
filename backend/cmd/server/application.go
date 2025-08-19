package main

import (
	"net/http"
	"os"
)

type application struct {
	config Config
	client *http.Client
}

func loadConfig() Config {
	var c Config

	c.weatherAPIKey = os.Getenv("WEATHER_API_KEY")
	c.apiURL = os.Getenv("API_URL")

	return c
}

type Config struct {
	weatherAPIKey string
	apiURL        string
}
