package resource

import (
	"kboard/internal"

	"gopkg.in/yaml.v2"
)

type IResNamespace interface {
	internal.IResource
	SetMetadataName(string) error
	SetNamespace(string) error
	SetLabels(map[string]string) error
}

// pod结构体
type ResNamespace struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name        string
		Namespace   string
		Labels      map[string]string
		Annotations map[string]string
	}
}

func NewResNamespace() *ResNamespace {
	return &ResNamespace{
		ApiVersion: "v1",
		Kind:       internal.RESOURCE_NAMESPACE,
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
		return internal.NewError("name is empty")
	}
	r.Metadata.Name = name
	return nil
}

func (r *ResNamespace) SetNamespace(ns string) error {
	if ns == "" {
		return internal.NewError("namespace is empty")
	}
	r.Metadata.Namespace = ns
	return nil
}

func (r *ResNamespace) SetLabels(labels map[string]string) error {
	if len(labels) <= 0 {
		return internal.NewError("labels is empty")
	}
	for k, v := range labels {
		if k == "" || v == "" {
			return internal.NewError("label's key or value is empty")
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
