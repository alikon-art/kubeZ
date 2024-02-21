package secrets

import (
	"kubez_project/controllers/secret"

	"github.com/gin-gonic/gin"
)

// RegisterSubRouters 在这里注册Secret的路由
func RegisterSubRouters(group *gin.RouterGroup) {
	// 第二层 : /api/secrets
	secretsGroup := group.Group("/secret")
	// 接口 : /api/secrets/create
	secretsGroup.POST("/add", secret.Create)
	// 接口 : /api/secrets/delete
	secretsGroup.POST("/delete", secret.Delete)
	// 接口 : /api/secrets/update
	secretsGroup.POST("/update", secret.Update)
	// 接口 : /api/secrets/list
	secretsGroup.POST("/list", secret.List)
	// 接口 : /api/secrets/get
	secretsGroup.POST("/get", secret.Get)
}
