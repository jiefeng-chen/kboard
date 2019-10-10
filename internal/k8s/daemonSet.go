package k8s

import (
	"kboard/config"
	"kboard/internal"
)

type IDaemonSet interface {
	internal.IK8sCore
}

type DaemonSet struct {
	internal.K8sCore
}

func NewDaemonSet(Config config.IConfig) IDaemonSet {
	return &DaemonSet{
		internal.K8sCore{
			Config: Config,
			Kind:   internal.RESOURCE_DAEMONSET,
			Urls: internal.Urls{
				Read:   "/apis/apps/v1beta2/namespaces/%s/daemonsets/%s",
				Create: "/apis/apps/v1beta2/namespaces/%s/daemonsets",
			},
		},
	}
}
