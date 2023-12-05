package routers

import (
	"kubez_project/routers/auth"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(g *gin.Engine) {
	// 第一层 : /api
	apigroup := g.Group("/api")

	auth.RegisterSubRouters(apigroup)
}
