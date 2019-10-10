package k8s

import (
	"kboard/config"
	"kboard/internal"
)

type IDeployment interface {
	internal.IK8sCore
}

type Deployment struct {
	internal.K8sCore
}

func NewDeployment(Config config.IConfig) IDeployment {
	return &Deployment{
		internal.K8sCore{
			Config: Config,
			Kind:   internal.RESOURCE_DEPLOYMENT,
			Urls: internal.Urls{
				Read:   "/apis/apps/v1/namespaces/%s/deployments/%s",
				Create: "/apis/apps/v1/namespaces/%s/deployments",
			},
		},
	}
}
