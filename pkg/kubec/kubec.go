package mykube

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// KubeClient is struct for Kubernetes Client
type KubeClient struct {
	Config *rest.Config
}

// NewKubeClient is func ...
func NewKubeClient() *KubeClient {
	k := new(KubeClient)

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	k.Config, _ = clientcmd.BuildConfigFromFlags("", *kubeconfig)

	return k
}
