package ingresses

import (
	"kubez_project/controllers/ingress"

	"github.com/gin-gonic/gin"
)

func Add(ingressesGroup *gin.RouterGroup) {
	ingressesGroup.POST("/add", ingress.Add)
}

func Delete(ingressesGroup *gin.RouterGroup) {
	ingressesGroup.POST("/delete", ingress.Delete)
}

func Update(ingressesGroup *gin.RouterGroup) {
	ingressesGroup.POST("/update", ingress.Update)
}

func List(ingressesGroup *gin.RouterGroup) {
	ingressesGroup.POST("/list", ingress.List)
}

func Get(ingressesGroup *gin.RouterGroup) {
	ingressesGroup.POST("/get", ingress.Get)
}

func RegisterSubRouters(group *gin.RouterGroup) {
	// 第二层 : /api/ingresses
	ingressesGroup := group.Group("/ingress")
	// 接口 : /api/ingresses/add
	Add(ingressesGroup)
	// 接口 : /api/ingresses/delete
	Delete(ingressesGroup)
	// 接口 : /api/ingresses/update
	Update(ingressesGroup)
	// 接口 : /api/ingresses/list
	List(ingressesGroup)
	// 接口 : /api/ingresses/get
	Get(ingressesGroup)
}
