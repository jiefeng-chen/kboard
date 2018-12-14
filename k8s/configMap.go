package k8s

import (
	"kboard/k8s/resource"
	"kboard/config"
)

type IConfigMap interface {
	IK8sCore
}

type ConfigMap struct {
	K8sCore
}

func NewConfigMap(Config *config.Config) *ConfigMap {
	return &ConfigMap{
		K8sCore{
			Config: Config,
			Kind:   resource.RESOURCE_CONFIG_MAP,
			Urls: Urls{
				Read:   "/api/v1/namespaces/%s/configmaps/%s",
				Create: "/api/v1/namespaces/%s/configmaps",
			},
		},
	}
}


