package resource

import (
	"dashboard/core"
	"errors"
	"gopkg.in/yaml.v2"
)

type IResDaemonSet interface {
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

type ResDaemonSet struct {
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
		Template *DaemonSetTemplate
	}
}

func NewResDaemonSet() *ResDaemonSet {
	return &ResDaemonSet{
		ApiVersion: "apps/v1",
		Kind:       RESOURCE_DAEMONSET,
		Metadata: struct {
			Name      string
			Namespace string
			Labels    map[string]string
		}{Name: "", Namespace: "", Labels: map[string]string{}},
	}
}

type DaemonSetTemplate struct {
	Metadata struct {
		Labels map[string]string `yaml:"labels"`
	}
	Spec *DaemonSetTemplateSpec
}

type DaemonSetTemplateSpec struct {
	Tolerations                   []*DaemonSetToleration
	Containers                    []*DaemonSetContainer
	TerminationGracePeriodSeconds string `yaml:"terminationGracePeriodSeconds"`
	Volumes                       []map[string]interface{}
	RestartPolicy                 string            `yaml:"restartPolicy"` // 默认为 Always
	ImagePullSecrets              map[string]string `yaml:"imagePullSecrets"`
	NodeSelector                  map[string]string `yaml:"nodeSelector"`
}

type DaemonSetToleration struct {
	Key               string
	Effect            string
	Value             string
	Operator          string
	TolerationSeconds string `yaml:"tolerationSeconds"`
}

type DaemonSetContainer struct {
	Name            string
	Image           string
	Resources       *DaemonSetResource
	VolumeMounts    []*DaemonSetVolumeMount `yaml:"volumeMounts"`
	Command         []string
	Args            []string
	ImagePullPolicy string `yaml:"imagePullPolicy"` // 默认为Always
	Ports           []*DaemonSetPort
	Env             []*DaemonSetEnv
}

type DaemonSetEnv struct {
	Name  string
	Value string
}

type DaemonSetPort struct {
	Name          string
	ContainerPort int    `yaml:"containerPort"` // 容器端口
	HostPort      int    `yaml:"hostPort"`      // 主机端口
	Protocol      string // 端口协议，默认为TCP
}

type DaemonSetResource struct {
	Limits struct {
		Cpu    string
		Memory string
	}
	Requests struct {
		Cpu    string
		Memory string
	}
}

type DaemonSetVolumeMount struct {
	Name      string
	MountPath string `yaml:"mountPath"`
	ReadOnly  bool   `yaml:"readOnly"` // 默认为读写模式
}

func (r *ResDaemonSet) SetMetaDataName(name string) error {
	if name == "" {
		return errors.New("metadata name is empty")
	}
	// 设置 .metadata.name
	r.Metadata.Name = name
	// 设置 .spec.selector.matchLabels.name
	r.Spec.Selector.MatchLabels["name"] = name
	// 设置 .spec.template.metadata.labels.name
	r.Spec.Template.Metadata.Labels["name"] = name
	return nil
}

func (r *ResDaemonSet) SetNamespace(ns string) error {
	if ns == "" {
		return errors.New("metadata namespace is empty")
	}
	r.Metadata.Namespace = ns
	return nil
}

func (r *ResDaemonSet) GetNamespace() string {
	return r.Metadata.Namespace
}

func (r *ResDaemonSet) SetMatchLabels(labels []map[string]string) error {
	if len(labels) > 0 {
		for _, v := range labels {
			for key, val := range v {
				if key == "" {
					return errors.New("matchLabels key is empty")
				}

				r.Spec.Selector.MatchLabels[key] = val
			}
		}
	}
	return nil
}

func (r *ResDaemonSet) SetTolerations(tolers []map[string]string) error {
	if len(tolers) > 0 {
		for _, v := range tolers {
			if v["key"] == "" {
				return errors.New("toleration key is empty")
			}
			toler := &DaemonSetToleration{
				Key:               v["key"],
				Operator:          v["operator"],
				Effect:            v["effect"],
				Value:             v["value"],
				TolerationSeconds: v["tolerationSeconds"],
			}
			r.Spec.Template.Spec.Tolerations = append(r.Spec.Template.Spec.Tolerations, toler)
		}
	}
	return nil
}

