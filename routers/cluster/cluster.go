package clusters

import (
	"kubez_project/controllers/cluster"

	"github.com/gin-gonic/gin"
)

func Add(clusterGroup *gin.RouterGroup) {
	clusterGroup.POST("/add", cluster.Add)
}

func Delete(clusterGroup *gin.RouterGroup) {
	clusterGroup.POST("/delete", cluster.Delete)
}

func Update(clusterGroup *gin.RouterGroup) {
	clusterGroup.POST("/update", cluster.Update)
}

func List(clusterGroup *gin.RouterGroup) {
	clusterGroup.GET("/list", cluster.List)
}

func Get(clusterGroup *gin.RouterGroup) {
	clusterGroup.GET("/get", cluster.Get)
}

func RegisterSubRouters(group *gin.RouterGroup) {
	// 第二层 : /api/cluster
	clusterGroup := group.Group("/cluster")
	// 接口 : /api/cluster/add
	Add(clusterGroup)
	// 接口 : /api/cluster/delete
	Delete(clusterGroup)
	// 接口 : /api/cluster/update
	Update(clusterGroup)
	// 接口 : /api/cluster/list
	List(clusterGroup)
	// 接口 : /api/cluster/get
	Get(clusterGroup)
}
