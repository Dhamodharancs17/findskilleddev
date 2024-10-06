package main

import (
	"github.com/Dhamodharancs17/findskilleddev/initializers"
	"github.com/Dhamodharancs17/findskilleddev/internal/assessment"
	"github.com/Dhamodharancs17/findskilleddev/internal/candidate"
	"github.com/Dhamodharancs17/findskilleddev/internal/organization"
	"github.com/Dhamodharancs17/findskilleddev/internal/question"
	"github.com/Dhamodharancs17/findskilleddev/internal/submission"
	"github.com/Dhamodharancs17/findskilleddev/internal/user"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.InitDB()
}

func main() {
	//migrate the schema
	initializers.DB.AutoMigrate(
		// Add the model files here
		&assessment.Assessment{},
		&candidate.Candidate{},
		&organization.Organization{},
		&question.Question{},
		&submission.Submission{},
		&user.User{},
	)
}
