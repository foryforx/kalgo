package models

import (
// "github.com/fatih/color"
// "github.com/jinzhu/gorm"
// "log"
// "math/rand"
//"time"
)

// Address and it's properties.
type CustAddMap struct {
	CustRootUniqueId string `gorm:"not null;" json:"cust_root_unique_id"`
	AddressID        string `gorm:"not null" json:"address_id"`
	AddressCode      string `gorm:"not null" json:"address_code"`
}
