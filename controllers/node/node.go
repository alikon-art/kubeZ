package node

import (
	"context"
	"kubez_project/utils/gins"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var nodeItem corev1.Node

func Add(c *gin.Context) {
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &nodeItem)
	if err != nil {
		return
	}
	nodeItem.SetName(requestData.Name)
	nodeItem.SetLabels(requestData.Labels)
	_, err = clientset.CoreV1().Nodes().Create(context.TODO(), &nodeItem, metav1.CreateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "创建node失败", err)
		return
	}
	gins.ReturnData(c, "200", "创建node成功", nil)
}
func Delete(c *gin.Context) {
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &nodeItem)
	if err != nil {
		return
	}
	err = clientset.CoreV1().Nodes().Delete(context.TODO(), requestData.Name, metav1.DeleteOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "删除node失败", err)
		return
	}
	gins.ReturnData(c, "200", "删除node成功", nil)
}
func Update(c *gin.Context) {
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &nodeItem)
	if err != nil {
		return
	}
	nodeItem.SetName(requestData.Name)
	nodeItem.SetLabels(requestData.Labels)
	_, err = clientset.CoreV1().Nodes().Update(context.TODO(), &nodeItem, metav1.UpdateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "更新node失败", err)
		return
	}
	gins.ReturnData(c, "200", "更新node成功", nil)
}
func List(c *gin.Context) {
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &nodeItem)
	if err != nil {
		return
	}
	nodeList, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取node列表失败", err)
		return
	}
	gins.ReturnData(c, "200", "获取node列表成功", nodeList)
}
func Get(c *gin.Context) {
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &nodeItem)
	if err != nil {
		return
	}
	node, err := clientset.CoreV1().Nodes().Get(context.TODO(), requestData.Name, metav1.GetOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取node失败", err)
		return
	}
	gins.ReturnData(c, "200", "获取node成功", node)
}
