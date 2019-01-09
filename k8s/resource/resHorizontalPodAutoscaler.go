package resource

import (
	"gopkg.in/yaml.v2"
	"kboard/exception"
)

type IResHorizontalPodAutoscaler interface {
	IResource
	SetMetadataName(string) error
	SetNamespace(string) error
	GetNamespace() string
	SetMatchLabels(map[string]string) error
}

type ResHorizontalPodAutoscaler struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name      string
		Namespace string
		Labels    map[string]string
	}
	Spec *HPASpec
}

type HPASpec struct {
	Selector                       *Selector
	ScaleTargetRef                 *ScaleTargetRef `yaml:"scaleTargetRef"`
	TargetCPUUtilizationPercentage int             `yaml:"targetCPUUtilizationPercentage"`
	MinReplicas                    int             `yaml:"minReplicas"`
	MaxReplicas                    int             `yaml:"maxReplicas"`
	Metrics                        []*Metric
}

// 弹性伸缩目标
type ScaleTargetRef struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Name       string
}

const (
	METRIC_TYPE_RESOURCE = "Resource"
	METRIC_TYPE_PODS     = "Pods"
	METRIC_TYPE_OBJECT   = "Object"
	METRIC_TYPE_EXTERNAL = "External"
)

type Metric struct {
	Type     string
	Resource *MetricResource
	Pods     *MetricPods
	Object   *MetricObject
}

type MetricResource struct {
	Name                     string
	TargetAverageUtilization int `yaml:"targetAverageUtilization"`
}

type MetricPods struct {
	MetricName         string `yaml:"metricName"`
	TargetAverageValue string `yaml:"targetAverageValue"`
}

type MetricObject struct {
	MetricName  string `yaml:"metricName"`
	Target      *MetricObjectTarget
	TargetValue string `yaml:"targetValue"`
}

type MetricObjectTarget struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Name       string
}

func NewResHorizontalPodAutoscaler() *ResHorizontalPodAutoscaler {
	return &ResHorizontalPodAutoscaler{
		ApiVersion: "autoscaling/v2beta1", // k8s>=v1.7
		Kind:       RESOURCE_HORIZONTAL_POD_AUTOSCALER,
		Metadata: struct {
			Name      string
			Namespace string
			Labels    map[string]string
		}{
			Name:      "",
			Namespace: "",
			Labels:    map[string]string{},
		},
		Spec: &HPASpec{
			Selector: &Selector{
				MatchLabels:      map[string]string{},
				MatchExpressions: nil,
			},
			ScaleTargetRef:                 nil,
			TargetCPUUtilizationPercentage: 0,
			MinReplicas:                    0,
			MaxReplicas:                    0,
			Metrics:                        nil,
		},
	}
}

func (r *ResHorizontalPodAutoscaler) SetMetadataName(name string) error {
	if name == "" {
		return exception.NewError("name is empty")
	}

	r.Metadata.Name = name
	return nil
}

func (r *ResHorizontalPodAutoscaler) SetNamespace(ns string) error {
	if ns == "" {
		return exception.NewError("namespace is empty")
	}
	r.Metadata.Namespace = ns
	return nil
}

func (r *ResHorizontalPodAutoscaler) GetNamespace() string {
	return r.Metadata.Namespace
}

func (r *ResHorizontalPodAutoscaler) SetMatchLabels(labels map[string]string) error {
	if len(labels) > 0 {
		for k, v := range labels {
			if k == "" || v == "" {
				return exception.NewError("match labels is empty")
			}
			r.Spec.Selector.MatchLabels[k] = v
		}
	}

	return nil
}

func (r *ResHorizontalPodAutoscaler) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}
