package routers

import (
	"kubez_project/routers/auth"
	"kubez_project/routers/clusters"
	"kubez_project/routers/namespaces"
	"kubez_project/routers/pods"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(g *gin.Engine) {
	// 第一层 : /api
	apiGroup := g.Group("/api")
	auth.RegisterSubRouters(apiGroup)
	clusters.RegisterSubRouters(apiGroup)
	namespaces.RegisterSubRouters(apiGroup)
	pods.RegisterSubRouters(apiGroup)

}