func (r *ResDaemonSet) SetContainers(containers []map[string]interface{}) error {
	if len(containers) > 0 {
		for _, v := range containers {
			if v["name"] == "" ||
				v["image"] == "" ||
				v["limit_memory"] == "" ||
				v["limit_cpu"] == "" ||
				v["request_cpu"] == "" ||
				v["request_memory"] == "" {
				errors.New("container parameters is empty")
			}
			container := &DaemonSetContainer{
				Name:  core.GetMapValueByKey(v, "name"),
				Image: core.GetMapValueByKey(v, "image"),
				Resources: &DaemonSetResource{
					Limits: struct {
						Cpu    string
						Memory string
					}{Cpu: core.GetMapValueByKey(v, "limit_cpu"), Memory: core.GetMapValueByKey(v, "limit_memory")},
					Requests: struct {
						Cpu    string
						Memory string
					}{Cpu: core.GetMapValueByKey(v, "request_cpu"), Memory: core.GetMapValueByKey(v, "request_memory")},
				},
				ImagePullPolicy: "IfNotPresent", // 如果本地有该镜像，则使用本地镜像
			}
			// 1. 处理command和args
			if cmds := v["command"].([]string); len(cmds) > 0 {
				for _, cmd := range cmds {
					if cmd == "" {
						continue
					}
					container.Command = append(container.Command, cmd)
				}
			}

			if args := v["args"].([]string); len(args) > 0 {
				for _, arg := range args {
					if arg == "" {
						continue
					}
					container.Args = append(container.Args, arg)
				}
			}
			// 2. 处理volumeMounts
			if volMounts := v["volumeMounts"].([]map[string]string); len(volMounts) > 0 {
				for _, volMount := range volMounts {
					if volMount["name"] == "" {
						return errors.New(".spec.template.spec.container.volumeMounts.name is empty")
					}
					container.VolumeMounts = append(container.VolumeMounts, &DaemonSetVolumeMount{
						Name:      volMount["name"],
						MountPath: volMount["mountPath"],
					})
				}
			}
			// 3. 处理ports
			if ports := v["ports"].([]map[string]string); len(ports) > 0 {
				for _, port := range ports {
					if port["name"] == "" || port["containerPort"] == "" || port["hostPort"] == "" || port["protocol"] == "" {
						return errors.New(".spec.container.ports's parameters error")
					}
					container.Ports = append(container.Ports, &DaemonSetPort{
						Name:          core.ToString(port["name"]),
						ContainerPort: core.ToInt(port["containerPort"]),
						HostPort:      core.ToInt(port["hostPort"]),
						Protocol:      core.ToString(port["protocol"]),
					})
				}
			}
			// 4. 处理env
			if envs := v["env"].([]map[string]string); len(envs) > 0 {
				for _, env := range envs {
					if env["name"] == "" {
						return errors.New(".spec.containers.env's name is empty")
					}
					container.Env = append(container.Env, &DaemonSetEnv{
						Name:  env["name"],
						Value: env["val"],
					})
				}
			}
			r.Spec.Template.Spec.Containers = append(r.Spec.Template.Spec.Containers, container)
		}
	}
	return nil
}

func (r *ResDaemonSet) SetTerminationGracePeriodSeconds(second string) error {
	if second == "" {
		return errors.New("termination grace period seconds is empty")
	}
	r.Spec.Template.Spec.TerminationGracePeriodSeconds = second
	return nil
}

type VolumeHostPath struct {
	Name     string
	HostPath struct {
		Path string
	} `yaml:"hostPath"`
}

type VolumeConfigMap struct {
	Name      string
	ConfigMap struct {
		Name string
	} `yaml:"configMap"`
}

type VolumeSecret struct {
}

type VolumeEmptyDir struct {
	Name     string
	EmptyDir struct{}
}

type VolumePersistentVolumeClaim struct {
	Name                  string
	PersistentVolumeClaim struct {
		ClaimName string `yaml:"claimName"`
	} `yaml:"persistentVolumeClaim"`
}

func (r *ResDaemonSet) SetVolumes(volumes []map[string]string, volumeType string) error {
	if len(volumes) > 0 {
		switch volumeType {
		case "hostPath":
		case "secret":
		case "configMap":
		case "emptyDir":
		case "persistentVolumeClaim":
		default:
			return errors.New("volume type[" + volumeType + "] is not supported")
		}
	}
	return nil
}

func (r *ResDaemonSet) SetRestartPolicy(rPolicy string) error {
	if rPolicy == "" {
		return errors.New("restart policy is empty")
	}
	r.Spec.Template.Spec.RestartPolicy = rPolicy
	return nil
}

func (r *ResDaemonSet) SetNodeSelector(selectors []map[string]string) error {
	if len(selectors) > 0 {
		for _, v := range selectors {
			if v["key"] == "" {
				return errors.New("node selector's key is empty")
			}
			r.Spec.Template.Spec.NodeSelector[v["key"]] = v["val"]
		}
	}
	return nil
}

func (r *ResDaemonSet) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}
