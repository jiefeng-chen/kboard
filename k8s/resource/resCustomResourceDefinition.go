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
}

type ResCustomResourceDefinition struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name      string // name must match the spec fields below, and be in the form: <plural>.<group>
	}
	Spec *CustomResourceDefinitionSpec
}

type CustomResourceDefinitionSpec struct {
	Group string // group name to use for REST API: /apis/<group>/<version>
	Version []*CrdVersion
	Scope string // either Namespaced or Cluster
	Names *CrdNames
}

type CrdNames struct {
	Plural string // plural name to be used in the URL: /apis/<group>/<version>/<plural>
	Singular string // singular name to be used as an alias on the CLI and for display
	Kind string // kind is normally the CamelCased singular type. Your resource manifests use this
	ShortNames []string // shortNames allow shorter string to match your resource on the CLI
}

type CrdVersion struct {
	Name string
	Served bool // Each version can be enabled/disabled by Served flag
	Storage bool // One and only one version must be marked as the storage version
}

func NewCustomResourceDefinition() *ResCustomResourceDefinition {
	return &ResCustomResourceDefinition{
		ApiVersion: "apiextensions.k8s.io/v1beta1",
		Kind:       RESOURCE_CUSTOM_RESOURCE_DEFINITION,
		Metadata: struct {
			Name        string
		}{Name: ""},
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

func (r *ResCustomResourceDefinition) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}



