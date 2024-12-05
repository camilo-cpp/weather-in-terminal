package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"camilo-cpp/weather-in-terminal/pkg/weather"
)

func main() {
	loc := "Bogotá"

	if len(os.Args) > 1 {
		loc = os.Args[1]
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("WEATHER_API_KEY")

	w, err := weather.GetWeather(loc, apiKey)
	if err != nil {
		panic(err)
	}

	weather.DisplayWeather(w)
}
