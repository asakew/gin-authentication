package handlers

import (
	"appGin/internal/db"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VerifyOTP(c *gin.Context) {
	email := c.Query("email")
	otp := c.Query("otp")

	storedOtp, err := db.RedisClient.Get(context.Background(), email).Result()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
		return
	}

	if storedOtp != otp {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
}
