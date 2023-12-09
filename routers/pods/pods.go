package pods

import (
	"kubez_project/controllers/pod"

	"github.com/gin-gonic/gin"
)

func Add(podsGroup *gin.RouterGroup) {
	podsGroup.POST("/add", pod.Add)
}

func Delete(podsGroup *gin.RouterGroup) {
	podsGroup.POST("/delete", pod.Delete)
}

func Update(podsGroup *gin.RouterGroup) {
	podsGroup.POST("/update", pod.Update)
}

func List(podsGroup *gin.RouterGroup) {
	podsGroup.POST("/list", pod.List)
}

func Get(podsGroup *gin.RouterGroup) {
	podsGroup.POST("/get", pod.Get)
}

func RegisterSubRouters(group *gin.RouterGroup) {
	// 第二层 : /api/pods
	podsGroup := group.Group("/pods")
	// 接口 : /api/pods/add
	Add(podsGroup)
	// 接口 : /api/pods/delete
	Delete(podsGroup)
	// 接口 : /api/pods/update
	Update(podsGroup)
	// 接口 : /api/pods/list
	List(podsGroup)
	// 接口 : /api/pods/get
	Get(podsGroup)
}
