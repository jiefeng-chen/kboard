package k8s

import (
	"kboard/config"
	"kboard/internal"
)

type IHorizontalPodAutoscaler interface {
	internal.IK8sCore
}

type HorizontalPodAutoscaler struct {
	internal.K8sCore
}

func NewHorizontalPodAutoscaler(Config config.IConfig) IHorizontalPodAutoscaler {
	return &HorizontalPodAutoscaler{
		internal.K8sCore{
			Config: Config,
			Kind:   internal.RESOURCE_HORIZONTAL_POD_AUTOSCALER,
			Urls: internal.Urls{
				Read:   "/apis/autoscaling/v1/namespaces/%s/horizontalpodautoscalers/%s",
				Create: "/apis/autoscaling/v1/namespaces/%s/horizontalpodautoscalers",
			},
		},
	}
}
