package pv

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

// Add 创建一个新的PersistentVolume
func Add(c *gin.Context) {
	var pvItem corev1.PersistentVolume
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &pvItem)
	if err != nil {
		return
	}

	_, err = clientset.CoreV1().PersistentVolumes().Create(context.TODO(), &pvItem, metav1.CreateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "创建PersistentVolume失败", err)
		return
	}

	gins.ReturnData(c, "200", "创建PersistentVolume成功", nil)
}

// Delete 删除一个PersistentVolume
func Delete(c *gin.Context) {
	var pvItem corev1.PersistentVolume
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &pvItem)
	if err != nil {
		return
	}

	err = clientset.CoreV1().PersistentVolumes().Delete(context.TODO(), requestData.Name, metav1.DeleteOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "删除PersistentVolume失败", err)
		return
	}

	gins.ReturnData(c, "200", "删除PersistentVolume成功", nil)
}

// Update 更新一个PersistentVolume
func Update(c *gin.Context) {
	var pvItem corev1.PersistentVolume
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &pvItem)
	if err != nil {
		return
	}

	_, err = clientset.CoreV1().PersistentVolumes().Update(context.TODO(), &pvItem, metav1.UpdateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "更新PersistentVolume失败", err)
		return
	}

	gins.ReturnData(c, "200", "更新PersistentVolume成功", nil)
}

// List 列出所有PersistentVolume
func List(c *gin.Context) {
	var pvItem corev1.PersistentVolume
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &pvItem)
	if err != nil {
		return
	}

	listOptions := metav1.ListOptions{
		LabelSelector: goclient.Convert2LabelSelector(requestData.Labels),
	}

	pvList, err := clientset.CoreV1().PersistentVolumes().List(context.TODO(), listOptions)
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取PersistentVolume列表失败", err)
		return
	}

	var returnDataList models.BasicReturnList
	for _, pv := range pvList.Items {
		returnData := models.BasicReturn{
			Name:       pv.Name,
			Labels:     pv.Labels,
			CreateTime: pv.CreationTimestamp.Format(config.TimestampFormat),
		}
		returnDataList.AddBasicReturn(returnData)
	}

	gins.ReturnData(c, "200", "获取PersistentVolume列表成功", returnDataList)
}

// Get 获取一个PersistentVolume的详情
func Get(c *gin.Context) {
	var pvItem corev1.PersistentVolume
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &pvItem)
	if err != nil {
		return
	}

	pv, err := clientset.CoreV1().PersistentVolumes().Get(context.TODO(), requestData.Name, metav1.GetOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取PersistentVolume详情失败", err)
		return
	}

	gins.ReturnData(c, "200", "获取PersistentVolume详情成功", pv)
}
