package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HTMLRendering(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
