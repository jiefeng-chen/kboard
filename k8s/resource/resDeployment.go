package resource

import (
	"gopkg.in/yaml.v2"
	"kboard/exception"
)

type IResDeployment interface {
	IResource
	SetMetadataName(string) error
	SetNamespace(string) error
	GetNamespace() string
	SetMatchLabels(map[string]string) error
	AddContainer(*Container) error
	SetTemplateLabels(map[string]string) error
}

type ResDeployment struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name      string
		Namespace string
		Labels    map[string]string
	}
	Spec struct {
		Selector *Selector // 圈定deployment管理的pod范围 跟下面的.spec.template.metadata.labels 匹配
		Template struct{ // pod模板，跟pod有一模一样的schema，但是不需要apiVersion和kind字段
			Metadata struct {
				Labels map[string]string
			}
			Spec struct {
				Containers []*Container
			}
		}
		Replicas string // replica副本数
	}
}

func NewResDeployment() *ResDeployment {
	return &ResDeployment{
		ApiVersion: "extensions/v1beta1",
		Kind:       RESOURCE_DEPLOYMENT,
		Metadata: struct {
			Name      string
			Namespace string
			Labels    map[string]string
		}{
			Name: "",
			Namespace: "",
			Labels: map[string]string{},
		},
		Spec: struct {
			Selector *Selector
			Template struct {
				Metadata struct{ Labels map[string]string };
				Spec     struct{ Containers []*Container }
			}
			Replicas string
		}{
			Selector: &Selector{
				MatchLabels: map[string]string{},
				MatchExpressions: nil,
			},
			Template: struct {
			Metadata struct{ Labels map[string]string }
			Spec     struct{ Containers []*Container }
		}{
			Metadata: struct{ Labels map[string]string }{Labels: map[string]string{}},
			Spec: struct{ Containers []*Container }{Containers: nil}},
			Replicas: ""},
	}
}

func (r *ResDeployment) SetMetadataName(name string) error {
	if name == "" {
		return exception.NewError("name is empty")
	}

	r.Metadata.Name = name

	return nil
}

func (r *ResDeployment) SetNamespace(nsName string) error {
	if nsName == "" {
		return exception.NewError("namespace is empty")
	}
	r.Metadata.Namespace = nsName

	return nil
}

func (r *ResDeployment) GetNamespace() string {
	return r.Metadata.Namespace
}

func (r *ResDeployment) AddContainer(container *Container) error {
	if container == nil {
		return exception.NewError("container is nil")
	}

	r.Spec.Template.Spec.Containers = append(r.Spec.Template.Spec.Containers, container)

	return nil
}

func (r *ResDeployment) SetMatchLabels(labels map[string]string) error {
	if len(labels) > 0 {
		for k,v := range labels {
			if k == "" || v == "" {
				return exception.NewError("match labels is empty")
			}
			r.Spec.Selector.MatchLabels[k] = v
		}

		return nil
	}else{
		return exception.NewError("labels is empty")
	}
}

func (r *ResDeployment) SetTemplateLabels(labels map[string]string) error {
	if len(labels) > 0 {
		for k, v := range labels {
			r.Spec.Template.Metadata.Labels[k] = v
		}
		return nil
	}
	return exception.NewError("labels are empty")
}

func (r *ResDeployment) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}

