package k8s

import (
	"kboard/config"
	"kboard/k8s/resource"
)

type IJob interface {
	IK8sCore
}

type Job struct {
	K8sCore
}

func NewJob(Config *config.Config) IJob {
	return &Job{
		K8sCore{
			Config: Config,
			Kind: resource.RESOURCE_JOB,
			Urls: Urls{
				Read: "/apis/batch/v1/namespaces/%s/jobs/%s",
				Create: "/apis/batch/v1/namespaces/%s/jobs",
			},
		},
	}
}


