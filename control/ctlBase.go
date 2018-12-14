package control

import (
	"kboard/config"
	"fmt"
	"log"
	"kboard/template"
	"kboard/exception"
	"net/http"
	"kboard/utils"
)

type IControl interface {
	Register(string, func()) *Control
	Run(string)
}

type Control struct {
	Config *config.Config
	TplEngine *template.TplEngine
	Module string
	Actions map[string]func()
	W http.ResponseWriter
	R *http.Request
}

func NewControl(config *config.Config, w http.ResponseWriter, r *http.Request) *Control {
	return &Control{
		Config: config,
		TplEngine: template.NewTplEngine(w, r),
		Module: "base",
		Actions: map[string]func(){},
		R: r,
		W: w,
	}
}

func (c *Control) GetString(name string) string {
	return utils.ToString(c.R.URL.Query().Get(name))
}

func (c *Control) GetInt(name string) int {
	return utils.ToInt(c.R.URL.Query().Get(name))
}

func (c *Control) PostString(name string) string {
	return utils.ToString(c.R.FormValue(name))
}

func (c *Control) PostInt(name string) int {
	return utils.ToInt(c.R.FormValue(name))
}

func (c *Control) Register(action string, f func()) *Control {
	if c.Actions == nil {
		c.Actions = map[string]func(){}
	}
	if c.Module == "" {
		exception.CheckError(exception.NewError("error: control is empty!"), 999)
	}
	c.Actions[action] = f
	return c
}


func (c *Control) Run(action string) {
	// 注册全局变量
	if c.TplEngine.TplData["GModule"] == nil || c.TplEngine.TplData["GModule"] == "" {
		c.TplEngine.TplData["GModule"] = c.Module
	}
	if c.TplEngine.TplData["GAction"] == nil || c.TplEngine.TplData["GAction"] == "" {
		c.TplEngine.TplData["GAction"] = action
	}
	// 检查action方法是否存在
	if c.Actions[action] == nil {
		if c.Actions["index"] == nil {
			fmt.Fprintln(c.TplEngine.W, "404 page not found!")
			log.Println("404")
		}else{
			c.TplEngine.TplData["GAction"] = "index"
			c.Actions["index"]()
		}
	}else{
		// run action
		c.Actions[action]()
	}
}

func (c *Control) Index() {
	fmt.Fprintln(c.TplEngine.W, "hello world, this is default index")
}


