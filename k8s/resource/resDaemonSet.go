package resource

import (
	"gopkg.in/yaml.v2"
	"kboard/exception"
)

type IResDaemonSet interface {
	IResource
	SetMetaDataName(string) error
	SetNamespace(string) error
	GetNamespace() string
	SetMatchLabels([]map[string]string) error
	SetTolerations([]map[string]string) error
	AddContainer(IContainer) error
	SetTerminationGracePeriodSeconds(string) error
	SetVolume(*Volume) error
	SetRestartPolicy(string) error
	SetNodeSelector([]map[string]string) error
}

type ResDaemonSet struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name      string
		Namespace string
		Labels    map[string]string
	}
	Spec *DaemonSetSpec
}

func NewResDaemonSet() IResDaemonSet {
	return &ResDaemonSet{
		ApiVersion: "extensions/v1beta1",
		Kind:       RESOURCE_DAEMONSET,
		Metadata: struct {
			Name      string
			Namespace string
			Labels    map[string]string
		}{Name: "", Namespace: "", Labels: map[string]string{}},
		Spec: &DaemonSetSpec{
			Selector: &Selector{
				MatchLabels:      map[string]string{},
				MatchExpressions: nil,
			},
			Template: &DaemonSetSpecTemplate{
				Metadata: struct{ Labels map[string]string }{
					Labels: map[string]string{}},
				Spec: &DaemonSetTemplateSpec{
					Tolerations:                   []*DaemonSetToleration{},
					Containers:                    []IContainer{},
					TerminationGracePeriodSeconds: "",
					Volumes:          []*Volume{},
					RestartPolicy:    "Always",
					ImagePullSecrets: map[string]string{},
					NodeSelector:     map[string]string{},
				},
			},
		},
	}
}

type DaemonSetSpec struct {
	Selector *Selector
	Template *DaemonSetSpecTemplate
}

type DaemonSetSpecTemplate struct {
	Metadata struct {
		Labels map[string]string
	}
	Spec *DaemonSetTemplateSpec
}

type DaemonSetTemplateSpec struct {
	Tolerations                   []*DaemonSetToleration
	Containers                    []IContainer
	TerminationGracePeriodSeconds string `yaml:"terminationGracePeriodSeconds"`
	Volumes                       []*Volume
	RestartPolicy                 string            `yaml:"restartPolicy"` // 默认为 Always
	ImagePullSecrets              map[string]string `yaml:"imagePullSecrets"`
	NodeSelector                  map[string]string `yaml:"nodeSelector"`
}

type DaemonSetToleration struct {
	Key               string
	Effect            string
	Value             string
	Operator          string
	TolerationSeconds string `yaml:"tolerationSeconds"`
}

func (r *ResDaemonSet) SetMetaDataName(name string) error {
	if name == "" {
		return exception.NewError("metadata name is empty")
	}
	// 设置 .metadata.name
	r.Metadata.Name = name
	// 设置 .spec.selector.matchLabels.name
	r.Spec.Selector.MatchLabels["name"] = name
	// 设置 .spec.template.metadata.labels.name
	r.Spec.Template.Metadata.Labels["name"] = name
	return nil
}

func (r *ResDaemonSet) SetNamespace(ns string) error {
	if ns == "" {
		return exception.NewError("metadata namespace is empty")
	}
	r.Metadata.Namespace = ns
	return nil
}

func (r *ResDaemonSet) GetNamespace() string {
	return r.Metadata.Namespace
}

func (r *ResDaemonSet) SetMatchLabels(labels []map[string]string) error {
	if len(labels) > 0 {
		for _, v := range labels {
			for key, val := range v {
				if key == "" {
					return exception.NewError("matchLabels key is empty")
				}

				r.Spec.Selector.MatchLabels[key] = val
			}
		}
	}
	return nil
}

func (r *ResDaemonSet) SetTolerations(tolers []map[string]string) error {
	if len(tolers) > 0 {
		for _, v := range tolers {
			if v["key"] == "" {
				return exception.NewError("toleration key is empty")
			}
			toler := &DaemonSetToleration{
				Key:               v["key"],
				Operator:          v["operator"],
				Effect:            v["effect"],
				Value:             v["value"],
				TolerationSeconds: v["tolerationSeconds"],
			}
			r.Spec.Template.Spec.Tolerations = append(r.Spec.Template.Spec.Tolerations, toler)
		}
	}
	return nil
}

func (r *ResDaemonSet) AddContainer(container IContainer) error {
	if container == nil {
		return exception.NewError("container is nil")
	}
	r.Spec.Template.Spec.Containers = append(r.Spec.Template.Spec.Containers, container)
	return nil
}

func (r *ResDaemonSet) SetTerminationGracePeriodSeconds(second string) error {
	if second == "" {
		return exception.NewError("termination grace period seconds is empty")
	}
	r.Spec.Template.Spec.TerminationGracePeriodSeconds = second
	return nil
}

type VolumeHostPath struct {
	Name     string
	HostPath struct {
		Path string
	} `yaml:"hostPath"`
}

type VolumeConfigMap struct {
	Name      string
	ConfigMap struct {
		Name string
	} `yaml:"configMap"`
}

type VolumeSecret struct {
}

type VolumeEmptyDir struct {
	Name     string
	EmptyDir struct{}
}

type VolumePersistentVolumeClaim struct {
	Name                  string
	PersistentVolumeClaim struct {
		ClaimName string `yaml:"claimName"`
	} `yaml:"persistentVolumeClaim"`
}

func (r *ResDaemonSet) SetVolume(vol *Volume) error {
	if vol == nil {
		return exception.NewError("volume is nil")
	}
	r.Spec.Template.Spec.Volumes = append(r.Spec.Template.Spec.Volumes, vol)
	return nil
}

func (r *ResDaemonSet) SetRestartPolicy(rPolicy string) error {
	if rPolicy == "" {
		return exception.NewError("restart policy is empty")
	}
	r.Spec.Template.Spec.RestartPolicy = rPolicy
	return nil
}

func (r *ResDaemonSet) SetNodeSelector(selectors []map[string]string) error {
	if len(selectors) > 0 {
		for _, v := range selectors {
			if v["key"] == "" {
				return exception.NewError("node selector's key is empty")
			}
			r.Spec.Template.Spec.NodeSelector[v["key"]] = v["val"]
		}
	}
	return nil
}

func (r *ResDaemonSet) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}
