package secret

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

// Create 创建一个新的Secret
func Create(c *gin.Context) {
	var secretItem corev1.Secret
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &secretItem)
	if err != nil {
		return
	}

	_, err = clientset.CoreV1().Secrets(secretItem.Namespace).Create(context.TODO(), &secretItem, metav1.CreateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "创建Secret失败", err)
		return
	}

	gins.ReturnData(c, "200", "创建Secret成功", nil)
}

// Delete 删除一个Secret
func Delete(c *gin.Context) {
	var secretItem corev1.Secret
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &secretItem)
	if err != nil {
		return
	}

	err = clientset.CoreV1().Secrets(requestData.Namespace).Delete(context.TODO(), requestData.Name, metav1.DeleteOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "删除Secret失败", err)
		return
	}

	gins.ReturnData(c, "200", "删除Secret成功", nil)
}

// Update 更新一个Secret
func Update(c *gin.Context) {
	var secretItem corev1.Secret
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &secretItem)
	if err != nil {
		return
	}

	_, err = clientset.CoreV1().Secrets(secretItem.Namespace).Update(context.TODO(), &secretItem, metav1.UpdateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "更新Secret失败", err)
		return
	}

	gins.ReturnData(c, "200", "更新Secret成功", nil)
}

// List 列出所有Secret
func List(c *gin.Context) {
	var secretItem corev1.Secret
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &secretItem)
	if err != nil {
		return
	}

	listOptions := metav1.ListOptions{
		LabelSelector: goclient.Convert2LabelSelector(requestData.Labels),
	}

	secretList, err := clientset.CoreV1().Secrets(requestData.Namespace).List(context.TODO(), listOptions)
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取Secret列表失败", err)
		return
	}

	var returnDataList models.BasicReturnList
	for _, secret := range secretList.Items {
		returnData := models.BasicReturn{
			Name:       secret.Name,
			Namespace:  secret.Namespace,
			Labels:     secret.Labels,
			CreateTime: secret.CreationTimestamp.Format(config.TimestampFormat),
		}
		returnDataList.AddBasicReturn(returnData)
	}

	gins.ReturnData(c, "200", "获取Secret列表成功", returnDataList)
}

// Get 获取一个Secret的详情
func Get(c *gin.Context) {
	var secretItem corev1.Secret
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &secretItem)
	if err != nil {
		return
	}

	secret, err := clientset.CoreV1().Secrets(requestData.Namespace).Get(context.TODO(), requestData.Name, metav1.GetOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取Secret详情失败", err)
		return
	}

	gins.ReturnData(c, "200", "获取Secret详情成功", secret)
}
