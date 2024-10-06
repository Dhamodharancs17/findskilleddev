package organization

import (
	// assessment "github.com/Dhamodharancs17/findskilleddev/internal/assessment"
	"github.com/Dhamodharancs17/findskilleddev/internal/assessment"
	"github.com/Dhamodharancs17/findskilleddev/internal/candidate"
	"github.com/Dhamodharancs17/findskilleddev/internal/user"
	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	Name        string                  `json:"name" gorm:"not null;size:255"`                                    // Added size constraint
	About       string                  `json:"about" gorm:"type:text"`                                           // Specified type as text for potentially long content
	LogoUrl     string                  `json:"logo_url" gorm:"size:512"`                                         // Added size constraint for URL
	Users       []user.User             `json:"users" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;foreignKey: OrganizationID, reference: ID"`       // Added constraint for user deletion
	Candidates  []candidate.Candidate   `json:"candidates" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;foreignKey: OrganizationID, reference: ID"`  // Added constraint for candidate deletion
	Assessments []assessment.Assessment `json:"assessments" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE; foreignKey: OrganizationID, reference: ID"` // Added constraint for assessment deletion
}
