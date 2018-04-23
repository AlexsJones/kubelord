package kubernetes

import (
	"log"
	"os"
	"path/filepath"

	"k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	//Required to work with gcp
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

//Configuration holds points to native and lib types
type Configuration struct {
	clientset *kubernetes.Clientset
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
	return &Configuration{clientset: clientset, config: config}, nil
}

//GetNamespace within kubernetes
func (k *Configuration) GetNamespace(namespace string) (*v1.Namespace, error) {

	ns, err := k.clientset.CoreV1().Namespaces().Get(namespace, meta.GetOptions{})
	return ns, err
}

//GetNamespaces within kubernetes
func (k *Configuration) GetNamespaces() (*v1.NamespaceList, error) {

	nl, err := k.clientset.CoreV1().Namespaces().List(meta.ListOptions{})

	return nl, err
}

//GetPods within kubernetes
func (k *Configuration) GetPods(namespace string) (*v1.PodList, error) {

	nl, err := k.clientset.CoreV1().Pods(namespace).List(meta.ListOptions{})

	return nl, err
}
