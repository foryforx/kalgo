package main

import (
	//"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	//"fmt"
	"kalgo/database"
	"kalgo/handlers"
	"kalgo/middleware"
	"kalgo/structs"
)

var db *gorm.DB
var err error

// // Address and it's properties.
// type Address struct {
// 	AddressID       string `json:"name"`
// 	AddressLine1    string `json:"address_line_1"`
// 	AddressLine2    string `json:"address_line_2"`
// 	AddressLine3    string `json:"address_line_3"`
// 	City            string `json:"city"`
// 	State           string `json:"state"`
// 	PostalCode      string `json:"postal_code"`
// 	IsoCountry      string `json:"country"`
// 	AddressVerified bool   `json:"address_verified"`
// }

// type Person struct {
// 	ID        uint   `json:"id"`
// 	FirstName string `json:"firstname"`
// 	LastName  string `json:"lastname"`
// 	City      string `json:"city"`
// }

// func main2() {
// 	//db, err = gorm.Open("sqlite3", "./gorm.db")
// 	db, err = gorm.Open("postgres", "host=localhost user=karups dbname=kalgo sslmode=disable password=")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer db.Close()

// 	db.AutoMigrate(&Person{})

// 	r := gin.Default()
// 	r.GET("/", GetPeople)
// 	r.GET("/people/", GetPeople)
// 	r.GET("/people/:id", GetPerson)
// 	r.POST("/people", CreatePerson)
// 	r.PUT("/people/:id", UpdatePerson)
// 	r.DELETE("/people/:id", DeletePerson)
// 	r.Run("localhost:9090")

// }

// func GetPerson(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var person Person
// 	if err := db.Where("id= ?", id).First(&person).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)

// 	} else {
// 		c.JSON(200, person)
// 	}
// }

// func GetPeople(c *gin.Context) {
// 	var people []Person
// 	if err := db.Find(&people).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		c.JSON(200, people)
// 	}
// }

// func CreatePerson(c *gin.Context) {
// 	var person Person
// 	c.BindJSON(&person)
// 	db.Create(&person)
// 	c.JSON(200, person)
// }

// func UpdatePerson(c *gin.Context) {
// 	var person Person
// 	id := c.Params.ByName("id")

// 	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	}
// 	c.BindJSON(&person)

// 	db.Save(&person)
// 	c.JSON(200, person)
// }

// func DeletePerson(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var person Person
// 	d := db.Where("id=?", id).Delete(&person)
// 	fmt.Println(d)
// 	c.JSON(200, gin.H{"id #" + id: "deleted"})
// }

// AddressHandler handles the address fetch.
// func GetAddress(c *gin.Context) {
// 	//session := sessions.Default(c)
// 	// userID := session.Get("user-id")
// 	var address []structs.Address

// 	if err := database.DBCon.Find(&address).Error; err != nil {
// 		c.AbortWithStatus(404)
// 	} else {
// 		c.JSON(200, address)
// 	}
// }

func main() {
	database.DBCon = database.Init()
	defer database.DBCon.Close()
	database.DBCon.AutoMigrate(&structs.Address{})
	router := gin.Default()
	store := sessions.NewCookieStore([]byte(handlers.RandToken(64)))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
	})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(sessions.Sessions("goquestsession", store))
	router.Static("/css", "./static/css")
	router.Static("/img", "./static/img")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", handlers.IndexHandler)
	router.GET("/login", handlers.LoginHandler)
	router.GET("/auth", handlers.AuthHandler)
	router.GET("/address", handlers.AddressHandler)

	authorized := router.Group("/battle")
	authorized.Use(middleware.AuthorizeRequest())
	{
		authorized.GET("/field", handlers.FieldHandler)
	}

	router.Run("localhost:9090")
}
