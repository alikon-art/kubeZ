package controllers

import (
	"kubez_project/models"
	"kubez_project/utils/logs"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 创建kube控制器
func InitKubeClient(kubeconfig string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeconfig))
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, err
}

// 新建一个error返回,并记录log
// 传入:状态,错误信息,err
func NewReturnErrorData(status string, msg string, err error) models.ReturnData {
	logs.Error(map[string]interface{}{"Error": err.Error()}, msg)
	returnData := models.ReturnData{
		Status:  status,
		Message: msg,
		Data:    map[string]interface{}{"Error": err.Error()},
	}
	return returnData
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
