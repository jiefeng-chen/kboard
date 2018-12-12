package k8s

import (
	"github.com/bitly/go-simplejson"
	"github.com/revel/config"
	"kboard/k8s/resource"
)

type IJob interface {
	IK8sCore
	Nodes() (*simplejson.Json, *HttpError)
}

type Job struct {
	K8sCore
}

func NewJob(Config *config.Context) *Job {
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


