package daemonset

import (
	"context"
	"fmt"
	"kubez_project/config"
	"kubez_project/models"
	"kubez_project/utils/gins"
	goclient "kubez_project/utils/go_client"

	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Add 创建一个新的DaemonSet
func Add(c *gin.Context) {
	var daemonsetItem appsv1.DaemonSet
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &daemonsetItem)
	if err != nil {
		return
	}

	fmt.Println(daemonsetItem)
	_, err = clientset.AppsV1().DaemonSets(daemonsetItem.Namespace).Create(context.TODO(), &daemonsetItem, metav1.CreateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "创建daemonset失败", err)
		return
	}

	gins.ReturnData(c, "200", "创建daemonset成功", nil)
}

// Delete 删除一个DaemonSet
func Delete(c *gin.Context) {
	var daemonsetItem appsv1.DaemonSet
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &daemonsetItem)
	if err != nil {
		return
	}

	err = clientset.AppsV1().DaemonSets(requestData.Namespace).Delete(context.TODO(), requestData.Name, metav1.DeleteOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "删除daemonset失败", err)
		return
	}

	gins.ReturnData(c, "200", "删除daemonset成功", nil)
}

// Update 更新一个DaemonSet
func Update(c *gin.Context) {
	var daemonsetItem appsv1.DaemonSet
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &daemonsetItem)
	if err != nil {
		return
	}

	_, err = clientset.AppsV1().DaemonSets(daemonsetItem.Namespace).Update(context.TODO(), &daemonsetItem, metav1.UpdateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "更新daemonset失败", err)
		return
	}

	gins.ReturnData(c, "200", "更新daemonset成功", nil)
}

// List 获取所有DaemonSet的列表
func List(c *gin.Context) {
	var daemonsetItem appsv1.DaemonSet
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &daemonsetItem)
	if err != nil {
		return
	}

	listOptions := metav1.ListOptions{
		LabelSelector: goclient.Convert2LabelSelector(requestData.Labels),
	}

	daemonsetList, err := clientset.AppsV1().DaemonSets(requestData.Namespace).List(context.TODO(), listOptions)
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取daemonset列表失败", err)
		return
	}

	var returnDataList models.BasicReturnList
	for _, daemonset := range daemonsetList.Items {
		returnData := models.BasicReturn{
			Name:       daemonset.Name,
			Namespace:  daemonset.Namespace,
			Labels:     daemonset.Labels,
			CreateTime: daemonset.CreationTimestamp.Format(config.TimestampFormat),
		}
		returnDataList.AddBasicReturn(returnData)
	}

	gins.ReturnData(c, "200", "获取daemonset列表成功", returnDataList)
}

// Get 获取一个DaemonSet的详情
func Get(c *gin.Context) {
	var daemonsetItem appsv1.DaemonSet
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &daemonsetItem)
	if err != nil {
		return
	}

	daemonset, err := clientset.AppsV1().DaemonSets(requestData.Namespace).Get(context.TODO(), requestData.Name, metav1.GetOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取daemonset详情失败", err)
		return
	}

	gins.ReturnData(c, "200", "获取daemonset详情成功", daemonset)
}
