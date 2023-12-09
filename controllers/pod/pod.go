package pod

import (
	"context"
	"kubez_project/utils/gins"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var podItem corev1.Pod

func Add(c *gin.Context) {

	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &podItem)
	if err != nil {
		return
	}

	_, err = clientset.CoreV1().Pods("default").Create(context.TODO(), &podItem, metav1.CreateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "创建pod失败", err)
		return
	}

	gins.ReturnData(c, "200", "创建pod成功", nil)
}

func Delete(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func List(c *gin.Context) {

}

func Get(c *gin.Context) {

}
