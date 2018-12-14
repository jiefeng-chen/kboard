package api


import (
	"kboard/template"
	"kboard/config"
	"net/http"
)

type ITeam struct {
	Api
}

func NewITeam(config *config.Config, w http.ResponseWriter, r *http.Request) *ITeam {
	return &ITeam{
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

func (this *ITeam) Index() {
	this.TplEngine.Response(100, "", "数据")
}

// @todo 团队列表
// @todo 创建团队
// @todo 撤销团队
// @todo 申请加入
// @todo 退出团队

// @todo 团队信息

// @todo 成员列表
// @todo 设置和撤销管理员
// @todo 转移所有权（leader转移）
// @todo 审核通过
// @todo 拒绝加入

// @todo 团队项目列表
// @todo 创建项目
// @todo 删除项目
// @todo 项目转移

// @todo 团队镜像列表
// @todo 创建dockerfile
// @todo 删除dockerfile
// @todo 构建镜像
// @todo 上传镜像

