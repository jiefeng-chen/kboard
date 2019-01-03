package api

import (
	"kboard/config"
	"net/http"
)

type IUser struct {
	IApi
}

func NewIUser(config *config.Config, w http.ResponseWriter, r *http.Request) *IUser {
	user := &IUser{
		IApi: *NewIApi(config, w, r)}
	user.Module = "user"
	return user
}

func (this *IUser) Index() {
	result := map[string]string{
		"email": "real_jf@163.com",
		"name": "real_jf",
	}
	this.TplEngine.ResponseWithHeader(100, result, "数据", this.Header)
}

// @todo 用户创建
func (this *IUser) Create() {

}

// @todo 角色关联

// @todo 用户注销


