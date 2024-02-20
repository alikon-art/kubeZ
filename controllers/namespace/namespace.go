package namespace

import (
	"context"
	// "kubez_project/config"
	"kubez_project/config"
	"kubez_project/models"
	"kubez_project/utils/gins"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// "context"
// "fmt"
// "kubez_project/config"
// "kubez_project/models"
// "kubez_project/utils/gins"

// // 	appsv1 "k8s.io/api/apps/v1"
// "github.com/gin-gonic/gin"

// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

var namespaceItem corev1.Namespace

func Add(c *gin.Context) {

	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &namespaceItem)
	if err != nil {
		return
	}

	namespaceItem.SetName(requestData.Name)
	namespaceItem.SetLabels(requestData.Labels)

	_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), &namespaceItem, metav1.CreateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "创建namespace失败", err)
		return
	}

	gins.ReturnData(c, "200", "创建namespace成功", nil)

}

func Delete(c *gin.Context) {
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &namespaceItem)
	if err != nil {
		return
	}

	err = clientset.CoreV1().Namespaces().Delete(context.TODO(), requestData.Name, metav1.DeleteOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "删除namespace失败", err)
		return
	}

	gins.ReturnData(c, "200", "删除namespace成功", nil)

}

func Update(c *gin.Context) {
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &namespaceItem)
	if err != nil {
		return
	}

	namespaceItem.SetLabels(requestData.Labels)

	_, err = clientset.CoreV1().Namespaces().Update(context.TODO(), &namespaceItem, metav1.UpdateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "更新namespace失败", err)
		return
	}
	gins.ReturnData(c, "200", "更新namespace成功", nil)
}

func List(c *gin.Context) {

	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &namespaceItem)
	if err != nil {
		return
	}

	namespaceList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "查询namespace列表失败", err)
		return
	}

	returnList := models.BasicReturnList{}
	for _, v := range namespaceList.Items {
		returnListData := models.BasicReturn{
			Name:       v.Name,
			Labels:     v.Labels,
			CreateTime: v.CreationTimestamp.Format(config.TimestampFormat),
			Item:       v,
		}
		returnList.AddBasicReturn(returnListData)
	}

	gins.ReturnData(c, "200", "ok", returnList)

}

func Get(c *gin.Context) {
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &namespaceItem)
	if err != nil {
		return
	}

	namespace, err := clientset.CoreV1().Namespaces().Get(context.TODO(), requestData.Name, metav1.GetOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取namespace失败", err)
		return
	}

	gins.ReturnData(c, "200", "获取namespace成功", namespace)

}
