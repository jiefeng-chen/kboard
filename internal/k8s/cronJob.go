package k8s

import (
	"kboard/config"
	"kboard/internal"
)

type ICronJob interface {
	internal.IK8sCore
}

type CronJob struct {
	internal.K8sCore
}

func NewCronJob(Config config.IConfig) ICronJob {
	return &CronJob{
		internal.K8sCore{
			Config: Config,
			Kind:   internal.RESOURCE_INGRESS,
			Urls: internal.Urls{
				Read:   "/apis/batch/v1beta1/namespaces/%s/cronjobs/%s",
				Create: "/apis/batch/v1beta1/namespaces/%s/cronjobs",
			},
		},
	}
}
