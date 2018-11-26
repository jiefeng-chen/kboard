package resource

type IResDeployment interface {
	IResource
	SetMetaDataName(string) error
	SetNamespace(string) error
	GetNamespace() string
	SetMatchLabels([]map[string]string) error
	SetTolerations([]map[string]string) error
	SetContainers([]map[string]interface{}) error
	SetTerminationGracePeriodSeconds(string) error
	SetVolumes([]map[string]string, string) error
	SetRestartPolicy(string) error
	SetNodeSelector([]map[string]string) error
}

type ResDeployment struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name      string
		Namespace string
		Labels    map[string]string
	}
	Spec struct {
		Selector struct {
			MatchLabels map[string]string `yaml:"matchLabels"`
		}
		Template *DeploymentTemplate
		Replicas string
	}
}

func NewResDeployment() *ResDeployment {
	return &ResDeployment{
		ApiVersion: "apps/v1",
		Kind:       RESOURCE_DEPLOYMENT,
		Metadata: struct {
			Name      string
			Namespace string
			Labels    map[string]string
		}{Name: "", Namespace: "", Labels: map[string]string{}},
	}
}

type DeploymentTemplate struct {
	Metadata struct {
		Labels map[string]string
	}
	Spec struct {
		Containers struct {
		}
	}
}
