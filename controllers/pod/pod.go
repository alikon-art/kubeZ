package pod

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

func Add(c *gin.Context) {
	var podItem corev1.Pod
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &podItem)
	if err != nil {
		return
	}

	_, err = clientset.CoreV1().Pods(podItem.Namespace).Create(context.TODO(), &podItem, metav1.CreateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "创建pod失败", err)
		return
	}

	gins.ReturnData(c, "200", "创建pod成功", nil)
}

func Delete(c *gin.Context) {
	var podItem corev1.Pod
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &podItem)
	if err != nil {
		return
	}

	err = clientset.CoreV1().Pods(requestData.Namespace).Delete(context.TODO(), requestData.Name, metav1.DeleteOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "删除pod失败", err)
		return
	}

	gins.ReturnData(c, "200", "删除pod成功", nil)
}

func Update(c *gin.Context) {
	var podItem corev1.Pod
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &podItem)
	if err != nil {
		return
	}

	_, err = clientset.CoreV1().Pods(podItem.Namespace).Update(context.TODO(), &podItem, metav1.UpdateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "更新pod失败", err)
		return
	}

	gins.ReturnData(c, "200", "更新pod成功", nil)

}

func List(c *gin.Context) {
	var podItem corev1.Pod
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &podItem)
	if err != nil {
		return
	}

	listOptions := metav1.ListOptions{
		LabelSelector: goclient.Convert2LabelSelector(requestData.Labels),
	}

	podList, err := clientset.CoreV1().Pods(requestData.Namespace).List(context.TODO(), listOptions)
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取pod列表失败", err)
		return
	}

	var returnDataList models.BasicReturnList
	for _, pod := range podList.Items {
		returnData := models.BasicReturn{
			Name:       pod.Name,
			Namespace:  pod.Namespace,
			Labels:     pod.Labels,
			CreateTime: pod.CreationTimestamp.Format(config.TimestampFormat),
		}
		returnDataList.AddBasicReturn(returnData)
	}

	gins.ReturnData(c, "200", "获取pod列表成功", returnDataList)
}

func Get(c *gin.Context) {
	var podItem corev1.Pod
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &podItem)
	if err != nil {
		return
	}

	pod, err := clientset.CoreV1().Pods(requestData.Namespace).Get(context.TODO(), requestData.Name, metav1.GetOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取pod详情失败", err)
		return
	}

	gins.ReturnData(c, "200", "获取pod详情成功", pod)
}
