package resource

import (
	"kboard/exception"
	"time"

	"gopkg.in/yaml.v2"
)

type IResNode interface {
	IResource
	SetMetadataName(string) error
	SetNamespace(string) error
	SetLabels(map[string]string) error
	SetUnschedulable(bool) error
	GetExternalID() string
}

type ResNode struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name        string
		Namespace   string
		Labels      map[string]string
		Annotations map[string]string
	}
	Spec   *NodeSpec
	Status *NodeStatus
}

type NodeSpec struct {
	ConfigSource  *NodeConfigSource `yaml:"configSource"`
	ExternalID    string            `yaml:"externalID"` // 节点ip
	PodCIDR       string            `yaml:"podCIDR"`
	ProviderID    string            `yaml:"providerID"`
	Taints        []*Taint
	Unschedulable bool // 将unschedulable设置为true实现隔离，恢复为false
}

type NodeConfigSource struct {
	ConfigMap *ConfigMapNodeConfigSource
}

type Taint struct {
	Effect    string
	Key       string
	TimeAdded *time.Time `yaml:"timeAdded"`
	Value     string
}

type ConfigMapNodeConfigSource struct {
	KubeletConfigKey string `yaml:"kubeletConfigKey"`
	Name             string
	Namespace        string
	ResourceVersion  string `yaml:"resourceVersion"`
	Uid              string
}

type NodeStatus struct {
	Capacity        *Capacity
	Allocatable     *Allocatable
	NodeInfo        *NodeInfo `yaml:"nodeInfo"`
	Addresses       []*Address
	Images          []*Image
	DaemonEndpoints *DaemonEndpoint `yaml:"daemonEndpoints"`
}

// 设置资源容量
type Capacity struct {
	Cpu              string
	EphemeralStorage string `yaml:"ephemeral-storage"` // 管理短暂存储
	Hugepages_1G     string `yaml:"hugepages-1Gi"`     // 大页
	Hugepages_2M     string `yaml:"Hugepages_2Mi"`
	Memory           string // 内存
	Pods             string
}

// 可分配数据
type Allocatable struct {
	Cpu              string
	EphemeralStorage string `yaml:"ephemeral-storage"` // 管理短暂存储
	Hugepages_1G     string `yaml:"hugepages-1Gi"`     // 大页
	Hugepages_2M     string `yaml:"Hugepages_2Mi"`
	Memory           string // 内存
	Pods             string
}

type NodeInfo struct {
	MachineID               string `yaml:"machineID"`               // 机器id
	SystemUUID              string `yaml:"systemUUID"`              // 系统uuid
	BootID                  string `yaml:"bootID"`                  // 启动id
	KernelVersion           string `yaml:"kernelVersion"`           // 内核版本
	OsImage                 string `yaml:"osImage"`                 // 操作系统镜像版本
	ContainerRuntimeVersion string `yaml:"containerRuntimeVersion"` // 容器运行时版本
	KubeletVersion          string `yaml:"kubeletVersion"`
	KubeProxyVersion        string `yaml:"kubeProxyVersion"`
	OperatingSystem         string `yaml:"operatingSystem"` // 操作系统
	Architecture            string `yaml:"architecture"`    // 架构
}

type Address struct {
	Type    string // InternalIP、Hostname
	Address string // ip地址
}

type Image struct {
	Names     []string // 镜像名称
	SizeBytes int      // 镜像大小
}

type DaemonEndpoint struct {
	KubeletEndpoint struct {
		Port int
	}
}

type Conditions struct {
}

func NewResNode(name string) *ResNode {
	return &ResNode{
		ApiVersion: "v1",
		Kind:       RESOURCE_NODE,
		Metadata: struct {
			Name        string
			Namespace   string
			Labels      map[string]string
			Annotations map[string]string
		}{Name: name, Namespace: "", Labels: map[string]string{}, Annotations: map[string]string{}},
		Spec: &NodeSpec{
			ConfigSource: &NodeConfigSource{
				ConfigMap: &ConfigMapNodeConfigSource{},
			},
			ExternalID:    "",
			ProviderID:    "",
			PodCIDR:       "",
			Unschedulable: false,
			Taints:        []*Taint{},
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

func (r *ResNode) SetLabels(labels map[string]string) error {
	if len(labels) <= 0 {
		return exception.NewError("labels is empty")
	}
	for k, v := range labels {
		if k == "" || v == "" {
			return exception.NewError("labels key or val is empty")
		}
		r.Metadata.Labels[k] = v
	}
	return nil
}

// 设置节点隔离和恢复状态 true-隔离 false-恢复
func (r *ResNode) SetUnschedulable(stat bool) error {
	r.Spec.Unschedulable = stat
	return nil
}

func (r *ResNode) GetExternalID() string {
	return r.Spec.ExternalID
}

func (r *ResNode) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}
