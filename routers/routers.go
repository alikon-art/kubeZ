package routers

import (
	// "kubez_project/routers/auth"
	"kubez_project/routers/cluster"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(g *gin.Engine) {
	// 第一层 : /api
	apiGroup := g.Group("/api")
	// auth.RegisterSubRouters(apiGroup)
	clusters.RegisterSubRouters(apiGroup)
}
