package question

import (
	"github.com/Dhamodharancs17/findskilleddev/internal/submission"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	AssessmentID   uint                  `json:"assessment_id" gorm:"not null"`
	Question       string                `json:"question" gorm:"type:text; not null"`
	ExpectedAnswer string                `json:"expected_answer" gorm:"type:text; not null"`
	TestCase       string                `json:"test_case" gorm:"type:json; not null"`
	Language       string                `json:"language" gorm:"not null"`
	Submissions   []submission.Submission `json:"submissions" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
}
