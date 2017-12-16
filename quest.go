package main

import (
	// "fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "net/http"
	// "runtime"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	//"fmt"
	"github.com/karuppaiah/kalgo/database"
	// "github.com/karuppaiah/kalgo/handlers"
	"github.com/karuppaiah/kalgo/middleware"
	"github.com/karuppaiah/kalgo/models"
	"github.com/karuppaiah/kalgo/routes"
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
	database.DBCon.AutoMigrate(&models.Address{})
	database.DBCon.AutoMigrate(&models.User{})
	database.DBCon.AutoMigrate(&models.Customer{})
	database.DBCon.AutoMigrate(&models.CustAddMap{})

	router := gin.Default()
	// store := sessions.NewCookieStore([]byte(handlers.RandToken(64)))
	// store.Options(sessions.Options{
	// 	Path:   "/",
	// 	MaxAge: 86400 * 7,
	// })
	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	router.Use(sessions.Sessions("kalgo-session", store))

	router.Use(middleware.CORSMiddleware())

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(sessions.Sessions("kalgo-session", store))
	router.Static("/css", "./static/css")
	router.Static("/img", "./static/img")
	router.LoadHTMLGlob("templates/*")
	routes.InitializeRoutes(router)

	// v1 := router.Group("/v1")
	// {
	// 	/*** START USER ***/
	// 	// user := new(controllers.UserController)

	// 	// v1.POST("/user/signin", user.Signin)
	// 	// v1.POST("/user/signup", user.Signup)
	// 	// v1.GET("/user/signout", user.Signout)

	// 	/*** START Article ***/
	// 	// article := new(controllers.ArticleController)

	// 	// v1.POST("/article", article.Create)
	// 	// v1.GET("/articles", article.All)
	// 	// v1.GET("/article/:id", article.One)
	// 	// v1.PUT("/article/:id", article.Update)
	// 	// v1.DELETE("/article/:id", article.Delete)
	// }
	// auth := new(handlers.AuthHandler)
	// index := new(handlers.IndexHandler)
	// address := new(handlers.AddressHandler)
	// router.GET("/", index.HomeHandler)
	// router.GET("/login", auth.LoginHandler)
	// router.GET("/auth", auth.AuthenticateHandler)
	// router.GET("/address", address.GetAllAddress)

	// authorized := router.Group("/battle")
	// authorized.Use(middleware.AuthorizeRequest())
	// {
	// 	authorized.GET("/field", handlers.FieldHandler)
	// }

	router.Run("localhost:9090")
}
