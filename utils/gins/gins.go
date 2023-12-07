package gins

import (
	"kubez_project/models"
	"kubez_project/utils/logs"

	"github.com/gin-gonic/gin"
)

// 绑定json至结构体的方法
func BoundJson(c *gin.Context, any interface{}) {
	if err := c.ShouldBindJSON(any); err != nil {
		ReturnErrorData(c, "400", "json绑定错误", err)
	}
}

// 发生错误返回并记录log的方法
// 此函数会终止contxt会话
func ReturnErrorData(c *gin.Context, status string, msg string, err error) {
	logs.Error(map[string]interface{}{"Error": err.Error()}, msg)
	returnData := models.ReturnData{
		Status:  status,
		Message: msg,
		Data:    map[string]interface{}{"Error": err.Error()},
	}
	c.JSON(200, returnData)
	c.Abort()
}

// 默认状态为200,message为ok,data为空
func NewReturnData() models.ReturnData {
	returnData := models.ReturnData{
		Status:  "200",
		Message: "ok",
		Data:    make(map[string]interface{}),
	}
	return returnData
}
