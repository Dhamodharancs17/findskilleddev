package organization

import (
	"encoding/json"
	"net/http"

	"github.com/Dhamodharancs17/findskilleddev/initializers"
)


func CreateOrganization(w http.ResponseWriter, r *http.Request) {
	// create user
	var organization Organization
	if err := json.NewDecoder(r.Body).Decode(&organization); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Create(&organization).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetOrganization(w http.ResponseWriter, r *http.Request) {
	// get user
	var organization Organization
	if err := initializers.DB.First(&organization, r.FormValue("id")).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Return the user
	json.NewEncoder(w).Encode(organization)
}
