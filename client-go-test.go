package main

import (
	"context"
	"encoding/json"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := "./config/kubeconfig"
	// 1.初始化config实例
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// 2.生成客户端工具集
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 3. 操作集群
	// 创建ns
	var newNameSpace corev1.Namespace
	newNameSpace.Name = "test1"
	_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), &newNameSpace, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}

	// // 创建deploy
	// newDeployment := &appsv1.Deployment{}
	// newDeployment.Name = "testdeploy"
	// labels := make(map[string]string) // 先创建map,直接赋值会报空指针,也便于下边统一赋值
	// labels["app"] = "testdeplpy"
	// labels["version"] = "1"
	// newDeployment.Labels = labels // 此处是deploy自身的label
	// newDeployment.Spec.Selector = &metav1.LabelSelector{}
	// newDeployment.Spec.Selector.MatchLabels = labels // 此处是deploy选择器的label
	// newDeployment.Spec.Template.Labels = labels      // 此处是pod模板的label
	// // 新建一个容器对象
	// newContainer := corev1.Container{
	// 	Name:  "testdeploy-pod",
	// 	Image: "testdeploy-pod",
	// }
	// // 使用append将容器对象添加到deploy对象中
	// newDeployment.Spec.Template.Spec.Containers = append(newDeployment.Spec.Template.Spec.Containers, newContainer)
	// // newDeployment.Spec.Template.Spec.Containers = []corev1.Container{}
	// // newDeployment.Spec.Template.Spec.Containers[0].Name = "testdeploy-pod"  // pod中容器的名字
	// // newDeployment.Spec.Template.Spec.Containers[0].Image = "testdeploy-pod" // pod中容器的镜像

	// clientset.AppsV1().Deployments("test1").Create(context.TODO(), newDeployment, metav1.CreateOptions{})

	deployjson := `{
		"kind": "Deployment",
		"apiVersion": "apps/v1",
		"metadata": {
			"name": "nginx",
			"creationTimestamp": null,
			"labels": {
				"app": "nginx"
			}
		},
		"spec": {
			"replicas": 1,
			"selector": {
				"matchLabels": {
					"app": "nginx"
				}
			},
			"template": {
				"metadata": {
					"creationTimestamp": null,
					"labels": {
						"app": "nginx"
					}
				},
				"spec": {
					"containers": [
						{
							"name": "test",
							"image": "test",
							"resources": {}
						}
					]
				}
			},
			"strategy": {}
		},
		"status": {}
	}



	`

	newDeployment2 := &appsv1.Deployment{}

	err = json.Unmarshal([]byte(deployjson), newDeployment2)
	if err != nil {
		fmt.Println(err.Error())
	}

	clientset.AppsV1().Deployments("default").Create(context.TODO(), newDeployment2, metav1.CreateOptions{})
}
