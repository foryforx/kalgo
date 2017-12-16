package handlers

import (
	// "github.com/gin-gonic/contrib/sessions"
	"fmt"
	"github.com/gin-gonic/gin"
	// "github.com/karuppaiah/kalgo/database"
	"github.com/karuppaiah/kalgo/models"
	"net/http"
	// "golang.org/x/oauth2"
	// "golang.org/x/oauth2/google"
	// "strconv"
)

type AddressHandler struct{}

// AddressHandler handles the address fetch.
func (hndlr AddressHandler) GetAllAddress(c *gin.Context) {
	// session := sessions.Default(c)
	// userID := session.Get("user-id")
	addressM := new(models.AddressModel)
	if address, err := addressM.AllAddress(); err != nil {
		c.AbortWithStatus(404)
	} else {
		// c.JSON(200, address)
		// fmt.Println(address)
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"addressList.html",
			// Pass the data that the page uses
			gin.H{
				"title":   "Address Page",
				"payload": address,
			})
	}
}

// AddressHandler handles the address fetch.
func (hndlr AddressHandler) GetAddress(c *gin.Context) {
	// session := sessions.Default(c)
	// userID := session.Get("user-id")
	addressM := new(models.AddressModel)
	addressID := c.Param("address_id")
	fmt.Println("hello")
	fmt.Println(addressID)
	if addressID != "" {
		if address, err := addressM.GetAddressById(addressID); err != nil {
			// If the address is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			// c.JSON(200, address)
			// fmt.Println(address)
			c.HTML(
				// Set the HTTP status to 200 (OK)
				http.StatusOK,
				// Use the index.html template
				"address.html",
				// Pass the data that the page uses
				gin.H{
					"title":   "Address Page",
					"payload": address,
				})
		}
	} else {
		// If an invalid address ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}

}

func (hndlr AddressHandler) AddAddress(c *gin.Context) {

}
