package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string `json:"name" gorm:"not null"`
	Email          string `json:"email" gorm:"not null;unique,index"`
	Password       string `json:"password" gorm:"not null"`
	Role           string `json:"role" gorm:"not null"`
	OrganizationID uint   `json:"organization_id" gorm:"not null"`
}
