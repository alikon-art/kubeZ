package namespaces

import (
	"kubez_project/controllers/namespace"

	"github.com/gin-gonic/gin"
)

func Add(namespaceGroup *gin.RouterGroup) {
	namespaceGroup.POST("/add", namespace.Add)
}

func Delete(namespaceGroup *gin.RouterGroup) {
	namespaceGroup.POST("/delete", namespace.Delete)
}

func Update(namespaceGroup *gin.RouterGroup) {
	namespaceGroup.POST("/update", namespace.Update)
}

func List(namespaceGroup *gin.RouterGroup) {
	namespaceGroup.POST("/list", namespace.List)
}

func Get(namespaceGroup *gin.RouterGroup) {
	namespaceGroup.POST("/get", namespace.Get)
}

func RegisterSubRouters(group *gin.RouterGroup) {
	// 第二层 : /api/namespace
	namespaceGroup := group.Group("/namespace")
	// 接口 : /api/namespace/add
	Add(namespaceGroup)
	// 接口 : /api/namespace/delete
	Delete(namespaceGroup)
	// 接口 : /api/namespace/update
	Update(namespaceGroup)
	// 接口 : /api/namespace/list
	List(namespaceGroup)
	// 接口 : /api/namespace/get
	Get(namespaceGroup)
}
