package middlewares

import (
	"errors"
	"kubez_project/utils/gins"
	"kubez_project/utils/jwts"

	"github.com/gin-gonic/gin"
)

// 全局检测请求是否携带jwt-token
func JwtCheck(c *gin.Context) {
	requestUrl := c.FullPath()
	if requestUrl == "/api/auth/login" || requestUrl == "/api/auth/logout" {
		// login和logout接口不需要带token
		c.Next()
	} else {
		token := c.GetHeader("Authorization")
		if claims, err := jwts.ParseToken(token); err != nil {
			// 验证失败
			gins.ReturnErrorData(c, "401", "用户认证失败", errors.New("invaid user token"))
			c.Abort()
			return
		} else {
			// 验证成功,将claims传给*gin.Context,以供后续使用
			c.Set("claims", claims)
		}
	}
}
