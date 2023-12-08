package controllers

import (
	"kubez_project/config"
	"kubez_project/utils/base64s"
	goclient "kubez_project/utils/go_client"
	"kubez_project/utils/logs"

	"k8s.io/client-go/kubernetes"
)

var InclusterKubeSet *kubernetes.Clientset

// 初始化控制器
func init() {

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
		InclusterKubeSet, err = goclient.InitKubeClient(decodeKubeconfig)
		if err != nil {
			logs.Panic(map[string]interface{}{"error": err.Error()}, "程序初始化InclusterKubeSet失败")
		}

	}
	serverVersion, err := InclusterKubeSet.Discovery().ServerVersion()
	if err != nil {
		logs.Error(map[string]interface{}{"error": err.Error()}, "获取集群版本失败,请检查网络或集群配置信息")
	} else {
		logs.Info(nil, "集群版本 : "+serverVersion.String())
	}

}
