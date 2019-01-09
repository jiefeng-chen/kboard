package k8s

import (
	"kboard/config"
	"kboard/k8s/resource"
)

type ISecret interface {
	IK8sCore
}

type Secret struct {
	K8sCore
}

func NewSecret(Config *config.Config) ISecret {
	return &Secret{
		K8sCore{
			Config: Config,
			Kind:   resource.RESOURCE_SECRET,
			Urls: Urls{
				Read:   "/api/v1/namespaces/%s/secrets/%s",
				Create: "/api/v1/namespaces/%s/secrets",
			},
		},
	}
}
