package models

import (
	// "github.com/fatih/color"
	// "github.com/jinzhu/gorm"
	// "log"
	// "math/rand"
	// "time"
	"errors"
	"fmt"
	"github.com/karuppaiah/kalgo/database"
)

// Address and it's properties.
type Address struct {
	AddressID         string `gorm:"primary_key" json:"address_id"`
	AddressLine1      string `gorm:"not null;" json:"address_line_1"`
	AddressLine2      string `gorm:"not null" json:"address_line_2"`
	AddressLine3      string `json:"address_line_3"`
	City              string `gorm:"not null;" json:"city"`
	State             string `gorm:"not null;" json:"state"`
	PostalCode        string `gorm:"not null" json:"postal_code"`
	IsoCountry        string `gorm:"not null" json:"country"`
	AddressVerified   bool   `json:"address_verified"`
	VerificationLevel string `gorm:"not null" json:"verification_level"`
	Visibility        string `gorm:"not null" jsong:"visibility"`
}

//AddressModel
type AddressModel struct{}

func (m AddressModel) CreateAddress(addressLine1 string, addressLine2 string, addressLine3 string, city string, state string, postal_code string, isoCountry string, addressVerified bool, verificationLevel string, visibility bool) (addressId string, err error) {
	return "", nil
}

func (m AddressModel) AllAddress() (address []Address, err error) {
	err = database.DBCon.Find(&address).Error
	return address, err
}

func (m AddressModel) UpsertAddress(addId string, addressLine1 string, addressLine2 string, addressLine3 string, city string, state string, postal_code string, isoCountry string, addressVerified bool, verificationLevel string, visibility bool) (addressId string, err error) {
	return "", nil
}

func (m AddressModel) GetAddressById(addId string) (*Address, error) {
	var address Address
	fmt.Println(addId)
	err := database.DBCon.Where("address_id = ?", addId).First(&address).Error
	fmt.Println(err)
	if err == nil {
		return &address, nil
	}

	return &address, errors.New("Address not found")
}
