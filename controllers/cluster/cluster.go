package cluster

import (
	"context"
	"fmt"
	"kubez_project/config"
	"kubez_project/controllers"
	"kubez_project/utils/logs"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Add(c *gin.Context) {
	clientset, err := controllers.InitKubeClient(config.Kubeconfig)
	if err != nil {
		logs.Error(map[string]interface{}{"Error : ": err.Error()}, "初始化contest失败!")
		c.Abort()
	}

	secret := corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test",
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
