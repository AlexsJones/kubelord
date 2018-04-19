package kubernetes

import (
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

//Configuration holds points to native and lib types
type Configuration struct {
	Clientset *kubernetes.Clientset
	config    *rest.Config
}

//NewConfiguration provides a kubernetes interface
func NewConfiguration(masterURL string, inclusterConfig bool) (*Configuration, error) {

	//InCluster...
	var config *rest.Config
	var err error
	if inclusterConfig {
		log.Println("Using in cluster configuration")
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
	} else {
		log.Println("Using out of cluster configuration")
		kubeconfig := filepath.Join(func() string {
			if h := os.Getenv("HOME"); h != "" {
				return h
			}
			return os.Getenv("USERPROFILE")
		}(), ".kube", "config")
		// use the current context in kubeconfig
		config, err = clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
		if err != nil {
			panic(err.Error())
		}
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &Configuration{Clientset: clientset, config: config}, nil
}
