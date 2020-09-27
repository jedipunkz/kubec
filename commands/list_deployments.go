package commands

import (
	"fmt"

	"github.com/mitchellh/cli"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ListDeployments is struct for Listing of Deployments
type ListDeployments struct {
	UI cli.Ui
}

var (
	namespace string
)

// Run function for listing deployments
func (c *ListDeployments) Run(args []string) int {
	config := loadConfig()

	clientset, err := kubernetes.NewForConfig(config)

	deploymentsClient := clientset.AppsV1().Deployments(namespace)
	list, err := deploymentsClient.List(metav1.ListOptions{})
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	fmt.Printf("Thare are %d deployment(s) in the cluster\n", len(list.Items))
	for _, d := range list.Items {
		fmt.Printf(" * %s (%d replicas\n", d.Name, *d.Spec.Replicas)
	}
	return 0
}
