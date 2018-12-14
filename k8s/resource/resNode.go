package resource

import (
	"time"
	"kboard/exception"
	"gopkg.in/yaml.v2"
)

type IResNode interface {
	IResource
	SetMetadataName(string) error
	SetNamespace(string) error
}

// pod结构体
type ResNode struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind string
	Metadata struct{
		Name string
		Namespace string
		Labels []map[string]string
		Annotations []map[string]string
	}
	Spec *NodeSpec
}

type NodeSpec struct {
	ConfigSource *NodeConfigSource `yaml:"configSource"`
	ExternalID string `yaml:"externalID"`
	PodCIDR string `yaml:"podCIDR"`
	ProviderID string `yaml:"providerID"`
	Taints []*Taint
	Unschedulable bool // 将unschedulable设置为true实现隔离，恢复为false
}

type NodeConfigSource struct {
	ConfigMap *ConfigMapNodeConfigSource
}

type Taint struct {
	Effect string
	Key string
	TimeAdded *time.Time `yaml:"timeAdded"`
	Value string
}

type ConfigMapNodeConfigSource struct {
	KubeletConfigKey string `yaml:"kubeletConfigKey"`
	Name string
	Namespace string
	ResourceVersion string `yaml:"resourceVersion"`
	Uid string
}

func NewResNode(name string) *ResNode {
	return &ResNode{
		ApiVersion: "v1",
		Kind: RESOURCE_NODE,
		Metadata: struct {
			Name        string
			Namespace   string
			Labels      []map[string]string
			Annotations []map[string]string
		}{Name: name, Namespace: "", Labels: []map[string]string{}, Annotations: []map[string]string{}},
		Spec: &NodeSpec{
			ConfigSource: &NodeConfigSource{
				ConfigMap: &ConfigMapNodeConfigSource{

				},
			},
			ExternalID: "",
			ProviderID: "",
			PodCIDR: "",
			Unschedulable: false,
			Taints: []*Taint{},
		},
	}
}


func (r *ResNode) SetMetadataName(name string) error {
	if name == "" {
		return exception.NewError("name is empty")
	}
	r.Metadata.Name = name
	return nil
}

func (r *ResNode) SetNamespace(ns string) error {
	if ns == "" {
		return exception.NewError("namespace is empty")
	}
	r.Metadata.Namespace = ns
	return nil
}


func (r *ResNode) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}

