package k8s

import (
	"kboard/config"
	"kboard/internal"
)

type IReplicaSet interface {
	internal.IK8sCore
}

type ReplicaSet struct {
	internal.K8sCore
}

func NewReplicaSet(Config config.IConfig) IReplicaSet {
	return &ReplicaSet{
		internal.K8sCore{
			Config: Config,
			Kind:   internal.RESOURCE_REPLICASET,
			Urls: internal.Urls{
				Read:   "/apis/apps/v1/namespaces/%s/replicasets/%s",
				Create: "/apis/apps/v1/namespaces/%s/replicasets",
			},
		},
	}
}
