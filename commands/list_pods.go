package commands

import (
	"context"
	"fmt"
	mykube "kubec/pkg/kubec"
	"strings"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const helpListPods = `
Usage: kubec list pods
`

// Help is func for Help
func (c *ListPods) Help() string {
	return strings.TrimSpace(helpListPods)
}

// Synopsis is func for Synopsis
func (c *ListPods) Synopsis() string {
	return "listing pods"
}

// Run function for listing pods
func (c *ListPods) Run(args []string) int {
	k := mykube.NewKubeClient()

	clientset, err := kubernetes.NewForConfig(k.Config)

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	fmt.Printf("Thare are %d pod(s) in the cluster\n", len(pods.Items))
	for _, d := range pods.Items {
		c.UI.Output(" * " + d.Name)
	}
	return 0
}
