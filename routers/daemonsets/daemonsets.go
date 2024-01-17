package daemonsets

import (
	"kubez_project/controllers/daemonset"

	"github.com/gin-gonic/gin"
)

// RegisterSubRouters 注册DaemonSet相关的路由
func RegisterSubRouters(group *gin.RouterGroup) {
	// 第二层 : /api/daemonsets
	daemonsetsGroup := group.Group("/daemonset")

	// 接口 : /api/daemonsets/add
	daemonsetsGroup.POST("/add", daemonset.Add)

	// 接口 : /api/daemonsets/delete
	daemonsetsGroup.POST("/delete", daemonset.Delete)

	// 接口 : /api/daemonsets/update
	daemonsetsGroup.POST("/update", daemonset.Update)

	// 接口 : /api/daemonsets/list
	daemonsetsGroup.POST("/list", daemonset.List)

	// 接口 : /api/daemonsets/get
	daemonsetsGroup.POST("/get", daemonset.Get)
}
