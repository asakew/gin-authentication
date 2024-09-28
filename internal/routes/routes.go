package routes

import (
	"github.com/gin-gonic/gin"
)

func TemplatesRoutes(r *gin.Engine) {
	r.Static("/css", "./web/assets/css")
	r.Static("/js", "./web/assets/js")
	r.LoadHTMLGlob("web/templates/*")
}
