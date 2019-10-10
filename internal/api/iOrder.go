package api

import (
	"kboard/config"
	"net/http"
)

type IOrder struct {
	IApi
}

func NewIOrder(config config.IConfig, w http.ResponseWriter, r *http.Request) *IOrder {
	order := &IOrder{
		IApi: *NewIApi(config, w, r),
	}
	order.Module = "order"
	return order
}

func (this *IOrder) Index() {

	this.TplEngine.Response(100, "", "")
}

// @todo 创建工单
func (this *IOrder) Edit() {

}

func (this *IOrder) Save() {

}

func (this *IOrder) List() {
	this.TplEngine.Response(100, "", "数据")
}
