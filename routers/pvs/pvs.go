package pvs

import (
	"kubez_project/controllers/pv"

	"github.com/gin-gonic/gin"
)

// 在这里注册pv的路由
func RegisterSubRouters(group *gin.RouterGroup) {
	// 第二层 : /api/pvs
	pvGroup := group.Group("/pv")
	// 接口 : /api/pvs/add
	pvGroup.POST("/add", pv.Add)
	// 接口 : /api/pvs/delete
	pvGroup.POST("/delete", pv.Delete)
	// 接口 : /api/pvs/update
	pvGroup.POST("/update", pv.Update)
	// 接口 : /api/pvs/list
	pvGroup.POST("/list", pv.List)
	// 接口 : /api/pvs/get
	pvGroup.POST("/get", pv.Get)
}
