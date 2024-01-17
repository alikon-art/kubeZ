package statefulsets

import (
	"kubez_project/controllers/statefulset"

	"github.com/gin-gonic/gin"
)

// 在这里注册StatefulSet的路由
func RegisterSubRouters(group *gin.RouterGroup) {
	// 第二层 : /api/statefulsets
	statefulSetsGroup := group.Group("/statefulset")
	// 接口 : /api/statefulsets/add
	statefulSetsGroup.POST("/add", statefulset.Add)
	// 接口 : /api/statefulsets/delete
	statefulSetsGroup.POST("/delete", statefulset.Delete)
	// 接口 : /api/statefulsets/update
	statefulSetsGroup.POST("/update", statefulset.Update)
	// 接口 : /api/statefulsets/list
	statefulSetsGroup.POST("/list", statefulset.List)
	// 接口 : /api/statefulsets/get
	statefulSetsGroup.POST("/get", statefulset.Get)
}
