package core

import (
	"html/template"
	"net/http"
	"encoding/json"
	"control"
)

type ITplEngine interface {
	NewTemplate(string) *TplEngine
	Parse(string) *TplEngine
	ParseFiles(...string) *TplEngine
	AssignMap(map[string]interface{}) *TplEngine
	Assign(string, interface{}) *TplEngine
	Render()
	Response(int, interface{}, string)
}

type ResponseData struct {
	Code int
	Result interface{}
	Message string
}

type TplEngine struct {
	tpl *template.Template
	TplData map[string]interface{}
	W http.ResponseWriter
	R *http.Request
	C *control.CtlBase
}

func NewTplEngine(w http.ResponseWriter, r *http.Request, c *control.CtlBase) *TplEngine  {
	return &TplEngine{
		W: w,
		R: r,
		C: c,
	}
}

func (t *TplEngine) NewTemplate(name string) *TplEngine {
	t.tpl = template.New(name)
	return t
}

func (t *TplEngine) Parse(content string) *TplEngine {
	var err error
	t.tpl, err = t.tpl.Parse(content)
	CheckError(err, 2003)
	return t
}

func getDefinedTpl(file string) []string {
	file = "tpl/" + file + ".html"
	tpls := []string{
		file,
		"template/header.html",
		"template/menu.html",
		"template/footer.html",
		"template/pager.html",
		"template/aside.html",
		"template/comm_js.html",
	}

	return tpls
}

func (t *TplEngine) ParseFiles(tplName ...string) *TplEngine {
	commFiles := getDefinedTpl(tplName[0])
	for _, f := range tplName {
		ff := "template/" + f + ".html"
		commFiles = append(commFiles, ff)
	}
	var err error
	t.tpl, err = t.tpl.ParseFiles(commFiles...)
	CheckError(err, 2001)
	return t
}

func (t *TplEngine) AssignMap(data map[string]interface{}) *TplEngine {
	for k, v := range data {
		t.TplData[k] = v
	}
	return t
}

func (t *TplEngine) Assign(key string, data interface{}) *TplEngine {
	t.TplData[key] = data
	return t
}

func (t *TplEngine) Render() {
	err := t.tpl.Execute(t.W, t.TplData)
	CheckError(err, 2004)
}

func (t *TplEngine) Response(code int, result interface{}, message string)  {
	data := ResponseData{
		Code: code,
		Result: result,
		Message: message,
	}
	jsonData, err := json.Marshal(data)
	CheckError(err, 2005)
	t.W.Write(jsonData)
}