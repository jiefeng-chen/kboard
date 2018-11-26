package k8s


import (
	"github.com/revel/config"
	"dashboard/resource"
)

type IDaemonSet interface {
	IK8sCore
}

type DaemonSet struct {
	K8sCore
}


func NewDaemonSet(Config *config.Context) *DaemonSet {
	return &DaemonSet{
		K8sCore{
			Config: Config,
			Kind: resource.RESOURCE_INGRESS,
			Urls: Urls{
				Read: "/apis/apps/v1beta2/namespaces/%s/daemonsets/%s",
				Create: "/apis/apps/v1beta2/namespaces/%s/daemonsets",
			},
		},
	}
}
