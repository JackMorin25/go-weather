package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

type Weather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}
type Location struct {
	Name      string  `json:"name"`
	Region    string  `json:"region"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	Localtime string  `json:"localtime"`
}
type Current struct {
	Temp_f    float64   `json:"temp_f"`
	Condition Condition `json:"condition"`
	Humidity  float64   `json:"humidity"`
}
type Condition struct {
	Text string `json:"text"`
}

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found, loading environment variables from system.")
	}
}

func format(weather Weather) {
	fmt.Print("WEATHER:\n")
	fmt.Printf("%s", weather.Location.Name)
	fmt.Printf("%f", weather.Current.Temp_f)
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
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}
	format(weather)
}
