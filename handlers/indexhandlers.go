package handlers

import (
	"net/http"
	// "github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	// "github.com/karuppaiah/kalgo/database"
	// "github.com/karuppaiah/kalgo/models"
	// "golang.org/x/oauth2"
	// "golang.org/x/oauth2/google"
)

type IndexHandler struct{}

// IndexHandler handels /.
func (hndlr IndexHandler) HomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page"})
}
