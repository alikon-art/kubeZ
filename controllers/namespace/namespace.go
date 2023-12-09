package namespace

import (
	"context"
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

func Add(c *gin.Context) {
	var requestData models.RequestData
	err := gins.BoundJson(c, &requestData)
	if err != nil {
		return
	}

	// 从controllers去对应集群的clientset
	clientset, err := gins.GetOutOfClusterClientSet(c, requestData.ClusterID)
	if err != nil {
		return
	}

	namespace := corev1.Namespace{}
	namespace.SetName(requestData.Name)
	_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), &namespace, metav1.CreateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "创建namespace失败", err)
		return
	}

	gins.ReturnData(c, "200", "创建namespace成功", nil)

}

func Delete(c *gin.Context) {
	var requestData models.RequestData
	err := gins.BoundJson(c, &requestData)
	if err != nil {
		return
	}

	// 从controllers去对应集群的clientset
	clientset, err := gins.GetOutOfClusterClientSet(c, requestData.ClusterID)
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
	// ns能更新啥?
}

func List(c *gin.Context) {
	var requestData models.RequestData
	err := gins.BoundJson(c, &requestData)
	if err != nil {
		return
	}

	// 从controllers去对应集群的clientset
	clientset, err := gins.GetOutOfClusterClientSet(c, requestData.ClusterID)
	if err != nil {
		return
	}

	namespaceList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "查询namespace列表失败", err)
		return
	}

	returnList := models.ReturnList{}
	for _, v := range namespaceList.Items {
		returnListData := models.ReturnListData{
			Name:       v.Name,
			CreateTime: v.CreationTimestamp.Format(config.TimestampFormat),
		}
		returnList.AddReturnListData(returnListData)
	}

	gins.ReturnData(c, "200", "ok", returnList)

}

func Get(c *gin.Context) {
	var requestData models.RequestData
	err := gins.BoundJson(c, &requestData)
	if err != nil {
		return
	}

	// 从controllers去对应集群的clientset
	clientset, err := gins.GetOutOfClusterClientSet(c, requestData.ClusterID)
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
