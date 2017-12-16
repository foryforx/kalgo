package models

import (
// "github.com/fatih/color"
// "github.com/jinzhu/gorm"
// "log"
// "math/rand"
//"time"
)

// Address and it's properties.
type Customer struct {
	Email         string `gorm:"primary_key" json:"email"`
	RootUniqueId  string `gorm:"primary_key;" json:"root_unique_id"`
	EmailVerified string `gorm:"not null" json:"email_verified"`
	Enabled       string `json:"enabled"`
	Sub           string `json:"sub"`
	GivenName     string `gorm:"not null;" json:"given_name"`
	FamilyName    string `gorm:"not null;" json:"last_name"`
	Profile       string `gorm:"not null" json:"profile"`
	Picture       string `gorm:"not null" json:"picture"`
	Gender        bool   `json:"gender"`
	Primary       bool   `gorm:"not null" json:"primary"`
	Type          string `gorm:"not null" json:"type"`
}
