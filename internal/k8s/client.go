package k8s

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Client wraps the Kubernetes clientset with app-specific helpers.
type Client struct {
	clientset *kubernetes.Clientset
	namespace string
}

// New returns a K8s client using in-cluster config or a local kubeconfig file.
func New(kubeconfig, namespace string, inCluster bool) (*Client, error) {
	var restCfg *rest.Config
	var err error

	if inCluster {
		restCfg, err = rest.InClusterConfig()
	} else {
		restCfg, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to build k8s config: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(restCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create k8s clientset: %w", err)
	}

	return &Client{clientset: clientset, namespace: namespace}, nil
}

// TODO: implement server lifecycle methods (Deploy, Delete, Scale, Status)
