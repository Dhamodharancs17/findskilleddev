package main

import "github.com/Dhamodharancs17/findskilleddev/initializers"

func init() {
	initializers.LoadEnvVariables()
	initializers.InitDB()
}

func main() {
	//migrate the schema
	initializers.DB.AutoMigrate(
		// Add the model files here
	)
}
