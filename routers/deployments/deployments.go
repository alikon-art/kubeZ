package deployments

import (
	"kubez_project/controllers/deployment"

	"github.com/gin-gonic/gin"
)

func Add(deploymentsGroup *gin.RouterGroup) {
	deploymentsGroup.POST("/add", deployment.Add)
}

func Delete(deploymentsGroup *gin.RouterGroup) {
	deploymentsGroup.POST("/delete", deployment.Delete)
}

func Update(deploymentsGroup *gin.RouterGroup) {
	deploymentsGroup.POST("/update", deployment.Update)
}

func List(deploymentsGroup *gin.RouterGroup) {
	deploymentsGroup.POST("/list", deployment.List)
}

func Get(deploymentsGroup *gin.RouterGroup) {
	deploymentsGroup.POST("/get", deployment.Get)
}

func RegisterSubRouters(group *gin.RouterGroup) {
	// 第二层 : /api/deployments
	deploymentsGroup := group.Group("/deployment")
	// 接口 : /api/deployments/add
	Add(deploymentsGroup)
	// 接口 : /api/deployments/delete
	Delete(deploymentsGroup)
	// 接口 : /api/deployments/update
	Update(deploymentsGroup)
	// 接口 : /api/deployments/list
	List(deploymentsGroup)
	// 接口 : /api/deployments/get
	Get(deploymentsGroup)
}
