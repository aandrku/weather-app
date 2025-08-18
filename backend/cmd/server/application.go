package main

import (
	"net/http"
	"os"
)

type application struct {
	config Config
}

func loadConfig() Config {
	var c Config

	c.client = &http.Client{}
	c.weatherAPIKey = os.Getenv("WEATHER_API_KEY")
	c.apiURL = os.Getenv("API_URL")

	return c
}

type Config struct {
	client        *http.Client
	weatherAPIKey string
	apiURL        string
}
