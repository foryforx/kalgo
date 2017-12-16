package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karuppaiah/kalgo/handlers"
	// "github.com/karuppaiah/kalgo/middleware"
)

func InitializeRoutes(router *gin.Engine) {
	auth := new(handlers.AuthHandler)
	index := new(handlers.IndexHandler)
	address := new(handlers.AddressHandler)
	router.GET("/", index.HomeHandler)
	router.GET("/login", auth.LoginHandler)
	router.GET("/auth", auth.AuthenticateHandler)
	// router.GET("/address", address.GetAllAddress)
	// router.GET("/address/view/:address_id", address.GetAddress)

	private_v1 := router.Group("/api/v1")
	// private_v1.Use(middleware.TokenAuthMiddleware)
	private_v1.GET("/address", address.GetAllAddress)
	private_v1.GET("/address/view/:address_id", address.GetAddress)
}
