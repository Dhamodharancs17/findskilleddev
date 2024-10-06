package candidate

import (
	"github.com/Dhamodharancs17/findskilleddev/internal/organization"
	"gorm.io/gorm"
)

type Candidate struct {
	gorm.Model
	OrganizationId  uint                      `json:"organization_id" gorm:"not null"`
	Organization    organization.Organization `json:"organization" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	Name            string                    `json:"name" gorm:"not null"`
	Email           string                    `json:"email" gorm:"not null;uniqueIndex"`
	EmailVerified   bool                      `json:"email_verified" gorm:"default:false"`
	Phone           string                    `json:"phone" gorm:"not null;uniqueIndex"`
	NoDualScreen    bool                      `json:"no_dual_screen" gorm:"not null;default:false"`
	Camera          bool                      `json:"camera" gorm:"not null;default:false"`
	Microphone      bool                      `json:"microphone" gorm:"not null;default:false"`
	LookedOutsideAI int                       `json:"looked_outside_ai" gorm:"not null;default:0"`
	ResumeURL       string                    `json:"resume_url" gorm:"size:512"`
	Status          string                    `json:"status" gorm:"not null"`
}
