package main

// Simple app connects to openweathermap.org API

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)


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
// Print API key:
fmt.Println("Your API Key is:", apiKey)
}
