package auth

import (
	"errors"
	"kubez_project/config"
	"kubez_project/models"
	"kubez_project/utils/gins"
	"kubez_project/utils/jwts"
	"kubez_project/utils/logs"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	userInfo := models.UserInfo{}
	gins.BoundJson(c, &userInfo)
	logs.Debug(map[string]interface{}{"用户名": userInfo.Username, "密码": userInfo.Password}, "取得登录信息")

	if userInfo.Username != config.Username || userInfo.Password != config.Password {
		gins.ReturnErrorData(c, "401", "用户名或密码错误", errors.New("用户名或密码错误"))
		c.Abort()
		return
	}

	// 验证密码正确
	token, err := jwts.GenToken(userInfo.Username)
	if err != nil {
		gins.ReturnErrorData(c, "500", "token生成失败", err)
	} else {
		// 成功生成jwt-token,构造返回数据
		data := map[string]interface{}{
			"token": token,
		}
		gins.ReturnData(c, "200", "ok", data)
	}

}

func Logout(c *gin.Context) {
	gins.ReturnData(c, "200", "用户已退出", nil)

}
