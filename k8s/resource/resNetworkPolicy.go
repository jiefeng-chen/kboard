package resource

import (
	"errors"
	"gopkg.in/yaml.v2"
)

// k8s >= v1.3
type IResNetworkPolicy interface {
	IResource
	SetMetadataName(string) error
	SetNamespace(string) error
	SetLabels(map[string]string) error
	SetPolicyType([]string) error
	AddPodSelector(map[string]string) error
	AddIngress(*Ingress) error
	AddEgress(*Egress) error
}

const (
	POLICY_TYPE_INGRESS = "Ingress"
	POLICY_TYPE_EGRESS = "Egress"
)


type ResNetworkPolicy struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind string
	Metadata struct{
		Name string
		Namespace string
		Labels map[string]string
		Annotations map[string]string
	}
	Spec *NetPolicySpec
}

type NetPolicySpec struct {
	PodSelector *PodSelector `yaml:"podSelector"`
	PolicyTypes []string `yaml:"policyTypes"` // Ingress、Egress
	Ingress []*Ingress // 入站规则
	Egress []*Egress // 出站规则
}

type PodSelector struct {
	MatchLabels map[string]string `yaml:"matchLabels"`
}

type Ingress struct {
	From struct{
		IpBlock *IpBlock `yaml:"ipBlock"`
		NamespaceSelector *NamespaceSelector `yaml:"namespaceSelector"`
		PodSelector *PodSelector `yaml:"podSelector"`
	}
	Ports []*NetPolicyPort
}


type Egress struct {
	To struct{
		IpBlock *IpBlock `yaml:"ipBlock"`
	}
	Ports []*NetPolicyPort
}

type IpBlock struct {
	Cidr string // 源地址ip地址段
	Except []string // 排除ip地址段
}

type NamespaceSelector struct {
	MatchLabels map[string]string `yaml:"matchLabels"`
}

type NetPolicyPort struct {
	Protocol string // TCP UDP
	Port int
}

func NewResNetworkPolicy() *ResNetworkPolicy {
	return &ResNetworkPolicy{
		ApiVersion: "networking.k8s.io/v1",
		Kind: RESOURCE_NETWORK_POLICY,
		Metadata: struct {
			Name        string
			Namespace   string
			Labels      map[string]string
			Annotations map[string]string
		}{
			Name: "",
			Namespace: "",
			Labels: map[string]string{},
			Annotations: map[string]string{}},
		Spec: &NetPolicySpec{
				PodSelector: &PodSelector{},
				PolicyTypes: []string{},
				Ingress: []*Ingress{},
				Egress: []*Egress{}},
	}
}


func (r *ResNetworkPolicy) SetMetadataName(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	r.Metadata.Name = name
	return nil
}

func (r *ResNetworkPolicy) SetNamespace(ns string) error {
	if ns == "" {
		return errors.New("namespace is empty")
	}
	r.Metadata.Namespace = ns
	return nil
}

func (r *ResNetworkPolicy) SetLabels(labels map[string]string) error {
	if len(labels) <= 0 {
		return errors.New("labels is empty")
	}
	for k, v := range labels {
		if k == "" || v == "" {
			return errors.New("label's key or value is empty")
		}
		if r.Metadata.Labels[k] != "" {
			return errors.New("label [" + k + "] is exist")
		}
		r.Metadata.Labels[k] = v
	}
	return nil
}

func (r *ResNetworkPolicy) SetPolicyType(types []string) error {
	if len(types) <= 0 {
		return errors.New("types is empty")
	}
	for _, v := range types {
		if v != POLICY_TYPE_EGRESS && v != POLICY_TYPE_INGRESS {
			return errors.New("type is invalid")
		}
		r.Spec.PolicyTypes = append(r.Spec.PolicyTypes, v)
	}
	return nil
}

func (r *ResNetworkPolicy) AddPodSelector(labels map[string]string) error {
	if len(labels) <= 0 {
		return errors.New("label is empty")
	}
	for k, v := range labels {
		if k == "" || v == "" {
			return errors.New("label's key or value is empty")
		}
		r.Spec.PodSelector.MatchLabels[k] = v
	}
	return nil
}

func (r *ResNetworkPolicy) AddIngress(ing *Ingress) error {
	if ing == nil {
		return errors.New("ingress is nil")
	}
	if len(ing.Ports) <= 0 {
		return errors.New("ports is empty")
	}
	if ing.From.IpBlock.Cidr == "" {
		return errors.New("ingress's source ip address is empty")
	}

	r.Spec.Ingress = append(r.Spec.Ingress, ing)
	return nil
}

func (r *ResNetworkPolicy) AddEgress(eg *Egress) error {
	if eg == nil {
		return errors.New("egress is nil")
	}
	if len(eg.Ports) <= 0 {
		return errors.New("ports is empty")
	}
	if eg.To.IpBlock.Cidr == "" {
		return errors.New("egress's source ip address is empty")
	}

	r.Spec.Egress = append(r.Spec.Egress, eg)
	return nil
}


func (r *ResNetworkPolicy) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}

