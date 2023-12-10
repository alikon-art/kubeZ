package goclient

import (
	"fmt"
	"strings"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 从kubeconfig变量初始化clientset的方法
// 传入解密后的string类型kubeconfig
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

// 将 map[string]string形式的label转为符合LabelSelector的格式
func Convert2LabelSelector(labels map[string]string) string {
	var selectorParts []string

	for key, value := range labels {
		selectorParts = append(selectorParts, fmt.Sprintf("%s=%s", key, value))
	}

	return strings.Join(selectorParts, ",")
}
