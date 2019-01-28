package resource

type IResRoleBinding interface {
	IResource
	SetMetadataName(string) error
	SetNamespace(string) error
	GetNamespace() string
}

type ResRoleBinding struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name      string
		Namespace string
	}
	Subjects []RoleBindingSubject
	RoleRef RoleRef
}

type RoleBindingSubject struct {
	Kind string
	Name string
	ApiGroup string `yaml:"apiGroup"`
}

type RoleRef struct {
	Kind string // this must be Role or ClusterRole
	Name string // this must match the name of the Role or ClusterRole you wish to bind to
	ApiGroup string `yaml:"apiGroup"`
}

func NewResRoleBinding() *ResRoleBinding {
	return &ResRoleBinding{
		ApiVersion: "rbac.authorization.k8s.io/v1",
		Kind:       RESOURCE_ROLE_BINDING,
		Metadata: struct {
			Name      string
			Namespace string
		}{Name: "", Namespace: ""},
		Subjects: []RoleBindingSubject{},
		RoleRef: RoleRef{},
	}
}




