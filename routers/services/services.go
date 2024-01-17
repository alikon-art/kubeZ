package services

import (
	"kubez_project/controllers/service"

	"github.com/gin-gonic/gin"
)

// RegisterSubRouters 注册服务相关的子路由
func RegisterSubRouters(group *gin.RouterGroup) {
	// 第二层 : /api/services
	servicesGroup := group.Group("/service")
	// 接口 : /api/services/add
	servicesGroup.POST("/add", service.Add)
	// 接口 : /api/services/delete
	servicesGroup.POST("/delete", service.Delete)
	// 接口 : /api/services/update
	servicesGroup.POST("/update", service.Update)
	// 接口 : /api/services/list
	servicesGroup.POST("/list", service.List)
	// 接口 : /api/services/get
	servicesGroup.POST("/get", service.Get)
}
