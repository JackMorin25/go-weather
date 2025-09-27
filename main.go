package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found, loading environment variables from system.")
	}
}

func main() {
	key := os.Getenv("KEY")
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=Memphis&aqi=no", key)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather api not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
