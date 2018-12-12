package k8s

import (
	"kboard/k8s/resource"
	"github.com/revel/config"
)

type IReplicaSet interface {
	IK8sCore
}

type ReplicaSet struct {
	K8sCore
}

func NewReplicaSet(Config *config.Context) *ReplicaSet {
	return &ReplicaSet{
		K8sCore{
			Config: Config,
			Kind:   resource.RESOURCE_REPLICASET,
			Urls: Urls{
				Read:   "/apis/apps/v1/namespaces/%s/replicasets/%s",
				Create: "/apis/apps/v1/namespaces/%s/replicasets",
			},
		},
	}
}


