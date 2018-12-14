package resource

import (
	"kboard/exception"
	"gopkg.in/yaml.v2"
)

type IResNamespace interface {
	IResource
	SetMetadataName(string) error
	SetNamespace(string) error
	SetLabels(map[string]string) error
}

// pod结构体
type ResNamespace struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind string
	Metadata struct{
		Name string
		Namespace string
		Labels map[string]string
		Annotations map[string]string
	}
}

func NewResNamespace() *ResNamespace {
	return &ResNamespace{
		ApiVersion: "v1",
		Kind: RESOURCE_NAMESPACE,
		Metadata: struct {
			Name        string
			Namespace   string
			Labels      map[string]string
			Annotations map[string]string
		}{Name: "", Namespace: "", Labels: map[string]string{}, Annotations: map[string]string{}},
	}
}

func (r *ResNamespace) SetMetadataName(name string) error {
	if name == "" {
		return exception.NewError("name is empty")
	}
	r.Metadata.Name = name
	return nil
}

func (r *ResNamespace) SetNamespace(ns string) error {
	if ns == "" {
		return exception.NewError("namespace is empty")
	}
	r.Metadata.Namespace = ns
	return nil
}

func (r *ResNamespace) SetLabels(labels map[string]string) error {
	if len(labels) <= 0 {
		return exception.NewError("labels is empty")
	}
	for k, v := range labels {
		if k == "" || v == "" {
			return exception.NewError("label's key or value is empty")
		}
		r.Metadata.Labels[k] = v
	}
	return nil
}

func (r *ResNamespace) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}

