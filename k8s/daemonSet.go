package k8s

import (
	"kboard/config"
	"kboard/k8s/resource"
)

type IDaemonSet interface {
	IK8sCore
}

type DaemonSet struct {
	K8sCore
}

func NewDaemonSet(Config *config.Config) IDaemonSet {
	return &DaemonSet{
		K8sCore{
			Config: Config,
			Kind:   resource.RESOURCE_DAEMONSET,
			Urls: Urls{
				Read:   "/apis/apps/v1beta2/namespaces/%s/daemonsets/%s",
				Create: "/apis/apps/v1beta2/namespaces/%s/daemonsets",
			},
		},
	}
}
