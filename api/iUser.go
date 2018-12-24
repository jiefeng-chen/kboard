package api

import (
	"kboard/template"
	"kboard/config"
	"net/http"
)

type IUser struct {
	Api
}

func NewIUser(config *config.Config, w http.ResponseWriter, r *http.Request) *IUser {
	return &IUser{
		Api{
			Config: config,
			TplEngine: template.NewTplEngine(w, r),
			Module: "index",
			Actions: map[string]func(){},
			R: r,
			W: w,
		},
	}
}

func (this *IUser) Index() {
	result := map[string]string{
		"email": "real_jf@163.com",
		"name": "real_jf",
	}
	this.TplEngine.Response(100, result, "数据")
}

// @todo 用户创建

// @todo 角色关联

// @todo 用户注销


