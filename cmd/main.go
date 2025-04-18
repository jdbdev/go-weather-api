// Open Weather Map (openweathermap.org) API client for go. Requires API key for One Call API 3.0 (include in .env file as API_KEY).
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jdbdev/go-weather-api/types"

	"github.com/joho/godotenv"
)

// Global variables
var (
	apiKey string = ""
	city string = ""
	url string = ""
)


// Load .env variables and check for API key:
func loadEnv() (string, error){
	err := godotenv.Load("../.env")
	if err != nil{
		log.Fatal("error loading .env file")
	}
	apiKey:= os.Getenv("API_KEY")
	if apiKey == ""{
		log.Fatal("Error: API_KEY has no assigned value")
	}
	return "success loading .env file!", nil
}
// Function prompts user for city name and updates value at variable city
func updateCity(str *string){
	var userCity string
	fmt.Println("Enter your city name:")
	fmt.Scanln(&userCity) 
	*str = userCity
}
// Build URL from latitude, longitued and api key
func buildURL(lat, long, key string) string {
	return ""
}
// Send HTTP Request to openweathermap.org
func sendRequest (x string) (string, error){
	response, err := http.Get(x)
	if err != nil{
		log.Fatal("error")
	}
	defer response.Body.Close()
	body, err:= io.ReadAll(response.Body)
	if err != nil{
		log.Fatal("there is no response body")
	}
	return string(body), nil
}

// Write .txt file from json data (string):
func writeText(x string){
	file, err := os.Create("../data_sample.txt")
	if err != nil{
		fmt.Println(err)
	} else{
		file.WriteString(x)
	}
}

func main(){
	// load packages:
	types.LoadTypes()
	// load .env file and check for API key:
	env, err := loadEnv()
	fmt.Println(env, err)
	apiKey = os.Getenv("API_KEY")
		if apiKey == ""{
			log.Fatal("Error: API_KEY has no assigned value")
		}

	// Prompt user for city: 
	updateCity(&city) 
	fmt.Printf("The weather in %s is:", city)

	// Get latitude and Longitude for user city: 

	// Build request URL: 
	url = fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=33.44&lon=-94.04&exclude=hourly,daily&appid=%s", apiKey)

	// send HTTP Request to openweathermap.org:
	responseBody, _ := sendRequest(url)

	// write .txt file from response json data:
	writeText(responseBody)
}
