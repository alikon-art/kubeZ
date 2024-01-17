// 在这里注册PVC的路由
package pvcs

import (
	"kubez_project/controllers/pvc"

	"github.com/gin-gonic/gin"
)

func RegisterSubRouters(group *gin.RouterGroup) {
	// 第二层 : /api/pvcs
	pvcsGroup := group.Group("/pvc")
	// 接口 : /api/pvcs/add
	pvcsGroup.POST("/add", pvc.Add)
	// 接口 : /api/pvcs/delete
	pvcsGroup.POST("/delete", pvc.Delete)
	// 接口 : /api/pvcs/update
	pvcsGroup.POST("/update", pvc.Update)
	// 接口 : /api/pvcs/list
	pvcsGroup.POST("/list", pvc.List)
	// 接口 : /api/pvcs/get
	pvcsGroup.POST("/get", pvc.Get)
}
