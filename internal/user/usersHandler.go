package user

import (
	"encoding/json"
	"net/http"

	"github.com/Dhamodharancs17/findskilleddev/initializers"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// create user
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	// get user
	var user User
	if err := initializers.DB.First(&user, r.FormValue("id")).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Return the user
	json.NewEncoder(w).Encode(user)
}
