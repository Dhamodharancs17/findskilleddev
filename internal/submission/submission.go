package submission

import (
	"time"

	"github.com/Dhamodharancs17/findskilleddev/internal/candidate"
	"github.com/Dhamodharancs17/findskilleddev/internal/question"
	"gorm.io/gorm"
)

type Submission struct {
	gorm.Model
	CandidateID      uint                `json:"candidate_id" gorm:"not null;index"`
	Candidate        candidate.Candidate `json:"candidate" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	Answer           string              `json:"answers" gorm:"type:text"`
	RecordURL        string              `json:"record_url" gorm:"size:512"`
	CodingResult     bool                `json:"coding_result" gorm:"not null;default:false"`
	NumberOfCodeRuns int                 `json:"number_of_code_runs" gorm:"not null;default:0"`
	Duration         int                 `json:"duration" gorm:"not null;default:0"`
	StartAt          time.Time           `json:"start_at" gorm:"not null"`
	EndAt            time.Time           `json:"end_at" gorm:"not null"`
	QuestionID       uint                `json:"question_id" gorm:"not null"`
	Question         question.Question   `json:"question" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
}

// Calculating the duration of the submission

func (s *Submission) CalculateDuration() {
	s.Duration = int(s.EndAt.Sub(s.StartAt).Minutes())
}
