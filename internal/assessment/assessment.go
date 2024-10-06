package assessment

import (
	"github.com/Dhamodharancs17/findskilleddev/internal/organization"
	"gorm.io/gorm"
)

type Assessment struct {
	gorm.Model
	OrganizationId uint                      `json:"organization_id" gorm:"not null"`
	Organization   organization.Organization `json:"organization" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	Name           string                    `json:"name" gorm:"not null"`
	JobName        string                    `json:"job_name" gorm:"not null"`
	ScoreLimit     int                       `json:"score_limit" gorm:"not null;default:0"`
	UniqueHash     string                    `json:"unique_hash" gorm:"not null;uniqueIndex"`
}
