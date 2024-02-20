package ingress

import (
	"context"
	"fmt"
	"kubez_project/config"
	"kubez_project/models"
	"kubez_project/utils/gins"
	goclient "kubez_project/utils/go_client"

	"github.com/gin-gonic/gin"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Add 创建一个新的Ingress
func Add(c *gin.Context) {
	var ingressItem networkingv1.Ingress
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &ingressItem)
	if err != nil {
		return
	}

	fmt.Println(ingressItem)
	_, err = clientset.NetworkingV1().Ingresses(ingressItem.Namespace).Create(context.TODO(), &ingressItem, metav1.CreateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "创建Ingress失败", err)
		return
	}

	gins.ReturnData(c, "200", "创建Ingress成功", nil)
}

// Delete 删除一个Ingress
func Delete(c *gin.Context) {
	var ingressItem networkingv1.Ingress
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &ingressItem)
	if err != nil {
		return
	}

	err = clientset.NetworkingV1().Ingresses(requestData.Namespace).Delete(context.TODO(), requestData.Name, metav1.DeleteOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "删除Ingress失败", err)
		return
	}

	gins.ReturnData(c, "200", "删除Ingress成功", nil)
}

// Update 更新一个Ingress
func Update(c *gin.Context) {
	var ingressItem networkingv1.Ingress
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &ingressItem)
	if err != nil {
		return
	}

	_, err = clientset.NetworkingV1().Ingresses(ingressItem.Namespace).Update(context.TODO(), &ingressItem, metav1.UpdateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "更新Ingress失败", err)
		return
	}

	gins.ReturnData(c, "200", "更新Ingress成功", nil)
}

// List 列出所有Ingress
func List(c *gin.Context) {
	var ingressItem networkingv1.Ingress
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &ingressItem)
	if err != nil {
		return
	}

	listOptions := metav1.ListOptions{
		LabelSelector: goclient.Convert2LabelSelector(requestData.Labels),
	}

	ingressList, err := clientset.NetworkingV1().Ingresses(requestData.Namespace).List(context.TODO(), listOptions)
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取Ingress列表失败", err)
		return
	}

	var returnDataList models.BasicReturnList
	for _, ingress := range ingressList.Items {
		returnData := models.BasicReturn{
			Name:       ingress.Name,
			Namespace:  ingress.Namespace,
			Labels:     ingress.Labels,
			CreateTime: ingress.CreationTimestamp.Format(config.TimestampFormat),
			Item:       ingress,
		}
		returnDataList.AddBasicReturn(returnData)
	}

	gins.ReturnData(c, "200", "获取Ingress列表成功", returnDataList)
}

// Get 获取一个Ingress的详情
func Get(c *gin.Context) {
	var ingressItem networkingv1.Ingress
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &ingressItem)
	if err != nil {
		return
	}

	ingress, err := clientset.NetworkingV1().Ingresses(requestData.Namespace).Get(context.TODO(), requestData.Name, metav1.GetOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取Ingress详情失败", err)
		return
	}

	gins.ReturnData(c, "200", "获取Ingress详情成功", ingress)
}
