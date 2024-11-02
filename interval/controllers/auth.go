package controllers

import (
	db "appGin/interval/database"
	"appGin/interval/models"
	"context"
	"gorm.io/gorm"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

var UserDB *gorm.DB

func GoogleLogin(c *gin.Context) {
	url := googleOAuthConfig().AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	token, err := googleOAuthConfig().Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
		return
	}

	oauth2Service, err := oauth2.NewService(context.Background(), option.WithTokenSource(oauth2.StaticTokenSource(token)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create oauth2 service"})
		return
	}

	userinfo, err := oauth2Service.Userinfo.Get().Do()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}

	var user models.User
	user, err = UserDB.GetUserByEmail(userinfo.Email)
	if err != nil {
		user = models.User{
			Email:     userinfo.Email,
			FirstName: userinfo.GivenName,
			LastName:  userinfo.FamilyName,
		}
		err := models.CreateUser(&user)
		if err != nil {
			return
		}
	}

	// Generate JWT token
	jwtToken := generateJWT(user)

	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}

func googleOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}
