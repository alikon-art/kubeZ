package gins

import (
	// "fmt"
	"kubez_project/models"
	"kubez_project/utils/base64s"
	goclient "kubez_project/utils/go_client"
	"kubez_project/utils/jsons"
	"kubez_project/utils/logs"

	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
)

// 绑定json至结构体的方法
// 传指针!传指针!传指针!!
func BoundJson(c *gin.Context, any interface{}) (err error) {
	if err := c.ShouldBindJSON(any); err != nil {
		ReturnErrorData(c, "400", "json绑定错误", err)
		c.Abort()
		return err
	}
	return err
}

func ReturnData(c *gin.Context, status string, msg string, data interface{}) {
	returnData := models.ReturnData{
		Status:  status,
		Message: msg,
		Data:    data,
	}
	c.JSON(200, returnData)
}

// 返回错误信息的方法
// 此函数会同时记录log
func ReturnErrorData(c *gin.Context, status string, msg string, err error) {
	var data map[string]interface{}
	if err != nil {
		data = map[string]interface{}{
			"error": err.Error(),
		}
	}
	ReturnData(c, status, msg, data)
	logs.Error(map[string]interface{}{"Error": err.Error()}, msg)
}

// 从kubeconfig初始化clientset的方法
// 若发生错误,此函数会返回错误信息并记录log
// 传入的kubeconfig应该是base64加密的
func InitKubeClient(c *gin.Context, kubeconfig string) (clientset *kubernetes.Clientset, err error) {

	decodeConfig, err := base64s.DeCode(kubeconfig)
	if err != nil {
		ReturnErrorData(c, "400", "解密kubeconfig失败", err)
		c.Abort()
		return nil, err
	}
	clientset, err = goclient.InitKubeClient(decodeConfig)
	if err != nil {
		ReturnErrorData(c, "400", "初始化clientset失败", err)
		c.Abort()
		return nil, err
	}
	return clientset, err
}

// 将cluster结构体转为map[string]string的方法
// 若发生错误,此函数会返回错误信息并记录log
func Struct2MapStr(c *gin.Context, s interface{}) (m map[string]string, err error) {
	m, err = jsons.Struct2MapStr(s)
	if err != nil {
		ReturnErrorData(c, "500", "结构体转mapstr失败", err)
	}
	return m, err
}

// 将cluster结构体转为map[string]interface{}的方法
// 若发生错误,此函数会返回错误信息并记录log
func Struct2MapInterface(c *gin.Context, s interface{}) (m map[string]interface{}, err error) {
	m, err = jsons.Struct2MapInterface(s)
	if err != nil {
		ReturnErrorData(c, "500", "结构体转mapstr失败", err)
		c.Abort()
	}
	return m, err
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

// 获取集群版本的方法
func GetClusterVersion(c *gin.Context, clientset *kubernetes.Clientset) (version string, err error) {
	version, err = goclient.GetClusterVersion(clientset)
	if err != nil {
		ReturnErrorData(c, "500", "获取集群版本失败", err)
		c.Abort()
		return "", err
	}
	return version, err
}
