package main

import (
	"fmt"

	"github.com/Dhamodharancs17/findskilleddev/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.InitDB()
}

func main() {
	router()

	fmt.Println("Demo configuration for findskilleddev")
	db, err := initializers.DB.DB()
	if err != nil {
		fmt.Println("Error getting DB instance:", err)
		return
	}
	defer db.Close()
	
}
