package resource

import "kboard/internal"

type IResRole interface {
	internal.IResource
	SetMetadataName(string) error
	SetNamespace(string) error
	GetNamespace() string
}

type ResRole struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name      string
		Namespace string
	}
	Rules []RoleRule
}

type RoleRule struct {
	ApiGroups []string `yaml:"apiGroups"`
	Resources []string
	Verbs     []string
}

func NewResRole() *ResRole {
	return &ResRole{
		ApiVersion: "rbac.authorization.k8s.io/v1",
		Kind:       internal.RESOURCE_ROLE,
		Metadata: struct {
			Name      string
			Namespace string
		}{Name: "", Namespace: ""},
		Rules: []RoleRule{},
	}
}
