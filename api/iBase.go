package api

import (
	"kboard/template"
	"kboard/config"
	"kboard/exception"
	"fmt"
	"log"
)

type IApi interface {
	Register(string, func()) *Api
	Run(string)
}

type Api struct {
	Config *config.Config
	TplEngine *template.TplEngine
	Module string
	Actions map[string]func()
}

func (i *Api) Register(action string, f func()) *Api {
	if i.Actions == nil {
		i.Actions = map[string]func(){}
	}
	if i.Module == "" {
		exception.CheckError(exception.NewError("error: api is empty!"), 999)
	}
	i.Actions[action] = f
	return i
}


func (i *Api) Run(action string) {
	// 注册全局变量
	if i.TplEngine.TplData["GModule"] == nil || i.TplEngine.TplData["GModule"] == "" {
		i.TplEngine.TplData["GModule"] = i.Module
	}
	if i.TplEngine.TplData["GAction"] == nil || i.TplEngine.TplData["GAction"] == "" {
		i.TplEngine.TplData["GAction"] = action
	}
	// 检查action方法是否存在
	if i.Actions[action] == nil {
		if i.Actions["index"] == nil {
			fmt.Fprintln(i.TplEngine.W, "404 page not found!")
			log.Println("404")
		}else{
			i.TplEngine.TplData["GAction"] = "index"
			i.Actions["index"]()
		}
	}else{
		// run action
		i.Actions[action]()
	}
}


