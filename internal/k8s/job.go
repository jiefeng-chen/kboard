package k8s

import (
	"kboard/config"
	"kboard/internal"
)

type IJob interface {
	internal.IK8sCore
}

type Job struct {
	internal.K8sCore
}

func NewJob(Config config.IConfig) IJob {
	return &Job{
		internal.K8sCore{
			Config: Config,
			Kind:   internal.RESOURCE_JOB,
			Urls: internal.Urls{
				Read:   "/apis/batch/v1/namespaces/%s/jobs/%s",
				Create: "/apis/batch/v1/namespaces/%s/jobs",
			},
		},
	}
}
