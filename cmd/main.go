// Open Weather Map (openweathermap.org) API client for go. Requires API key for One Call API 3.0 (include in .env file as API_KEY).
package main

import (
	"errors"
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


// loadEnv loads .env file and checks for required variables:
func loadEnv() (error) {
	err := godotenv.Load("../.env")
	if err != nil{
		return errors.New("missing: required .env file not found")
	}
	apiKey:= os.Getenv("API_KEY")
	if apiKey == ""{
		return errors.New("empty value: API_KEY has no assigned value")
	}
	return nil
}

// updateCity prompts user for city name and updates value at variable city
func updateCity(str *string){
	var userCity string
	fmt.Println("Enter your city name:")
	fmt.Scanln(&userCity) 
	*str = userCity
}

// Build URL from latitude, longitued and api key
func buildURL() string {
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
	err := loadEnv()
	if err != nil{
		log.Fatal("Fatal error:", err)
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
