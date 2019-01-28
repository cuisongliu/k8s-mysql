package utils

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/Sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func KubeClient() *kubernetes.Clientset {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}

//KubeClientFromKubeconfigPath is
func KubeClientFromKubeconfigPath(path string) (clientSet *kubernetes.Clientset, err error) {
	cfg, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		return nil, fmt.Errorf("new kube config error: %s", err)
	}
	if clientSet, err = kubernetes.NewForConfig(cfg); err != nil {
		return nil, fmt.Errorf("new kube config error: %s", err)
	}
	return clientSet, nil
}

//KubeClientFromKubeconfigStringBody is
func KubeClientFromKubeconfigStringBody(body string) (*kubernetes.Clientset, error) {
	b, err := base64.StdEncoding.DecodeString(body)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("Fetch kubeconfig string: %s", string(b))
	clientConfig, err := clientcmd.NewClientConfigFromBytes(b)
	if err != nil {
		return nil, fmt.Errorf("new client config from body error: %s", err)
	}
	cfg, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("new kube config from body error: %s", err)
	}
	clientSet, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, fmt.Errorf("new kube config from body error: %s", err)
	}
	return clientSet, nil
}
