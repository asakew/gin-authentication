package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TemplatesRoutes(r *gin.Engine) {
	r.Static("/css", "./web/assets/css")
	r.Static("/js", "./web/assets/js")
	r.LoadHTMLGlob("web/templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
