package resource

import (
	"errors"
)

type IContainer interface {
	AppendPort(*Port) error
	AppendVolumeMount(map[string]interface{}) error
	AppendEnv(*Env) error
	SetResource(*Resource) error
	SetImage(string) error
	SetName(string) error
	SetLivenessProbe(LivenessProbe) error
	SetReadinessProbe(ReadinessProbe) error
}


// 容器结构体
type Container struct {
	Name string
	Image string
	ImagePullPolicy string	`yaml:"imagePullPolicy"` // [Always | Never | IfNotPresent]
	Command []string
	Args []string
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
		SecurityContext: struct{ Privileged bool }{
			Privileged: false},
		Env: []*Env{},
		Ports: []*Port{},
		Args: []string{},
		Command: []string{},
		ImagePullPolicy: "",
		WorkingDir: "",
		VolumeMounts: []map[string]interface{}{},
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

func (r *Container) SetResource(res *Resource) error {
	if res == nil {
		return errors.New("resource is nil")
	}
	if res.Requests.Cpu == "" || res.Requests.Memory == "" {
		return errors.New("request cpu or memory is empty")
	}
	r.Resources = res
	return nil
}

func (r *Container) SetImage(img string) error {
	if img == "" {
		return errors.New("image is empty")
	}
	r.Image = img
	return nil
}


func (r *Container) SetName(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	r.Name = name
	return nil
}

func (r *Container) SetLivenessProbe(liveness LivenessProbe) error {
	if liveness.HttpGet != nil {
		// http get 方式检查 port、path
		if liveness.Path == "" || liveness.Port == 0 {
			return errors.New("liveness probe use httpget way, need path and port")
		}
	}else if liveness.TcpSocket.Port == 0 {
		return errors.New("tcp socket port way need port")
	}

	if len(liveness.Exec.Command) <= 0 {
		return errors.New("command is empty")
	}
	if liveness.FailureThreshold <= 0 {
		return errors.New("failure threshold is invalid")
	}
	if liveness.InitialDelaySeconds <= 0 {
		return errors.New("initial delay second is invalid")
	}
	if liveness.PeriodSeconds <= 0 {
		return errors.New("period second is invalid")
	}
	if liveness.SuccessThreshold <= 0 {
		return errors.New("success second is invalid")
	}
	if liveness.TimeoutSeconds <= 0 {
		return errors.New("timeout second is invalid")
	}
	return nil
}

func (r *Container) SetReadinessProbe(readiness ReadinessProbe) error {
	if readiness.HttpGet != nil {
		// http get 方式检查 port、path
		if readiness.Path == "" || readiness.Port == 0 {
			return errors.New("liveness probe use httpget way, need path and port")
		}
	}else if readiness.TcpSocket.Port == 0 {
		return errors.New("tcp socket port way need port")
	}

	if readiness.FailureThreshold <= 0 {
		return errors.New("failure threshold is invalid")
	}
	if readiness.InitialDelaySeconds <= 0 {
		return errors.New("initial delay second is invalid")
	}
	if readiness.PeriodSeconds <= 0 {
		return errors.New("period second is invalid")
	}
	if readiness.SuccessThreshold <= 0 {
		return errors.New("success second is invalid")
	}
	if readiness.TimeoutSeconds <= 0 {
		return errors.New("timeout second is invalid")
	}
	return nil
}

func NewLivenessProbe() *LivenessProbe {
	return &LivenessProbe{
	}
}

type ProbeAction struct {
	Exec *ExecAction `yaml:"exec"`
	HttpGet *HttpGetAction `yaml:"httpGet"`
	TcpSocket *TcpSocketAction `yaml:"tcpSocket"`
}

type TcpSocketAction struct {
	Port int
}

type ExecAction struct {
	Command []string
}


type HttpGetAction struct {
	Path string
	Port string
	Host string
	Scheme string
	HttpHeaders []map[string]string `yaml:"httpHeaders"`
}

type LivenessProbe struct {
	ProbeAction `yaml:",inline"`
	InitialDelaySeconds int `yaml:"initialDelaySeconds"`
	TimeoutSeconds int `yaml:"timeoutSeconds"`
	PeriodSeconds int `yaml:"periodSeconds"`
	SuccessThreshold int `yaml:"successThreshold"`
	FailureThreshold int `yaml:"failureThreshold"`
	Path string
	Port int
}

type ReadinessProbe struct {
	ProbeAction
	Path string
	Port int
	InitialDelaySeconds int `yaml:"initialDelaySeconds"`
	TimeoutSeconds int `yaml:"timeoutSeconds"`
	PeriodSeconds int `yaml:"periodSeconds"`
	SuccessThreshold int `yaml:"successThreshold"`
	FailureThreshold int `yaml:"failureThreshold"`
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
	Protocol string // 仅支持 TCP UDP
}

func NewPort(name string) *Port {
	return &Port{
		Name: name,
		ContainerPort: 0,
		HostPort: 0,
		Protocol: "",
	}
}


