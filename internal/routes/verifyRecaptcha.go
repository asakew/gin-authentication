package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

var secretKey = os.Getenv("RECAPTCHA_SECRET")

type RecaptchaResponse struct {
	Success     bool   `json:"success"`
	ChallengeTs string `json:"challenge_ts"`
	Hostname    string `json:"hostname"`
}

func verifyRecaptcha(token string) (bool, error) {
	response, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", url.Values{
		"secret":   {secretKey},
		"response": {token},
	})
	if err != nil {
		return false, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	var recaptchaResponse RecaptchaResponse
	if err := json.NewDecoder(response.Body).Decode(&recaptchaResponse); err != nil {
		return false, err
	}
	return recaptchaResponse.Success, nil
}

func Register(c *gin.Context) {
	token := c.PostForm("g-recaptcha-response")
	if valid, err := verifyRecaptcha(token); err != nil || !valid {
		log.Println("Invalid reCAPTCHA.")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reCAPTCHA"})
		return
	}
	// Handle registration logic here
	c.JSON(http.StatusOK, gin.H{"message": "Registration successful!"})
}

func Login(c *gin.Context) {
	token := c.PostForm("g-recaptcha-response")
	if valid, err := verifyRecaptcha(token); err != nil || !valid {
		log.Println("Invalid reCAPTCHA.")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reCAPTCHA"})
		return
	}
	// Handle login logic here
	c.JSON(http.StatusOK, gin.H{"message": "Login successful!"})
}

func SetupAuth(router *gin.Engine) {
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	router.POST("/login", Login)
	router.POST("/register", Register)
}
