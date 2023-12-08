package cluster

import (
	"context"
	"fmt"
	"kubez_project/config"
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
	// 绑定json到结构体
	if gins.BoundJson(c, &clusterDataWConfig) != nil {
		c.Abort()
		return
	}
	// 从结构体创建clientset,测试连通性
	_, err := gins.InitKubeClient(c, clusterDataWConfig.ClusterConfig)
	if err != nil {
		c.Abort()
		return
	}
	// 将结构体的数据转为secret数据
	data, err := gins.Struct2MapStr(c, clusterDataWConfig)
	if err != nil {
		c.Abort()
		return
	}
	// secret对象
	secret := corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterDataWConfig.ClusterID,
			Namespace: config.KubeZNamespace,
			Labels:    config.KubeZLabels,
		},
		// 将ClusterDataWConfig结构体的所有数据作为secret的data部分
		// StringData自动base64加密
		StringData: data,
	}
	// 创建secret
	createdSecret, err := controllers.InclusterKubeSet.CoreV1().Secrets(config.KubeZNamespace).Create(context.TODO(), &secret, metav1.CreateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "创建secret失败", err)
		c.Abort()
		return
	}
	// 构建返回data
	secretinfo, err := gins.Struct2MapInterface(c, createdSecret)
	if err != nil {
		c.Abort()
		return
	}
	gins.ReturnData(c, "200", "添加集群成功", secretinfo)

}

func Delete(c *gin.Context) {

}

func Update(c *gin.Context) {
	if gins.BoundJson(c, &clusterDataWConfig) != nil {
		c.Abort()
		return
	}
	// 从结构体创建clientset,测试连通性
	_, err := gins.InitKubeClient(c, clusterDataWConfig.ClusterConfig)
	if err != nil {
		c.Abort()
		return
	}
	// 将结构体的数据转为secret数据
	data, err := gins.Struct2MapStr(c, clusterDataWConfig)
	if err != nil {
		c.Abort()
		return
	}
	// secret对象
	secret := corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterDataWConfig.ClusterID,
			Namespace: config.KubeZNamespace,
			Labels:    config.KubeZLabels,
		},
		// 将ClusterDataWConfig结构体的所有数据作为secret的data部分
		// StringData自动base64加密
		StringData: data,
	}
	// 创建secret
	createdSecret, err := controllers.InclusterKubeSet.CoreV1().Secrets(config.KubeZNamespace).Update(context.TODO(), &secret, metav1.UpdateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "更新secret失败", err)
		c.Abort()
		return
	}
	// 构建返回data
	secretinfo, err := gins.Struct2MapInterface(c, createdSecret)
	if err != nil {
		c.Abort()
		return
	}
	gins.ReturnData(c, "200", "更新集群成功", secretinfo)
}

func List(c *gin.Context) {
	gins.BoundJson(c, &clusterDataWConfig)
	clientset, err := gins.InitKubeClient(c, clusterDataWConfig.ClusterConfig)
	if err != nil {
		c.Abort()
	}
	SecretList, err := clientset.CoreV1().Secrets("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "创建secret失败", err)
		c.Abort()
	}
	fmt.Println(SecretList)
	/// todo
	gins.ReturnData(c, "200", "ok", nil)
}

func Get(c *gin.Context) {
	// clusterid := c.Query("clusterid")

}
