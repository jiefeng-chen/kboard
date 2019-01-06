package resource

import (
	"gopkg.in/yaml.v2"
	"kboard/exception"
)

/**
	创建自定义资源，
**/

type IResCustomResourceDefinition interface {
	IResource
	SetMetadataName(string) error
	SetNamespace(string) error
	GetNamespace() string
}

type ResCustomResourceDefinition struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name      string
		Namespace string
		Annotations map[string]string
		Labels      map[string]string
	}
	Spec *CustomResourceDefinitionSpec
}

type CustomResourceDefinitionSpec struct {

}

func NewCustomResourceDefinition() *ResCustomResourceDefinition {
	return &ResCustomResourceDefinition{
		ApiVersion: "apiextensions.k8s.io/v1beta1",
		Kind:       RESOURCE_CUSTOM_RESOURCE_DEFINITION,
		Metadata: struct {
			Name        string
			Namespace   string
			Annotations map[string]string
			Labels      map[string]string
		}{Name: "", Namespace: "", Annotations: map[string]string{}, Labels: map[string]string{}},
		Spec: nil,
	}
}

func (r *ResCustomResourceDefinition) SetMetadataName(name string) error {
	if name == "" {
		return exception.NewError("name is empty")
	}
	r.Metadata.Name = name
	return nil
}

func (r *ResCustomResourceDefinition) SetNamespace(ns string) error {
	if ns == "" {
		return exception.NewError("namespace is empty")
	}
	r.Metadata.Namespace = ns
	return nil
}

func (r *ResCustomResourceDefinition) GetNamespace() string {
	return r.Metadata.Namespace
}


func (r *ResCustomResourceDefinition) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}



