package statefulset

import (
	"context"
	"kubez_project/config"
	"kubez_project/models"
	"kubez_project/utils/gins"
	goclient "kubez_project/utils/go_client"

	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Add 创建一个新的StatefulSet
func Add(c *gin.Context) {
	var statefulsetItem appsv1.StatefulSet
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &statefulsetItem)
	if err != nil {
		return
	}

	_, err = clientset.AppsV1().StatefulSets(statefulsetItem.Namespace).Create(context.TODO(), &statefulsetItem, metav1.CreateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "创建StatefulSet失败", err)
		return
	}

	gins.ReturnData(c, "200", "创建StatefulSet成功", nil)
}

// Delete 删除一个StatefulSet
func Delete(c *gin.Context) {
	var statefulsetItem appsv1.StatefulSet
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &statefulsetItem)
	if err != nil {
		return
	}

	err = clientset.AppsV1().StatefulSets(requestData.Namespace).Delete(context.TODO(), requestData.Name, metav1.DeleteOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "删除StatefulSet失败", err)
		return
	}

	gins.ReturnData(c, "200", "删除StatefulSet成功", nil)
}

// Update 更新一个StatefulSet
func Update(c *gin.Context) {
	var statefulsetItem appsv1.StatefulSet
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &statefulsetItem)
	if err != nil {
		return
	}

	_, err = clientset.AppsV1().StatefulSets(statefulsetItem.Namespace).Update(context.TODO(), &statefulsetItem, metav1.UpdateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "更新StatefulSet失败", err)
		return
	}

	gins.ReturnData(c, "200", "更新StatefulSet成功", nil)
}

// List 列出所有StatefulSet
func List(c *gin.Context) {
	var statefulsetItem appsv1.StatefulSet
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &statefulsetItem)
	if err != nil {
		return
	}

	listOptions := metav1.ListOptions{
		LabelSelector: goclient.Convert2LabelSelector(requestData.Labels),
	}

	statefulSetList, err := clientset.AppsV1().StatefulSets(requestData.Namespace).List(context.TODO(), listOptions)
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取StatefulSet列表失败", err)
		return
	}

	var returnDataList models.BasicReturnList
	for _, statefulSet := range statefulSetList.Items {
		returnData := models.BasicReturn{
			Name:       statefulSet.Name,
			Namespace:  statefulSet.Namespace,
			Labels:     statefulSet.Labels,
			CreateTime: statefulSet.CreationTimestamp.Format(config.TimestampFormat),
			Item:       statefulSet,
		}
		returnDataList.AddBasicReturn(returnData)
	}

	gins.ReturnData(c, "200", "获取StatefulSet列表成功", returnDataList)
}

// Get 获取一个StatefulSet的详情
func Get(c *gin.Context) {
	var statefulsetItem appsv1.StatefulSet
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &statefulsetItem)
	if err != nil {
		return
	}

	statefulSet, err := clientset.AppsV1().StatefulSets(requestData.Namespace).Get(context.TODO(), requestData.Name, metav1.GetOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取StatefulSet详情失败", err)
		return
	}

	gins.ReturnData(c, "200", "获取StatefulSet详情成功", statefulSet)
}
