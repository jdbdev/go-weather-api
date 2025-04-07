package main

// Simple app connects to openweathermap.org API. Abbreviated to owm in app.

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)


type weatherData struct{
	Name string
}

func main(){

// 01 ENV SETUP: 

// Load .env file:
	err := godotenv.Load()
	if err != nil{
		log.Fatal("error loading .env file")
	}
// Get API key: 
apiKey := os.Getenv("API_KEY")
if apiKey == ""{
	log.Fatal("Error: API_KEY has no assigned value")
}
// Print API key:
fmt.Println("Your API Key is:", apiKey)

// Prompt user for city name: 
var city string 
fmt.Println("enter name of city:")
fmt.Scanln(&city)
fmt.Println("Your city is:", city)

// 02 BUILD REQUEST URL:

// Sample API Call from owm:
urlDemo:= fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=33.44&lon=-94.04&exclude=hourly,daily&appid=%s", apiKey)
fmt.Println(urlDemo)

// // Build the request URL
// url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric&lang=fr", city, apiKey)
// fmt.Println(url)

// Send HTTP request: 
response, err := http.Get(urlDemo)
	if err != nil{
		log.Fatal("Error sending request to www.openweathermap.org")
		fmt.Println()
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		log.Fatal("Something went wrong!")
	}
	fmt.Println("The HTTP status code is:", response.StatusCode)
	
	body, err:= io.ReadAll(response.Body)
	if err != nil{
		log.Fatal("there is no response body")
	}
	fmt.Println(string(body))
}
