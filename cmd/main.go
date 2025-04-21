// Open Weather Map (openweathermap.org) API client for go. Requires API key for One Call API 3.0 (include in .env file as API_KEY).
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Global variables
var (
	apiKey string = ""
	url string = ""
)

type Location struct {
	CityName 		string
	CountryCode 	string 
	Latitude 		string
	Longitude 		string
}


func (l *Location) setLocationFields() {
	var inputName string
	var inputCode string
	fmt.Println("Enter city name:")
	fmt.Scanln(&inputName)
	fmt.Println("Enter two letter country code (example: Canada = CA):")
	fmt.Scanln(&inputCode)
	l.CityName = inputName
	l.CountryCode = inputCode
}

// loadEnv loads .env file required variables, assigns API_KEY value to apiKey:
func loadEnv(str *string) (error) {
	err := godotenv.Load("../.env")
	if err != nil{
		return errors.New("missing: required .env file not found")
	}
	key:= os.Getenv("API_KEY")
	if key == ""{
		return errors.New("empty value: API_KEY has no assigned value")
	}else{
		*str = key
	}
	return nil
}

// inputCity prompts user for city name and updates value at variable city
func inputCity(str *string){
	var inputCity string
	fmt.Println("enter your city name:")
	fmt.Scanln(&inputCity) 
	*str = inputCity
}

// inputCountryCode prompts user for ISO 3166 Country code (2 letters)
func inputCountryCode(str *string) error{
	var inputCountryCode string
	fmt.Println("enter your country code (ISO 2 letter code. Example: Canada = CA)")
	fmt.Scanln(&inputCountryCode)
	*str = inputCountryCode
	return errors.New("error: country code must be 2 letters long")
}

// Build URL from latitude, longitude and api key
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

	// 01 load .env file and check for API key:
	err := loadEnv(&apiKey)
	if err != nil{
		log.Fatal("Fatal error:", err)
	}
	
	// 02 Build location data from user input and API call:
	city := Location{"","","",""} // instantiate empty Location struct 
	city.setLocationFields() 
	fmt.Println(city.CityName)
	fmt.Println(city.CountryCode)

	// // Get latitude and Longitude for user city: 
	// urlDirect = fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q={city name},{state code},{country code}&limit={limit}&appid=%s", apiKey)

	// // Build request URL: 
	// url = fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=33.44&lon=-94.04&exclude=hourly,daily&appid=%s", apiKey)

	// // send HTTP Request to openweathermap.org:
	// responseBody, _ := sendRequest(url)
	// fmt.Println(responseBody)

	// // write .txt file from response json data:
	// writeText(responseBody)
}
