package k8s


import (
	"kboard/k8s/resource"
	"kboard/config"
)

type IStatefulSet interface {
	IK8sCore
}

type StatefulSet struct {
	K8sCore
}

func NewStatefulSet(Config *config.Config) *StatefulSet {
	return &StatefulSet{
		K8sCore{
			Config: Config,
			Kind:   resource.RESOURCE_STATEFULE_SET,
			Urls: Urls{
				Read:   "/apis/apps/v1/namespaces/%s/statefulsets/%s",
				Create: "/apis/apps/v1/namespaces/%s/statefulsets",
			},
		},
	}
}






