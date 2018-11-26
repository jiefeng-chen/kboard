package k8s

import (
	"github.com/revel/config"
	"dashboard/resource"
)

type ICronJob interface {
	IK8sCore
}

type CronJob struct {
	K8sCore
}


func NewCronJob(Config *config.Context) *CronJob {
	return &CronJob{
		K8sCore{
			Config: Config,
			Kind: resource.RESOURCE_INGRESS,
			Urls: Urls{
				Read: "/apis/batch/v1beta1/namespaces/%s/cronjobs/%s",
				Create: "/apis/batch/v1beta1/namespaces/%s/cronjobs",
			},
		},
	}
}