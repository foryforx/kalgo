package middleware

import (
	"fmt"
	// jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthorizeRequest is used to authorize a request for a certain end-point group.
func AuthorizeRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		v := session.Get("user-id")
		if v == nil {
			c.HTML(http.StatusUnauthorized, "error.tmpl", gin.H{"message": "Please login."})
			c.Abort()
		}
		c.Next()
	}
}

//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

// func TokenAuthMiddleware(c *gin.Context) {

// 	token, err := jwt.ParseFromRequest(c.Request, func(token *jwt.Token) (interface{}, error) {
// 		b := ([]byte(JWTKEY))

// 		return b, nil
// 	})

// 	if err != nil {
// 		abortWithError(c, http.StatusUnauthorized, "Invaild User Token")
// 		return
// 	}

// 	if token.Claims[USERID] != nil {
// 		db := c.MustGet(mongodb.DBMiddlewareName).(*mgo.Database)
// 		user, getUserErr := GetUserByEmail(token.Claims[USERID].(string), db)
// 		if getUserErr != nil {
// 			abortWithError(c, http.StatusUnauthorized, "Invalid User Token")
// 			return
// 		}
// 		c.Set(User, user)
// 	} else {
// 		abortWithError(c, http.StatusUnauthorized, "Invalid User Token")
// 		return
// 	}

// 	c.Next()
// }
