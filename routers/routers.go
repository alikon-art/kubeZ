package routers

import (
<<<<<<< HEAD
	// "kubez_project/routers/auth"
=======
	"kubez_project/routers/auth"
>>>>>>> 6512db693d369297878d5d9833b4c1d8f000e958
	"kubez_project/routers/cluster"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(g *gin.Engine) {
	// 第一层 : /api
	apiGroup := g.Group("/api")
<<<<<<< HEAD
	// auth.RegisterSubRouters(apiGroup)
=======
	auth.RegisterSubRouters(apiGroup)
>>>>>>> 6512db693d369297878d5d9833b4c1d8f000e958
	clusters.RegisterSubRouters(apiGroup)
}
