package controllers

import (
	"context"
	"kubez_project/config"
	"kubez_project/utils/base64s"
	"kubez_project/utils/cryptos"
	goclient "kubez_project/utils/go_client"
	"kubez_project/utils/logs"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var InclusterClientSet *kubernetes.Clientset

var OutOfClusterClientSet map[string]*kubernetes.Clientset

// 初始化外部集群的clientset
func initOutOfClusterClientSet() {
	// 初始化map
	OutOfClusterClientSet = make(map[string]*kubernetes.Clientset)
	// 查询label
	listOptions := metav1.ListOptions{
		LabelSelector: config.KubeZLabelsKey + "=" + config.KubeZLabelsValue,
	}
	// 查询集群列表
	querySecretList, err := InclusterClientSet.CoreV1().Secrets(config.KubeZNamespace).List(context.TODO(), listOptions)
	if err != nil {
		logs.Error(map[string]interface{}{"error": err.Error()}, "查询集群列表失败")
		return
	}
	// 取出kubecofnig
	for _, v := range querySecretList.Items {
		clusterid := v.Name
		kubeconfig := string(v.Data["clusterconfig"])
		decodeKubeconfig, _ := cryptos.Decrypt(kubeconfig)
		logs.Info(nil, "正在生成"+clusterid+"的clusterset")
		clientset, err := goclient.InitKubeClient(decodeKubeconfig)
		if err != nil {
			logs.Error(map[string]interface{}{"error": err.Error()}, "生成"+clusterid+"的clusterset失败")
		} else {
			logs.Info(nil, "生成"+clusterid+"的clusterset成功")
			OutOfClusterClientSet[clusterid] = clientset
		}

	}
}

// 初始化内部集群的clientset
func initInClusterClientSet() {

	if config.KubeConfig == "" {
		// 没有kubeconfig文件,采用incluster方式
		logs.Debug(nil, "在配置文件中没有检测到kubeconfig,使用incluster模式")
		// todo
	} else {
		// 有kubeconfig文件,可能是docker部署或者外部集群
		logs.Debug(nil, "在配置文件中检测到kubeconfig,使用外部集群模式")
		decodeKubeconfig, err := base64s.DeCode(config.KubeConfig)
		if err != nil {
			logs.Panic(map[string]interface{}{"error": err.Error()}, "解码kubeconfig失败")
		}
		InclusterClientSet, err = goclient.InitKubeClient(decodeKubeconfig)
		if err != nil {
			logs.Panic(map[string]interface{}{"error": err.Error()}, "程序初始化InclusterKubeSet失败")
		}

	}
	// 查询集群版本
	serverVersion, err := InclusterClientSet.Discovery().ServerVersion()
	if err == nil {
		logs.Info(nil, "集群版本 : "+serverVersion.String())
	} else {
		logs.Error(map[string]interface{}{"error": err.Error()}, "获取集群版本失败,请检查网络或集群配置信息")
	}

}

func init() {
	initInClusterClientSet()
	initOutOfClusterClientSet()
}
