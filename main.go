package main

import (
	"fmt"
	utils "github.com/cuisongliu/k8s-mysql/cmd/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	clientSet, _ := utils.KubeClientFromKubeconfigPath("C:\\Users\\cuisongliu\\.kube\\config")
	list, _ := clientSet.CoreV1().Namespaces().List(metav1.ListOptions{})
	for _, d := range list.Items {
		fmt.Printf("Namespace: %s\n", d.Name)
	}

	utils.ToJson()
	clientSetNew := utils.KubeClient()

	listNew, _ := clientSetNew.CoreV1().Nodes().List(metav1.ListOptions{})
	for _, d := range listNew.Items {
		fmt.Printf("NodeName: %s\n", d.Name)
	}
}
