package database

import (
	"fmt"
	"github.com/jinzhu/gorm"

	//"github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var DBCon *gorm.DB
var err error

//Init ...
func Init() *gorm.DB {
	DBCon, err = gorm.Open("postgres", "host=localhost user=karups dbname=kalgo sslmode=disable password=")
	if err != nil {
		fmt.Println(err)
	}

	return DBCon

}

//GetDB ...
func GetDB() *gorm.DB {
	return DBCon

}
