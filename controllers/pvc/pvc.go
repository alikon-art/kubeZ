package pvc

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

// Add 创建一个新的PVC
func Add(c *gin.Context) {
	var pvcItem corev1.PersistentVolumeClaim
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &pvcItem)
	if err != nil {
		return
	}

	_, err = clientset.CoreV1().PersistentVolumeClaims(pvcItem.Namespace).Create(context.TODO(), &pvcItem, metav1.CreateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "创建PVC失败", err)
		return
	}

	gins.ReturnData(c, "200", "创建PVC成功", nil)
}

// Delete 删除一个PVC
func Delete(c *gin.Context) {
	var pvcItem corev1.PersistentVolumeClaim
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &pvcItem)
	if err != nil {
		return
	}

	err = clientset.CoreV1().PersistentVolumeClaims(requestData.Namespace).Delete(context.TODO(), requestData.Name, metav1.DeleteOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "删除PVC失败", err)
		return
	}

	gins.ReturnData(c, "200", "删除PVC成功", nil)
}

// Update 更新一个PVC
func Update(c *gin.Context) {
	var pvcItem corev1.PersistentVolumeClaim
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &pvcItem)
	if err != nil {
		return
	}

	_, err = clientset.CoreV1().PersistentVolumeClaims(pvcItem.Namespace).Update(context.TODO(), &pvcItem, metav1.UpdateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "更新PVC失败", err)
		return
	}

	gins.ReturnData(c, "200", "更新PVC成功", nil)
}

// List 列出所有PVC
func List(c *gin.Context) {
	var pvcItem corev1.PersistentVolumeClaim
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &pvcItem)
	if err != nil {
		return
	}

	listOptions := metav1.ListOptions{
		LabelSelector: goclient.Convert2LabelSelector(requestData.Labels),
	}

	pvcList, err := clientset.CoreV1().PersistentVolumeClaims(requestData.Namespace).List(context.TODO(), listOptions)
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取PVC列表失败", err)
		return
	}

	var returnDataList models.BasicReturnList
	for _, pvc := range pvcList.Items {
		returnData := models.BasicReturn{
			Name:       pvc.Name,
			Namespace:  pvc.Namespace,
			Labels:     pvc.Labels,
			CreateTime: pvc.CreationTimestamp.Format(config.TimestampFormat),
		}
		returnDataList.AddBasicReturn(returnData)
	}

	gins.ReturnData(c, "200", "获取PVC列表成功", returnDataList)
}

// Get 获取一个PVC的详情
func Get(c *gin.Context) {
	var pvcItem corev1.PersistentVolumeClaim
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &pvcItem)
	if err != nil {
		return
	}

	pvc, err := clientset.CoreV1().PersistentVolumeClaims(requestData.Namespace).Get(context.TODO(), requestData.Name, metav1.GetOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取PVC详情失败", err)
		return
	}

	gins.ReturnData(c, "200", "获取PVC详情成功", pvc)
}
