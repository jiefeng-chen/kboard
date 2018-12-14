package resource

import (
	"errors"
	"gopkg.in/yaml.v2"
	"strings"
	"kboard/utils"
	"kboard/exception"
)

type IResStorageClass interface {
	IResource
	GetProvisioner() string
	SetMetaDataName(string) error
	SetProvisioner(string) error
	SetReclaimPolicy(string) error
	SetParameters(interface{}) error
	GetReclaimPolicy() string
	SetAllowVolumeExpansion(bool) error
	GetParameters() interface{}
	GetAllowVolumeExpansion() bool
}

type ResStorageClass struct {
	Kind       string `yaml:"kind"`
	ApiVersion string `yaml:"apiVersion"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	}
	Provisioner          string
	ReclaimPolicy        string `yaml:"reclaimPolicy"`
	AllowVolumeExpansion bool   `yaml:"allowVolumeExpansion"`
	Parameters           interface{}
}

func NewStorageClass() *ResStorageClass {
	return &ResStorageClass{
		Kind:                 RESOURCE_STORAGE_CLASS,
		ApiVersion:           "storage.k8s.io/v1",
		AllowVolumeExpansion: true,
	}
}

func (r *ResStorageClass) GetReclaimPolicy() string {
	return r.ReclaimPolicy
}

func (r *ResStorageClass) GetProvisioner() string {
	return r.Provisioner
}

func (r *ResStorageClass) SetProvisioner(provisioner string) error {
	r.Provisioner = provisioner
	return nil
}

func (r *ResStorageClass) SetMetaDataName(name string) error {
	r.Metadata.Name = name
	return nil
}

func (r *ResStorageClass) SetReclaimPolicy(rp string) error {
	r.ReclaimPolicy = rp
	return nil
}

func (r *ResStorageClass) SetParameters(params interface{}) error {
	r.Parameters = params
	return nil
}

func (r *ResStorageClass) SetAllowVolumeExpansion(ave bool) error {
	r.AllowVolumeExpansion = ave
	return nil
}

func (r *ResStorageClass) GetAllowVolumeExpansion() bool {
	return r.AllowVolumeExpansion
}

func (r *ResStorageClass) GetParameters() interface{} {
	switch r.GetProvisioner() {
	case "kubernetes.io/rbd":
		var cephrbd *CephRbd
		return cephrbd
	}
	return ""
}

func (r *ResStorageClass) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}

// ceph rbd
type cephRbd interface {
	SetMonitors(string) error
	SetAdminId(string) error
	SetAdminSecretName(string) error
	SetAdminSecretNamespace(string) error
	SetPool(string) error
	SetUserId(string) error
	SetUserSecretName(string) error
	SetFsType(string) error
	SetData([]map[string]string) error
	SetImageFormat(string) error
	SetImageFeatures(string) error
}

type CephRbd struct {
	Monitors             string `yaml:"monitors"`
	AdminId              string `yaml:"adminId"`
	AdminSecretName      string `yaml:"adminSecretName"`
	AdminSecretNamespace string `yaml:"adminSecretNamespace"`
	Pool                 string `yaml:"pool"`
	UserId               string `yaml:"userId"`
	UserSecretName       string `yaml:"userSecretName"`
	FsType               string `yaml:"fsType"`
	ImageFormat          string `yaml:"imageFormat"`
	ImageFeatures        string `yaml:"imageFeatures"`
}

func NewCephRbd() *CephRbd {
	return &CephRbd{
		Monitors:             "",
		AdminId:              "admin",
		AdminSecretName:      "admin",
		AdminSecretNamespace: "",
		Pool:                 "",
		UserId:               "",
		UserSecretName:       "",
		FsType:               "xfs",
		ImageFormat:          "2",
		ImageFeatures:        "layering",
	}
}

func (r *CephRbd) SetMonitors(monitors string) error {
	// 检查ip地址
	monis := strings.Split(monitors, ",")
	if len(monis) <= 0 {
		return exception.NewError("monitor is empty")
	}
	for _, v := range monis {
		// 检查ip:port格式
		url := strings.Split(v, ":")
		if !utils.IsIP(url[0]) {
			return exception.NewError("ip is empty")
		}
		// 端口检查
		if len(url) <= 1 || url[1] == "" || utils.ToInt(url[1]) <= 0 {
			return exception.NewError("port is empty")
		}
	}
	r.Monitors = monitors
	return nil
}

func (r *CephRbd) SetAdminId(adminId string) error {
	r.AdminId = adminId
	return nil
}

func (r *CephRbd) SetAdminSecretName(adminSN string) error {
	r.AdminSecretName = adminSN
	return nil
}

func (r *CephRbd) SetAdminSecretNamespace(adminSNs string) error {
	r.AdminSecretNamespace = adminSNs
	return nil
}

func (r *CephRbd) SetPool(pool string) error {
	r.Pool = pool
	return nil
}

func (r *CephRbd) SetUserId(uid string) error {
	r.UserId = uid
	return nil
}

func (r *CephRbd) SetUserSecretName(userSN string) error {
	r.UserSecretName = userSN
	return nil
}

func (r *CephRbd) SetFsType(fst string) error {
	r.FsType = fst
	return nil
}

func (r *CephRbd) SetData(data []map[string]string) error {
	if len(data) <= 0 {
		return errors.New("no data to set")
	}
	for _, v := range data {
		if v["val"] == "" {
			return errors.New(v["key"] + " is empty")
		}
		switch v["key"] {
		case "monitors":
			monitors := strings.Split(v["val"], ",")
			if err := r.SetMonitors(strings.Join(monitors, ",")); err != nil {
				return exception.NewError(v["key"] + " format error")
			}
		case "adminId":
			r.SetAdminId(v["val"])
		case "adminSecretName":
			r.SetAdminSecretName(v["val"])
		case "adminSecretNamespace":
			r.SetAdminSecretNamespace(v["val"])
		case "pool":
			r.SetPool(v["val"])
		case "userId":
			r.SetUserId(v["val"])
		case "userSecretName":
			r.SetUserSecretName(v["val"])
		case "fsType":
			r.SetFsType(v["val"])
		case "imageFormat":
			r.SetImageFormat(v["val"])
		case "imageFeatures":
			r.SetImageFeatures(v["val"])
		}
	}
	return nil
}

func (r *CephRbd) SetImageFormat(imgFmt string) error {
	r.ImageFormat = imgFmt
	return nil
}

func (r *CephRbd) SetImageFeatures(imgFeat string) error {
	r.ImageFeatures = imgFeat
	return nil
}
