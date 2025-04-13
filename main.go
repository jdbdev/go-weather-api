package main

// Simple app connects to openweathermap.org API. Abbreviated to owm in app.
// Requires API key for One Call API 3.0 (include in .env file as API_KEY)

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
	City []location
}
type location struct{
	City string
}

func main(){

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

// Prompt user for city name: 
var city string 
fmt.Println("enter name of city:")
fmt.Scanln(&city)
fmt.Println("Your city is:", city)

// ---------- 02 BUILD REQUEST URL ---------- //

// Sample API Call from owm:
urlDemo:= fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=33.44&lon=-94.04&exclude=hourly,daily&appid=%s", apiKey)
fmt.Println(urlDemo)

// ---------- 03 HTTP REQUEST AND RESPONSE HANDLING ---------- //

response, err := http.Get(urlDemo)
	if err != nil{
		log.Fatal("Error sending request to www.openweathermap.org")
	}
	defer response.Body.Close()
	// if response.StatusCode != 200 {
	// 	log.Fatal("Something went wrong!")
	// }
	fmt.Println("The HTTP status code is:", response.StatusCode)
	
	body, err:= io.ReadAll(response.Body)
	if err != nil{
		log.Fatal("there is no response body")
	}
	fmt.Println(string(body))
}

// Functions called from main{}
func getCity (x string) (string, error){
	return "", nil
}
func buildUrl (city,url string) (string, error){
	return "", nil
}
func getWeather (api_key,url string) (string, error){
	response, err := http.Get(url)
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