package k8s

import (
	"github.com/revel/config"
	"resource"
)

type ISecret interface {
	IK8sCore
}

type Secret struct {
	K8sCore
}


func NewSecret(Config *config.Context) *Secret {
	return &Secret{
		K8sCore{
			Config: Config,
			Kind: resource.RESOURCE_SECRET,
			Urls: Urls{
				Read: "/api/v1/namespaces/%s/secrets/%s",
				Create: "/api/v1/namespaces/%s/secrets",
			},
		},
	}
}



