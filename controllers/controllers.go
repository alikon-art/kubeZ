package controllers

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 创建kube控制器的方法
func InitKubeClient(kubeconfig string) (*kubernetes.Clientset, error) {
	if err != nil {
	if err != nil {
		return nil, err
	}
	return clientset, err
<<<<<<< HEAD
=======
}

