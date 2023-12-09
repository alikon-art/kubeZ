package cluster

import (
	"context"
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

// add方法会验证集群连通性,并获取集群版本
func Add(c *gin.Context) {
	// 绑定json到结构体
	if gins.BoundJson(c, &clusterDataWConfig) != nil {
		c.Abort()
		return
	}

	// 从结构体创建clientset,测试连通性
	clusterset, err := gins.InitKubeClient(c, clusterDataWConfig.ClusterConfig)
	if err != nil {
		c.Abort()
		return
	}

	// 获取集群版本
	clusterVersion, err := gins.GetClusterVersion(c, clusterset)
	if err != nil {
		return
	}

	// 将版本放到data结构体中
	clusterDataWConfig.Version = clusterVersion

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
	createdSecret, err := controllers.InclusterClientSet.CoreV1().Secrets(config.KubeZNamespace).Create(context.TODO(), &secret, metav1.CreateOptions{})
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
	clusterid := c.Query("clusterid")

	err := controllers.InclusterClientSet.CoreV1().Secrets(config.KubeZNamespace).Delete(context.TODO(), clusterid, metav1.DeleteOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "删除失败", err)
		return
	}

	gins.ReturnData(c, "200", "ok", nil)

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
	createdSecret, err := controllers.InclusterClientSet.CoreV1().Secrets(config.KubeZNamespace).Update(context.TODO(), &secret, metav1.UpdateOptions{})
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

	// 查询label
	listOptions := metav1.ListOptions{
		LabelSelector: config.KubeZLabelsKey + "=" + config.KubeZLabelsValue,
	}

	querySecretList, err := controllers.InclusterClientSet.CoreV1().Secrets(config.KubeZNamespace).List(context.TODO(), listOptions)
	if err != nil {
		gins.ReturnErrorData(c, "500", "查询失败", err)
		c.Abort()
		return
	}

	var clusterList models.ClusterList
	var clusterdata models.ClusterListData
	for _, v := range querySecretList.Items {
		clusterdata = models.ClusterListData{
			ClusterID:   v.Name,
			ClusterName: string(v.Data["clustername"]),
			Version:     string(v.Data["version"]),
			// Annotations: v.Data["annotations"],
			CreateTime: v.CreationTimestamp.Time.Format(config.TimestampFormat),
		}
		clusterList.AddCluster(clusterdata)

	}

	gins.ReturnData(c, "200", "ok", clusterList)
}

func Get(c *gin.Context) {
	var requestData models.RequestDataFrame
	err := gins.BoundJson(c, &requestData)
	if err != nil {
		return
	}

	cluster, err := controllers.InclusterClientSet.CoreV1().Secrets(config.KubeZNamespace).Get(context.TODO(), requestData.ClusterID, metav1.GetOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "查询失败", err)
		return
	}

	gins.ReturnData(c, "200", "ok", cluster)
}
