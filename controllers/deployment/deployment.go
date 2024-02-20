package deployment

import (
	"context"
	"fmt"
	"kubez_project/config"
	"kubez_project/models"
	"kubez_project/utils/gins"
	goclient "kubez_project/utils/go_client"

	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"

	// corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Add(c *gin.Context) {
	var deploymentItem appsv1.Deployment
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &deploymentItem)
	if err != nil {
		return
	}

	fmt.Println(deploymentItem)
	_, err = clientset.AppsV1().Deployments(deploymentItem.Namespace).Create(context.TODO(), &deploymentItem, metav1.CreateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "创建deployment失败", err)
		return
	}

	gins.ReturnData(c, "200", "创建deployment成功", nil)
}

func Delete(c *gin.Context) {
	var deploymentItem appsv1.Deployment
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &deploymentItem)
	if err != nil {
		return
	}

	err = clientset.AppsV1().Deployments(requestData.Namespace).Delete(context.TODO(), requestData.Name, metav1.DeleteOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "删除deployment失败", err)
		return
	}

	gins.ReturnData(c, "200", "删除deployment成功", nil)
}

func Update(c *gin.Context) {
	var deploymentItem appsv1.Deployment
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &deploymentItem)
	if err != nil {
		return
	}

	_, err = clientset.AppsV1().Deployments(deploymentItem.Namespace).Update(context.TODO(), &deploymentItem, metav1.UpdateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "更新deployment失败", err)
		return
	}

	gins.ReturnData(c, "200", "更新deployment成功", nil)

}

func List(c *gin.Context) {
	var deploymentItem appsv1.Deployment
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &deploymentItem)
	if err != nil {
		return
	}

	listOptions := metav1.ListOptions{
		LabelSelector: goclient.Convert2LabelSelector(requestData.Labels),
	}

	deploymentList, err := clientset.AppsV1().Deployments(requestData.Namespace).List(context.TODO(), listOptions)
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取deployment列表失败", err)
		return
	}

	var returnDataList models.BasicReturnList
	for _, deployment := range deploymentList.Items {
		returnData := models.BasicReturn{
			Name:       deployment.Name,
			Namespace:  deployment.Namespace,
			Labels:     deployment.Labels,
			CreateTime: deployment.CreationTimestamp.Format(config.TimestampFormat),
			Item:       deployment,
		}
		returnDataList.AddBasicReturn(returnData)
	}

	gins.ReturnData(c, "200", "获取deployment列表成功", returnDataList)
}

func Get(c *gin.Context) {
	var deploymentItem appsv1.Deployment
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &deploymentItem)
	if err != nil {
		return
	}

	deployment, err := clientset.AppsV1().Deployments(requestData.Namespace).Get(context.TODO(), requestData.Name, metav1.GetOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取deployment详情失败", err)
		return
	}

	gins.ReturnData(c, "200", "获取deployment详情成功", deployment)
}
