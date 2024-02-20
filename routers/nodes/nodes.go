package nodes

import (
	"kubez_project/controllers/node"

	"github.com/gin-gonic/gin"
)

func Add(nodeGroup *gin.RouterGroup) {
	nodeGroup.POST("/add", node.Add)
}

func Delete(nodeGroup *gin.RouterGroup) {
	nodeGroup.POST("/delete", node.Delete)
}

func Update(nodeGroup *gin.RouterGroup) {
	nodeGroup.POST("/update", node.Update)
}

func List(nodeGroup *gin.RouterGroup) {
	nodeGroup.POST("/list", node.List)
}

func Get(nodeGroup *gin.RouterGroup) {
	nodeGroup.POST("/get", node.Get)
}

func RegisterSubRouters(group *gin.RouterGroup) {
	// 第二层 : /api/node
	nodeGroup := group.Group("/node")
	// 接口 : /api/node/add
	Add(nodeGroup)
	// 接口 : /api/node/delete
	Delete(nodeGroup)
	// 接口 : /api/node/update
	Update(nodeGroup)
	// 接口 : /api/node/list
	List(nodeGroup)
	// 接口 : /api/node/get
	Get(nodeGroup)
}
