package initializers

import (
    "log"
    "github.com/joho/godotenv"
)

func LoadEnvVariables(){
	err := godotenv.Load()
	if err != nil {
	//   log.Fatal("Error loading .env file")
	  log.Printf("Error loading .env file: %v", err)
	}
}