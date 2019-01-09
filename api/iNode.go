package api

import (
	"kboard/config"
	"net/http"
	"kboard/k8s"
	"kboard/k8s/resource"
	"log"
)

type INode struct {
	IApi
}

func NewINode(config *config.Config, w http.ResponseWriter, r *http.Request) *INode {
	node := &INode{
		IApi: *NewIApi(config, w, r),
	}
	node.Module = "node"
	return node
}

func (this *INode) Index() {
	// 获取node状态
	node := this.GetString("name")
	lib := k8s.NewNode(this.Config)
	if node == ""{
		this.TplEngine.Response(101, "", "缺少节点名称")
		return
	}
	data, httpErr := lib.Read(node)
	if httpErr.Code == 200 {
		this.TplEngine.Response(100, data, httpErr.Message)
		return
	}
	this.TplEngine.Response(99, "", httpErr.Message)
}


// @todo 节点扩容
func (this *INode) Scale() {
	lib := k8s.NewStatefulSet(this.Config)
	statefulSet := resource.NewResStatefulSet()
	statefulSet.SetMetaDataName("nginx")
	statefulSet.SetNamespace("myapp")
	statefulSet.SetReplicas(3)
	container := resource.NewContainer("mycontainer", "image")
	container.SetResource(resource.Resource{
		Limits: resource.Limits{
			Cpu: "0.5",
			Memory:"100Mi",
		},
		Requests: resource.Request{
			Cpu: "0.1",
			Memory: "50Mi",
		},
	})
	statefulSet.AddContainer(container)

	annos := map[string]string{
		"app":"nginx",
	}
	statefulSet.SetAnnotations(annos)
	labels := map[string]string{
		"app":"nginx",
	}
	statefulSet.SetLabels(labels)
	statefulSet.SetSelector(&resource.Selector{
		MatchLabels: labels,
	})
	statefulSet.SetServiceName("service name")
	statefulSet.SetStorage("1Gi")
	statefulSet.SetReplicas(3)
	statefulSet.SetVolumeClaimName("volume claim name")
	statefulSet.SetAccessMode("ReadWriteOnce")

	yamlData, err := statefulSet.ToYamlFile()
	if err != nil {
		log.Printf("%v", err)
	}
	res := lib.WriteToEtcd("myapp", "mystateful", yamlData)

	this.TplEngine.Response(100, res, "数据")
}

// @todo 节点隔离与恢复



// @todo 节点移除
func (this *INode) Delete()  {

}

