package organization

import (
	"github.com/Dhamodharancs17/findskilleddev/internal/user"
	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	Name    string    `json:"name" gorm:"not null;size:255"`                             // Added size constraint
	UserID  uint      `json:"user_id" gorm:"not null;unique;index"`                      // Added index tag explicitly
	User    user.User `json:"user" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"` // Added OnUpdate constraint
	About   string    `json:"about" gorm:"type:text"`                                    // Specified type as text for potentially long content
	LogoUrl string    `json:"logo_url" gorm:"size:512"`                                  // Added size constraint for URL
}
