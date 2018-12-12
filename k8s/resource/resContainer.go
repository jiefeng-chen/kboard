package resource

import "errors"

type IContainer interface {
	AppendPort(*Port) error
	AppendVolumeMount(map[string]interface{}) error
	AppendEnv(*Env) error
}


// 容器结构体
type Container struct {
	Name string
	Image string
	ImagePullPolicy string	`yaml:"imagePullPolicy"` // [Always | Never | IfNotPresent]
	Command string
	Args string
	WorkingDir string `yaml:"workingDir"`	// 当前工作目录
	VolumeMounts []map[string]interface{} `yaml:"volumeMounts"`	// 挂载卷
	Resources *Resource
	Env []*Env
	Ports []*Port	// 端口号
	LivenessProbe *LivenessProbe  `yaml:"livenessProbe"`
	SecurityContext struct{
		Privileged bool // true-容器运行在特权模式
	} `yaml:"securityContext"`
}

func NewContainer(name string, image string) *Container {
	return &Container{
		Name: name,
		Image: image,
		Resources:NewResource(),
		LivenessProbe:NewLivenessProbe(),
		SecurityContext: struct{ Privileged bool }{Privileged: false},
		Env: []*Env{},
		Ports: []*Port{},
	}
}

func (r *Container) AppendPort(port *Port) error {
	if port == nil {
		return errors.New("port is nil")
	}
	r.Ports = append(r.Ports, port)
	return nil
}

func (r *Container) AppendVolumeMount(vol map[string]interface{}) error {
	if vol == nil {
		return errors.New("volume is nil")
	}
	r.VolumeMounts = append(r.VolumeMounts, vol)
	return nil
}

type Env struct {
	Name string
	ValueFrom *ValueFrom `yaml:"valueFrom"`
}

type ValueFrom struct {
	FieldRef *FieldRef `yaml:"fieldRef"`
	ResourceFieldRef *ResourceFieldRef `yaml:"resourceFieldRef"`
}

type FieldRef struct {
	FieldPath string `yaml:"fieldPath"`
}

type ResourceFieldRef struct {
	ContainerName string `yaml:"containerName"`
	Resource string
}

func NewEnv() *Env {
	return &Env{
		Name: "",
		ValueFrom: &ValueFrom{
			FieldRef: &FieldRef{
				FieldPath: "",
			},
			ResourceFieldRef: &ResourceFieldRef{
				ContainerName: "",
				Resource: "",
			},
		}}
}

func (r *Container) AppendEnv(env *Env) error {
	if env == nil {
		return errors.New("env is nil")
	}
	r.Env = append(r.Env, env)
	return nil
}

func NewLivenessProbe() *LivenessProbe {
	return &LivenessProbe{
		Exec: struct{ Command []string }{Command: []string{}},
		HttpGet: &HttpGet{
			Path: "",
			Port: "",
			Host: "",
			Scheme: "",
			HttpHeaders: []map[string]string{},
		},
		TcpSocket: struct{ Port string }{Port: ""},
	}
}

type LivenessProbe struct {
	Exec struct{
		Command []string
	}
	HttpGet *HttpGet `yaml:"httpGet"`
	TcpSocket struct{
		Port string
	} `yaml:"tcpSocket"`
	InitialDelaySeconds int `yaml:"initialDelaySeconds"`
	TimeoutSeconds int `yaml:"timeoutSeconds"`
	PeriodSeconds int `yaml:"periodSeconds"`
	SuccessThreshold int `yaml:"successThreshold"`
	FailureThreshold int `yaml:"failureThreshold"`
}

type HttpGet struct {
	Path string
	Port string
	Host string
	Scheme string
	HttpHeaders []map[string]string `yaml:"httpHeaders"`
}


type Secret struct {
	SecretName string `yaml:"secretName"`
	Items []map[string]string // [key:string, path:string]
}

func NewResource() *Resource {
	return &Resource{
		Limits: &Limits{Cpu: "", Memory: ""},
		Requests: &Request{Cpu: "", Memory: ""},
	}
}

type Resource struct{
	Limits *Limits
	Requests *Request
}

type Port struct {
	Name string
	ContainerPort int `yaml:"containerPort"`
	HostPort int `yaml:"hostPort"`
	Protocol string
}

func NewPort(name string) *Port {
	return &Port{
		Name: name,
		ContainerPort: 0,
		HostPort: 0,
		Protocol: "",
	}
}


