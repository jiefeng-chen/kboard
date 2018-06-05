package resource

type resource interface {
	ToYamlFile() ([]byte, error)
}

// 资源类型
const (
	RESOURCE_CONFIG_MAP              = "configMap"
	RESOURCE_PERSISTENT_VOLUME_CLAIM = "pvc"
	RESOURCE_PERSISTENT_VOLUME       = "pv"
	RESOURCE_SECRET                  = "secret"
	RESOURCE_STORAGE_CLASS           = "storageClass"
)

func GetKinds() map[string]string {
	return map[string]string{
		RESOURCE_CONFIG_MAP:              "ConfigMap",
		RESOURCE_PERSISTENT_VOLUME_CLAIM: "PersistentVolumeClaim",
		RESOURCE_PERSISTENT_VOLUME:       "PersistentVolume",
		RESOURCE_SECRET:                  "Secret",
		RESOURCE_STORAGE_CLASS:           "storageClass",
	}
}

func GetVolumePlugins() map[string]string {
	return map[string]string{
		"RBD": "RBD(Ceph Block Device)",
	}
}

func GetAccessModes() map[string]string {
	return map[string]string{
		"ReadWriteOnce": "ReadWriteOnce(单点读写)", // 单节点读写 RWO
		"ReadOnlyMany":  "ReadOnlyMany(多点只读)",  // 多节点只读	 ROX
		"ReadWriteMany": "ReadWriteMany(多点读写)", // 多节点读写 RWX ceph rbd 不支持
	}
}

func GetVolumeModes() map[string]string {
	return map[string]string{
		"Block": "raw block devices(原始块设备)", // 原始块设备
		//"Filesystem": "Filesystem(文件系统)",		// 文件系统
	}
}

// 返回可用的provisioners
func GetProvisioners() map[string]string {
	return map[string]string{
		"kubernetes.io/rbd": "Ceph RBD", // {priovisioner} : 说明
	}
}

func GetFsTypes() map[string]string {
	return map[string]string{
		"xfs": "xfs",
		"nfs": "nfs",
	}
}

// 回收策略
func GetReclaimPolicy() map[string]string {
	return map[string]string{
		"Delete": "Delete",
		"Retain": "Retain",
	}
}
