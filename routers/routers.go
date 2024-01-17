package routers

import (
	"kubez_project/routers/auth"
	"kubez_project/routers/clusters"
	"kubez_project/routers/configmaps"
	"kubez_project/routers/cronjobs"
	"kubez_project/routers/daemonsets"
	"kubez_project/routers/deployments"
	"kubez_project/routers/ingresses"
	"kubez_project/routers/namespaces"
	"kubez_project/routers/pods"
	"kubez_project/routers/pvcs"
	"kubez_project/routers/pvs"
	"kubez_project/routers/secrets"
	"kubez_project/routers/services"
	"kubez_project/routers/statefulsets"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(g *gin.Engine) {
	// 第一层 : /api
	// 注册api组
	apiGroup := g.Group("/api")
	// 向组中添加子组
	// 后端认证
	auth.RegisterSubRouters(apiGroup)
	// 集群管理
	clusters.RegisterSubRouters(apiGroup)
	namespaces.RegisterSubRouters(apiGroup)
	// 工作负载
	pods.RegisterSubRouters(apiGroup)
	deployments.RegisterSubRouters(apiGroup)
	statefulsets.RegisterSubRouters(apiGroup)
	daemonsets.RegisterSubRouters(apiGroup)
	cronjobs.RegisterSubRouters(apiGroup)
	// 服务发布
	services.RegisterSubRouters(apiGroup)
	ingresses.RegisterSubRouters(apiGroup)
	// 配置管理
	configmaps.RegisterSubRouters(apiGroup)
	secrets.RegisterSubRouters(apiGroup)
	// 存储管理
	pvs.RegisterSubRouters(apiGroup)
	pvcs.RegisterSubRouters(apiGroup)
}
