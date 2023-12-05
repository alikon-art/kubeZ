package auth

import (
	"kubez_project/config"
	"kubez_project/models"
	"kubez_project/utils/jwts"
	"kubez_project/utils/logs"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	returnData := models.NewRetunData()
	userInfo := UserInfo{}

	if err := c.ShouldBindJSON(&userInfo); err != nil {
		// 绑定json到结构体失败,一般是传入的数据格式有问题
		returnData.Status = 401
		returnData.Message = err.Error()
		c.JSON(200, returnData)
		return
	}
	logs.Debug(map[string]interface{}{"用户名": userInfo.Username, "密码": userInfo.Password}, "取得登录信息")

	if userInfo.Username == config.Username && userInfo.Password == config.Password {
		// 验证密码正确
		token, err := jwts.GenToken(userInfo.Username)
		if err != nil {
			// 生成jwt-token失败
			returnData.Status = 401
			returnData.Message = err.Error()
			c.JSON(200, returnData)
			return
		} else {
			// 成功生成jwt-token,构造返回数据
			returnData.Data["token"] = token
			c.JSON(200, returnData)
		}
	} else {
		// 验证密码失败
		returnData.Status = 401
		returnData.Message = "用户名或密码错误"
		c.JSON(200, returnData)
		return
	}

}

func Logout(c *gin.Context) {
	returnData := models.NewRetunData()
	returnData.Message = "用户已退出"
	c.JSON(200, returnData)
	return
}
