package k8s

import (
	"kboard/config"
	"kboard/k8s/resource"
)

type IHorizontalPodAutoscaler interface {
	IK8sCore
}

type HorizontalPodAutoscaler struct {
	K8sCore
}

func NewHorizontalPodAutoscaler(Config config.IConfig) IHorizontalPodAutoscaler {
	return &HorizontalPodAutoscaler{
		K8sCore{
			Config: Config,
			Kind:   resource.RESOURCE_HORIZONTAL_POD_AUTOSCALER,
			Urls: Urls{
				Read:   "/apis/autoscaling/v1/namespaces/%s/horizontalpodautoscalers/%s",
				Create: "/apis/autoscaling/v1/namespaces/%s/horizontalpodautoscalers",
			},
		},
	}
}
