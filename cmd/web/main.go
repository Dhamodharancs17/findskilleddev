package main

import (
	"fmt"

	"github.com/Dhamodharancs17/findskilleddev/initializers"
)

func init () {
	initializers.LoadEnvVariables()
	initializers.InitDB()
}

func main () {
	fmt.Println("Demo configuration for findskilleddev")
}