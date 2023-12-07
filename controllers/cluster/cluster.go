package cluster

import (
	"context"
	"fmt"
	"kubez_project/controllers"
	"kubez_project/models"
	"kubez_project/utils/gins"

	// 	appsv1 "k8s.io/api/apps/v1"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var clusterDataWConfig models.ClusterDataWConfig

func Add(c *gin.Context) {
	gins.BoundJson(c, &clusterDataWConfig)
	clientset, err := controllers.InitKubeClient(clusterDataWConfig.ClusterConfig)
	if err != nil {
		gins.ReturnErrorData(c, "400", "创建clientset失败", err)
	}

	secret := corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterDataWConfig.ClusterID,
			Namespace: "default",
			Labels: map[string]string{
				"app": "kubez",
			},
		},
		// StringData自动base64加密
		StringData: map[string]string{
			"clusterid": "123",
			"anon":      "null",
		},
	}

	createdSecret, err := clientset.CoreV1().Secrets("default").Create(context.TODO(), &secret, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(createdSecret)
		c.JSON(200, "ok")
	}

}

func Delete(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func List(c *gin.Context) {

}
