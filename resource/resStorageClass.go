package resource

import (
	"dashboard/core"
	"errors"
	"gopkg.in/yaml.v2"
	"strings"
)

type IResStorageClass interface {
	IResource
	GetProvisioner() string
	SetMetaDataName(string) bool
	SetProvisioner(string) bool
	SetReclaimPolicy(string) bool
	SetParameters(interface{}) bool
	GetReclaimPolicy() string
	SetAllowVolumeExpansion(bool) bool
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

func (r *ResStorageClass) SetProvisioner(provisioner string) bool {
	r.Provisioner = provisioner
	return true
}

func (r *ResStorageClass) SetMetaDataName(name string) bool {
	r.Metadata.Name = name
	return true
}

func (r *ResStorageClass) SetReclaimPolicy(rp string) bool {
	r.ReclaimPolicy = rp
	return true
}

func (r *ResStorageClass) SetParameters(params interface{}) bool {
	r.Parameters = params
	return true
}

func (r *ResStorageClass) SetAllowVolumeExpansion(ave bool) bool {
	r.AllowVolumeExpansion = ave
	return true
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
	SetMonitors(string) bool
	SetAdminId(string) bool
	SetAdminSecretName(string) bool
	SetAdminSecretNamespace(string) bool
	SetPool(string) bool
	SetUserId(string) bool
	SetUserSecretName(string) bool
	SetFsType(string) bool
	SetData([]map[string]string) error
	SetImageFormat(string) bool
	SetImageFeatures(string) bool
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

func (r *CephRbd) SetMonitors(monitors string) bool {
	// 检查ip地址
	monis := strings.Split(monitors, ",")
	if len(monis) <= 0 {
		return false
	}
	for _, v := range monis {
		// 检查ip:port格式
		url := strings.Split(v, ":")
		if !core.IsIP(url[0]) {
			return false
		}
		// 端口检查
		if len(url) <= 1 || url[1] == "" || core.ToInt(url[1]) <= 0 {
			return false
		}
	}
	r.Monitors = monitors
	return true
}

func (r *CephRbd) SetAdminId(adminId string) bool {
	r.AdminId = adminId
	return true
}

func (r *CephRbd) SetAdminSecretName(adminSN string) bool {
	r.AdminSecretName = adminSN
	return true
}

func (r *CephRbd) SetAdminSecretNamespace(adminSNs string) bool {
	r.AdminSecretNamespace = adminSNs
	return true
}

func (r *CephRbd) SetPool(pool string) bool {
	r.Pool = pool
	return true
}

func (r *CephRbd) SetUserId(uid string) bool {
	r.UserId = uid
	return true
}

func (r *CephRbd) SetUserSecretName(userSN string) bool {
	r.UserSecretName = userSN
	return true
}

func (r *CephRbd) SetFsType(fst string) bool {
	r.FsType = fst
	return true
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
			if !r.SetMonitors(strings.Join(monitors, ",")) {
				return errors.New(v["key"] + " format error")
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

func (r *CephRbd) SetImageFormat(imgFmt string) bool {
	r.ImageFormat = imgFmt
	return true
}

func (r *CephRbd) SetImageFeatures(imgFeat string) bool {
	r.ImageFeatures = imgFeat
	return true
}
