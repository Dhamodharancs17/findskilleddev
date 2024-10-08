package organization

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Dhamodharancs17/findskilleddev/initializers"
)

func CreateOrganization(w http.ResponseWriter, r *http.Request) {
	var organization Organization
	if err := json.NewDecoder(r.Body).Decode(&organization); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		log.Printf("Error decoding organization: %v", err)
		return
	}

	if err := initializers.DB.Create(&organization).Error; err != nil {
		http.Error(w, "Failed to create organization: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Error creating organization: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Organization '" + organization.Name + "' created successfully"))
}

func GetOrganization(w http.ResponseWriter, r *http.Request) {
	var organization Organization
	if err := initializers.DB.First(&organization, r.FormValue("id")).Error; err != nil {
		http.Error(w, "Organization not found: "+err.Error(), http.StatusNotFound)
		log.Printf("Error retrieving organization: %v", err)
		return
	}

	if err := json.NewEncoder(w).Encode(organization); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Error encoding organization: %v", err)
	}

	w.WriteHeader(http.StatusOK)
}

func GetAllOrganizations(w http.ResponseWriter, r *http.Request) {
	var organizations []Organization
	if err := initializers.DB.Find(&organizations).Error; err != nil {
		http.Error(w, "Failed to retrieve organizations: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Error retrieving organizations: %v", err)
		return
	}

	if err := json.NewEncoder(w).Encode(organizations); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Error encoding organizations: %v", err)
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateOrganization(w http.ResponseWriter, r *http.Request) {
	var organization Organization
	// Use PUT to handle update
	if err := json.NewDecoder(r.Body).Decode(&organization); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		log.Printf("Error decoding organization: %v", err)
		return
	}

	if organization.ID == 0 {
		http.Error(w, "Organization ID is required", http.StatusBadRequest)
		log.Printf("Organization ID is missing")
		return
	}

	if err := initializers.DB.First(&organization, organization.ID).Error; err != nil {
		http.Error(w, "Organization not found: "+err.Error(), http.StatusNotFound)
		log.Printf("Error retrieving organization: %v", err)
		return
	}

	if err := initializers.DB.Save(&organization).Error; err != nil {
		http.Error(w, "Failed to update organization: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Error updating organization: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Organization '" + organization.Name + "' updated successfully"))
}

func DeleteOrganization(w http.ResponseWriter, r *http.Request) {
	var organization Organization
	if err := initializers.DB.First(&organization, r.FormValue("id")).Error; err != nil {
		http.Error(w, "Organization not found: "+err.Error(), http.StatusNotFound)
		log.Printf("Error retrieving organization: %v", err)
		return
	}
	if err := initializers.DB.Delete(&organization).Error; err != nil {
		http.Error(w, "Failed to delete organization: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Error deleting organization: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Organization '" + organization.Name + "' deleted successfully"))
}


func GetOrganizationUsers(w http.ResponseWriter, r *http.Request) {
	var organization Organization
	if err := initializers.DB.Preload("Users").First(&organization, r.FormValue("id")).Error; err != nil {
		http.Error(w, "Organization not found: "+err.Error(), http.StatusNotFound)
		log.Printf("Error retrieving organization: %v", err)
		return
	}

	if err := json.NewEncoder(w).Encode(organization.Users); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Error encoding users: %v", err)
	}
	w.WriteHeader(http.StatusOK)
}

func GetOrganizationCandidates(w http.ResponseWriter, r *http.Request) {
	var organization Organization
	if err := initializers.DB.Preload("Candidates").First(&organization, r.FormValue("id")).Error; err != nil {
		http.Error(w, "Organization not found: "+err.Error(), http.StatusNotFound)
		log.Printf("Error retrieving organization: %v", err)
		return
	}

	if err := json.NewEncoder(w).Encode(organization.Candidates); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Error encoding candidates: %v", err)
	}
	w.WriteHeader(http.StatusOK)
}

func GetOrganizationAssessments(w http.ResponseWriter, r *http.Request) {
	var organization Organization
	if err := initializers.DB.Preload("Assessments").First(&organization, r.FormValue("id")).Error; err != nil {
		http.Error(w, "Organization not found: "+err.Error(), http.StatusNotFound)
		log.Printf("Error retrieving organization: %v", err)
		return
	}

	if err := json.NewEncoder(w).Encode(organization.Assessments); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Error encoding assessments: %v", err)
	}
	w.WriteHeader(http.StatusOK)
}
