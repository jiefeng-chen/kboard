package k8s

import (
	"kboard/k8s/resource"
	"github.com/revel/config"
)

type IService interface {
	IK8sCore
}

type Service struct {
	K8sCore
}

func NewService(Config *config.Context) *Service {
	return &Service{
		K8sCore{
			Config: Config,
			Kind:   resource.RESOURCE_SERVICE,
			Urls: Urls{
				Read:   "/api/v1/namespaces/%s/services/%s",
				Create: "/api/v1/namespaces/%s/services",
			},
		},
	}
}




