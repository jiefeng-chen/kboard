package k8s

import (
	"kboard/config"
	"kboard/internal"
)

type ISecret interface {
	internal.IK8sCore
}

type Secret struct {
	internal.K8sCore
}

func NewSecret(Config config.IConfig) ISecret {
	return &Secret{
		internal.K8sCore{
			Config: Config,
			Kind:   internal.RESOURCE_SECRET,
			Urls: internal.Urls{
				Read:   "/api/v1/namespaces/%s/secrets/%s",
				Create: "/api/v1/namespaces/%s/secrets",
			},
		},
	}
}
