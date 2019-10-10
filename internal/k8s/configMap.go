package k8s

import (
	"kboard/config"
	"kboard/internal"
)

type IConfigMap interface {
	internal.IK8sCore
}

type ConfigMap struct {
	internal.K8sCore
}

func NewConfigMap(Config config.IConfig) IConfigMap {
	return &ConfigMap{
		internal.K8sCore{
			Config: Config,
			Kind:   internal.RESOURCE_CONFIG_MAP,
			Urls: internal.Urls{
				Read:   "/api/v1/namespaces/%s/configmaps/%s",
				Create: "/api/v1/namespaces/%s/configmaps",
			},
		},
	}
}
