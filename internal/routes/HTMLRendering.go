package routes

import (
	v1 "appGin/api/v1"
	"appGin/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HTMLRendering(r *gin.Engine) {
	r.Static("/css", "./web/assets/css")
	r.Static("/js", "./web/assets/js")
	r.LoadHTMLGlob("web/templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(200, "register.html", nil)
	})

	//r.POST("/login", handlers.Login)
	//r.POST("/register", handlers.Register)
	//r.GET("/verify-otp", func(c *gin.Context) {
	//	c.HTML(200, "verify_otp.html", nil)
	//})
	////r.POST("/verify-otp", handlers.VerifyOTP)
	//
	//auth := r.Group("/auth")
	//auth.Use(handlers.ValidateToken)
	//{
	//	auth.GET("/profile", func(c *gin.Context) {
	//		c.JSON(200, gin.H{"message": "Welcome to your profile"})
	//	})
	//}

	api := r.Group("/api/v1")
	{
		api.POST("/register", v1.Register)
		api.POST("/login", v1.Login)
	}

	auth := r.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/protected", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "You are authorized"})
		})
	}
}
