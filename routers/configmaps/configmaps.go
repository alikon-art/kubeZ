package configmaps

import (
	"kubez_project/controllers/configmap"

	"github.com/gin-gonic/gin"
)

func Add(configMapsGroup *gin.RouterGroup) {
	configMapsGroup.POST("/add", configmap.Add)
}

func Delete(configMapsGroup *gin.RouterGroup) {
	configMapsGroup.POST("/delete", configmap.Delete)
}

func Update(configMapsGroup *gin.RouterGroup) {
	configMapsGroup.POST("/update", configmap.Update)
}

func List(configMapsGroup *gin.RouterGroup) {
	configMapsGroup.POST("/list", configmap.List)
}

func Get(configMapsGroup *gin.RouterGroup) {
	configMapsGroup.POST("/get", configmap.Get)
}

func RegisterSubRouters(group *gin.RouterGroup) {
	configMapsGroup := group.Group("/configmap")
	Add(configMapsGroup)
	Delete(configMapsGroup)
	Update(configMapsGroup)
	List(configMapsGroup)
	Get(configMapsGroup)
}
