package cronjobs

import (
	"kubez_project/controllers/cronjob"

	"github.com/gin-gonic/gin"
)

// RegisterSubRouters 注册 CronJob 相关的路由
func RegisterSubRouters(group *gin.RouterGroup) {
	cronJobsGroup := group.Group("/cronjob")
	cronJobsGroup.POST("/add", cronjob.Add)
	cronJobsGroup.POST("/delete", cronjob.Delete)
	cronJobsGroup.POST("/update", cronjob.Update)
	cronJobsGroup.POST("/list", cronjob.List)
	cronJobsGroup.POST("/get", cronjob.Get)
}
