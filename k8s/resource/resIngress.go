package resource

import (
	"gopkg.in/yaml.v2"
	"kboard/utils"
	"kboard/exception"
)

type IResIngress interface {
	IResource
	SetMetadataName(string) error
	GetMetadataName() string
	SetNamespace(string) error
	SetRules(string, []map[string]string) error
	SetAnnotations(map[string]string) error
	SetLabels([]map[string]string) error
	SetTls([]string, string) error
}

type ResIngress struct {
	Kind       string
	ApiVersion string `yaml:"apiVersion"`
	MetaData   struct {
		Name        string
		Namespace   string
		Annotations map[string]string
		Labels      map[string]string
	}
	Spec struct {
		Rules []*IngressRule
		Tls []Tls
	}
}

type Tls struct {
	Hosts []string
	SecretName string `yaml:"secretName"`
}

type IngressRule struct {
	Host string
	Http struct {
		Paths []IngressPath
	}
}

type IngressPath struct {
	Path    string
	Backend IngressBackend
}

type IngressBackend struct {
	ServiceName string `yaml:"serviceName"`
	ServicePort int    `yaml:"servicePort"`
}

const (
	ANNOTATIONS_INGRESS_CLASS = "kubernetes.io/ingress.class"
	ANNOTATIONS_WHITELIST_X_FORWARDED_FOR = "ingress.kubernetes.io/whitelist-x-forwarded-for" // 是否开启ip白名单
	ANNOTATIONS_WHITELIST_SOURCE_RANGE = "traefik.ingress.kubernetes.io/whitelist-source-range" // ip白名单列表
)

func NewIngress() *ResIngress {
	return &ResIngress{
		Kind:       RESOURCE_INGRESS,
		ApiVersion: "extensions/v1beta1",
		MetaData: struct {
			Name        string
			Namespace   string
			Annotations map[string]string
			Labels      map[string]string
		}{Name: "", Namespace: "", Annotations: map[string]string{}, Labels: map[string]string{}},
	}
}

func (r *ResIngress) SetAnnotations(annot map[string]string) error {
	if len(annot) <= 0 {
		return exception.NewError("Annotations is empty")
	}
	for k, v := range annot {
		r.MetaData.Annotations[k] = v
	}
	return nil
}

func (r *ResIngress) SetMetadataName(name string) error {
	if name == "" {
		return exception.NewError("name is empty")
	}
	r.MetaData.Name = name
	return nil
}

func (r *ResIngress) SetNamespace(ns string) error {
	if ns == "" {
		return exception.NewError("namespace is empty")
	}
	r.MetaData.Namespace = ns
	return nil
}

func (r *ResIngress) GetMetadataName() string {
	return r.MetaData.Name
}

func (r *ResIngress) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}

func (r *ResIngress) SetRules(host string, rules []map[string]string) error {
	if len(rules) > 0 {
		rule := new(IngressRule)
		rule.Host = host
		for _, v := range rules {
			path := new(IngressPath)
			if v["serviceName"] == "" || v["servicePort"] == "" {
				return exception.NewError("服务名称或服务端口为空")
			}
			if v["path"] != "" {
				// 这里允许访问路径为空，因为可以直接通过域名访问
				path.Path = v["path"]
			}
			path.Backend = IngressBackend{
				ServiceName: v["serviceName"],
				ServicePort: utils.ToInt(v["servicePort"]),
			}
			rule.Http.Paths = append(rule.Http.Paths, *path)
		}
		r.Spec.Rules = append(r.Spec.Rules, rule)
	}
	return nil
}

func (r *ResIngress) SetLabels(labels []map[string]string) error {
	if len(labels) > 0 {
		for _, v := range labels {
			for key, val := range v {
				r.MetaData.Labels[key] = val
			}
		}
	}
	return nil
}

func (r *ResIngress) SetTls(hosts []string, secretName string) error {
	if len(hosts) < 0 {
		return exception.NewError("缺少域名")
	}
	if secretName == "" {
		return exception.NewError("缺少tls密钥对")
	}
	tls := Tls{
		Hosts: []string{},
		SecretName: secretName,
	}
	for _, host := range hosts {
		tls.Hosts = append(tls.Hosts, host)
	}
	r.Spec.Tls = append(r.Spec.Tls, tls)
	return nil
}