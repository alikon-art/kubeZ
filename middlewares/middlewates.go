package middlewares

<<<<<<< HEAD
// import (
// 	"kubez_project/controllers"
// 	"kubez_project/utils/jwts"
=======
import (
	"kubez_project/controllers"
	"kubez_project/utils/jwts"
>>>>>>> 6512db693d369297878d5d9833b4c1d8f000e958

// 	"github.com/gin-gonic/gin"
// )

// 全局检测请求是否携带jwt-token
<<<<<<< HEAD
// func JwtCheck(c *gin.Context) {
// 	requestUrl := c.FullPath()
// 	if requestUrl == "/api/auth/login" || requestUrl == "/api/auth/logout" {
// 		// login和logout接口不需要带token
// 		c.Next()
// 	}
// 	token := c.GetHeader("Authorization")
// 	if claims, err := jwts.ParseToken(token); err != nil {
// 		c.JSON(200, controllers.NewReturnErrorData("401", "token验证失败", err))
// 		c.Abort()
// 		return
// 	} else {
// 		// 验证成功,将claims传给*gin.Context,以供后续使用
// 		c.Set("claims", claims)
=======
func JwtCheck(c *gin.Context) {
	requestUrl := c.FullPath()
	if requestUrl == "/api/auth/login" || requestUrl == "/api/auth/logout" {
		// login和logout接口不需要带token
		c.Next()
	}
	token := c.GetHeader("Authorization")
	if claims, err := jwts.ParseToken(token); err != nil {
		c.JSON(200, controllers.NewReturnErrorData("401", "token验证失败", err))
		c.Abort()
		return
	} else {
		// 验证成功,将claims传给*gin.Context,以供后续使用
		c.Set("claims", claims)
>>>>>>> 6512db693d369297878d5d9833b4c1d8f000e958

// 	}
// }
