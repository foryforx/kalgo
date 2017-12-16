package models

// import (
// 	"crypto/rand"
// 	"encoding/base64"
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/gin-gonic/contrib/sessions"
// 	"github.com/gin-gonic/gin"
// 	// "github.com/karuppaiah/kalgo/database"
// 	"golang.org/x/oauth2"
// 	"golang.org/x/oauth2/google"
// )

// User is a retrieved and authentiacted user.
type User struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
}
