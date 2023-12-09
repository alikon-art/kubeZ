package goclient

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 从kubeconfig变量初始化clientset的方法
func InitKubeClient(kubeconfig string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeconfig))
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, err
}

// 获取集群版本的方法
func GetClusterVersion(clientset *kubernetes.Clientset) (version string, err error) {
	serverVirsion, err := clientset.Discovery().ServerVersion()
	if err != nil {
		return "", err
	}
	return serverVirsion.String(), err
}
