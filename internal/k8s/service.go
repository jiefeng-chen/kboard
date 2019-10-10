package k8s

import (
	"kboard/config"
	"kboard/internal"
)

type IService interface {
	internal.IK8sCore
}

type Service struct {
	internal.K8sCore
}

func NewService(Config config.IConfig) IService {
	return &Service{
		internal.K8sCore{
			Config: Config,
			Kind:   internal.RESOURCE_SERVICE,
			Urls: internal.Urls{
				Read:   "/api/v1/namespaces/%s/services/%s",
				Create: "/api/v1/namespaces/%s/services",
			},
		},
	}
}
