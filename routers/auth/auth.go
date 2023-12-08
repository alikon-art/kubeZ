package auth

import (
	"kubez_project/controllers/auth"

	"github.com/gin-gonic/gin"
)

func login(authGroup *gin.RouterGroup) {
	authGroup.POST("/login", auth.Login)
}

func logout(authGroup *gin.RouterGroup) {
	authGroup.POST("/logout", auth.Logout)
}

func RegisterSubRouters(group *gin.RouterGroup) {
	// 第二层 : /api/auth
	authGroup := group.Group("/auth")
	// 接口 : /api/auth/login
	login(authGroup)
	// 接口 : /api/auth/logout
	logout(authGroup)

}
