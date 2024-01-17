package service

import (
	"context"
	"fmt"
	"kubez_project/config"
	"kubez_project/models"
	"kubez_project/utils/gins"
	goclient "kubez_project/utils/go_client"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Add 创建一个新的服务
func Add(c *gin.Context) {
	var serviceItem corev1.Service
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &serviceItem)
	if err != nil {
		return
	}

	fmt.Println(serviceItem)
	_, err = clientset.CoreV1().Services(serviceItem.Namespace).Create(context.TODO(), &serviceItem, metav1.CreateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "创建service失败", err)
		return
	}

	gins.ReturnData(c, "200", "创建service成功", nil)
}

// Delete 删除一个服务
func Delete(c *gin.Context) {
	var serviceItem corev1.Service
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &serviceItem)
	if err != nil {
		return
	}

	err = clientset.CoreV1().Services(requestData.Namespace).Delete(context.TODO(), requestData.Name, metav1.DeleteOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "删除service失败", err)
		return
	}

	gins.ReturnData(c, "200", "删除service成功", nil)
}

// Update 更新一个服务
func Update(c *gin.Context) {
	var serviceItem corev1.Service
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &serviceItem)
	if err != nil {
		return
	}

	_, err = clientset.CoreV1().Services(serviceItem.Namespace).Update(context.TODO(), &serviceItem, metav1.UpdateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "更新service失败", err)
		return
	}

	gins.ReturnData(c, "200", "更新service成功", nil)
}

// List 列出所有服务
func List(c *gin.Context) {
	var serviceItem corev1.Service
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &serviceItem)
	if err != nil {
		return
	}

	listOptions := metav1.ListOptions{
		LabelSelector: goclient.Convert2LabelSelector(requestData.Labels),
	}

	serviceList, err := clientset.CoreV1().Services(requestData.Namespace).List(context.TODO(), listOptions)
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取service列表失败", err)
		return
	}

	var returnDataList models.BasicReturnList
	for _, service := range serviceList.Items {
		returnData := models.BasicReturn{
			Name:       service.Name,
			Namespace:  service.Namespace,
			Labels:     service.Labels,
			CreateTime: service.CreationTimestamp.Format(config.TimestampFormat),
		}
		returnDataList.AddBasicReturn(returnData)
	}

	gins.ReturnData(c, "200", "获取service列表成功", returnDataList)
}

// Get 获取一个服务的详细信息
func Get(c *gin.Context) {
	var serviceItem corev1.Service
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &serviceItem)
	if err != nil {
		return
	}

	service, err := clientset.CoreV1().Services(requestData.Namespace).Get(context.TODO(), requestData.Name, metav1.GetOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取service详情失败", err)
		return
	}

	gins.ReturnData(c, "200", "获取service详情成功", service)
}
