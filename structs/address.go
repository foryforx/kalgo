package structs

import (
// "github.com/fatih/color"
// "github.com/jinzhu/gorm"
// "log"
// "math/rand"
// "time"
)

// Address and it's properties.
type Address struct {
	AddressID       string `gorm:"primary_key" json:"name"`
	AddressLine1    string `gorm:"not null;" json:"address_line_1"`
	AddressLine2    string `gorm:"not null" json:"address_line_2"`
	AddressLine3    string `json:"address_line_3"`
	City            string `gorm:"not null;" json:"city"`
	State           string `gorm:"not null;" json:"state"`
	PostalCode      string `gorm:"not null"`
	IsoCountry      string `gorm:"not null" json:"country"`
	AddressVerified bool   `json:"address_verified"`
}
