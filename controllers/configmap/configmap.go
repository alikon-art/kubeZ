package configmap

import (
	"context"
	"kubez_project/config"
	"kubez_project/models"
	"kubez_project/utils/gins"
	goclient "kubez_project/utils/go_client"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Add 创建一个新的ConfigMap
func Add(c *gin.Context) {
	var configMapItem corev1.ConfigMap
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &configMapItem)
	if err != nil {
		return
	}

	_, err = clientset.CoreV1().ConfigMaps(configMapItem.Namespace).Create(context.TODO(), &configMapItem, metav1.CreateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "创建ConfigMap失败", err)
		return
	}

	gins.ReturnData(c, "200", "创建ConfigMap成功", nil)
}

// Delete 删除一个ConfigMap
func Delete(c *gin.Context) {
	var configMapItem corev1.ConfigMap
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &configMapItem)
	if err != nil {
		return
	}

	err = clientset.CoreV1().ConfigMaps(requestData.Namespace).Delete(context.TODO(), requestData.Name, metav1.DeleteOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "删除ConfigMap失败", err)
		return
	}

	gins.ReturnData(c, "200", "删除ConfigMap成功", nil)
}

// Update 更新一个ConfigMap
func Update(c *gin.Context) {
	var configMapItem corev1.ConfigMap
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &configMapItem)
	if err != nil {
		return
	}

	_, err = clientset.CoreV1().ConfigMaps(configMapItem.Namespace).Update(context.TODO(), &configMapItem, metav1.UpdateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "更新ConfigMap失败", err)
		return
	}

	gins.ReturnData(c, "200", "更新ConfigMap成功", nil)
}

// List 获取ConfigMap列表
func List(c *gin.Context) {
	var configMapItem corev1.ConfigMap
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &configMapItem)
	if err != nil {
		return
	}

	listOptions := metav1.ListOptions{
		LabelSelector: goclient.Convert2LabelSelector(requestData.Labels),
	}

	configMapList, err := clientset.CoreV1().ConfigMaps(requestData.Namespace).List(context.TODO(), listOptions)
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取ConfigMap列表失败", err)
		return
	}

	var returnDataList models.BasicReturnList
	for _, configMap := range configMapList.Items {
		returnData := models.BasicReturn{
			Name:       configMap.Name,
			Namespace:  configMap.Namespace,
			Labels:     configMap.Labels,
			CreateTime: configMap.CreationTimestamp.Format(config.TimestampFormat),
			Item:       configMap,
		}
		returnDataList.AddBasicReturn(returnData)
	}

	gins.ReturnData(c, "200", "获取ConfigMap列表成功", returnDataList)
}

// Get 获取一个ConfigMap的详情
func Get(c *gin.Context) {
	var configMapItem corev1.ConfigMap
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &configMapItem)
	if err != nil {
		return
	}

	configMap, err := clientset.CoreV1().ConfigMaps(requestData.Namespace).Get(context.TODO(), requestData.Name, metav1.GetOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取ConfigMap详情失败", err)
		return
	}

	gins.ReturnData(c, "200", "获取ConfigMap详情成功", configMap)
}
