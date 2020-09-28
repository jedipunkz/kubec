package commands

import (
	"fmt"
	mykube "kubec/pkg/kubec"
	"strconv"
	"strings"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const helpListDeployments = `
Usage: kubec list deployments
`

// Help is func for Help
func (c *ListDeployments) Help() string {
	return strings.TrimSpace(helpListDeployments)
}

// Synopsis is func for Synopsis
func (c *ListDeployments) Synopsis() string {
	return "listing deployments"
}

// Run function for listing deployments
func (c *ListDeployments) Run(args []string) int {
	k := mykube.NewKubeClient()

	clientset, err := kubernetes.NewForConfig(k.Config)

	deploymentsClient := clientset.AppsV1().Deployments(namespace)
	list, err := deploymentsClient.List(v1.ListOptions{})
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	fmt.Printf("Thare are %d deployment(s) in the cluster\n", len(list.Items))
	for _, d := range list.Items {
		c.UI.Output(" * " + d.Name + " (" + strconv.Itoa(int(*d.Spec.Replicas)) + " replicas)")
	}
	return 0
}
