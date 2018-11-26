package resource

import "gopkg.in/yaml.v2"

type IResHorizontalPodAutoscaler interface {
	IResource
	SetMetaDataName(string) error
	SetNamespace(string) error
	GetNamespace() string
	SetMatchLabels([]map[string]string) error
	SetTolerations([]map[string]string) error
	SetContainers([]map[string]interface{}) error
	SetTerminationGracePeriodSeconds(string) error
	SetVolumes([]map[string]string, string) error
	SetRestartPolicy(string) error
	SetNodeSelector([]map[string]string) error
}

type ResHorizontalPodAutoscaler struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name      string
		Namespace string
		Labels    map[string]string
	}
	Spec struct {
		Selector struct {
			MatchLabels map[string]string `yaml:""`
		}
		ScaleTargetRef *ScaleTargetRef `yaml:"scaleTargetRef"`
		TargetCPUUtilizationPercentage int `yaml:"targetCPUUtilizationPercentage"`
		MinReplicas int `yaml:"minReplicas"`
		MaxReplicas int `yaml:"maxReplicas"`
		Metrics *[]metric
	}
}

type ScaleTargetRef struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind string
	Name string
}

const (
	METRIC_TYPE_RESOURCE = "Resource"
	METRIC_TYPE_PODS = "Pods"
	METRIC_TYPE_OBJECT = "Object"
	METRIC_TYPE_EXTERNAL = "External"
)

type metric interface {

}

type MetricResource struct {
	metric
	Type string
	Resource struct{
		Name string
		TargetAverageUtilization int `yaml:"targetAverageUtilization"`
	}
}

type MetricPods struct {
	metric
	Type string
	Pods struct{
		MetricName string `yaml:"metricName"`
		TargetAverageValue string `yaml:"targetAverageValue"`
	}
}

type MetricObject struct {
	metric
	Type string
	Object struct{
		MetricName string `yaml:"metricName"`
		Target struct{
			ApiVersion string `yaml:"apiVersion"`
			Kind string
			Name string
		}
		TargetValue string `yaml:"targetValue"`
	}
}

func NewResHorizontalPodAutoscaler() *ResHorizontalPodAutoscaler {
	return &ResHorizontalPodAutoscaler{
		ApiVersion: "autoscaling/v2beta1",
		Kind:       RESOURCE_HORIZONTAL_POD_AUTOSCALER,
		Metadata: struct {
			Name      string
			Namespace string
			Labels    map[string]string
		}{Name: "", Namespace: "", Labels: map[string]string{}},
	}
}

func (r *ResHorizontalPodAutoscaler) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}

