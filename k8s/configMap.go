package k8s

import (
	"kboard/config"
	"kboard/k8s/resource"
)

type IConfigMap interface {
	IK8sCore
}

type ConfigMap struct {
	K8sCore
}

func NewConfigMap(Config config.IConfig) IConfigMap {
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
