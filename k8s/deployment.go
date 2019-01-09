package k8s

import (
	"kboard/config"
	"kboard/k8s/resource"
)

type IDeployment interface {
	IK8sCore
}

type Deployment struct {
	K8sCore
}

func NewDeployment(Config *config.Config) IDeployment {
	return &Deployment{
		K8sCore{
			Config: Config,
			Kind:   resource.RESOURCE_DEPLOYMENT,
			Urls: Urls{
				Read:   "/apis/apps/v1/namespaces/%s/deployments/%s",
				Create: "/apis/apps/v1/namespaces/%s/deployments",
			},
		},
	}
}
