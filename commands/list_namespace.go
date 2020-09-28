package commands

import (
	"fmt"
	mykube "kubec/pkg/kubec"
	"strings"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const helpListNamespaces = `
Usage: kubec list namespaces
`

// Help is func for Help
func (c *ListNamespaces) Help() string {
	return strings.TrimSpace(helpListNamespaces)
}

// Synopsis is func for Synopsis
func (c *ListNamespaces) Synopsis() string {
	return "listing namespaces"
}

// Run function for listing deployments
func (c *ListNamespaces) Run(args []string) int {
	k := mykube.NewKubeClient()

	clientset, err := kubernetes.NewForConfig(k.Config)

	pods, err := clientset.CoreV1().Namespaces().List(v1.ListOptions{})
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	fmt.Printf("Thare are %d deployment(s) in the cluster\n", len(pods.Items))
	for _, d := range pods.Items {
		c.UI.Output(" * " + d.Name)
	}
	return 0
}
