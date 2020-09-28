package commands

import "github.com/mitchellh/cli"

// ListDeployments is struct for Listing of Deployments
type ListDeployments struct {
	UI cli.Ui
}

// ListPods is struct for Listing of Pods
type ListPods struct {
	UI cli.Ui
}

var (
	namespace string
)
