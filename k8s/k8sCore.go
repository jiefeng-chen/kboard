package k8s

import (
	"bytes"
	"kboard/exception"
	"log"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/revel/config"
	"io/ioutil"
	"net/http"
)

type ResultData struct {
	Code int
	Msg  string
	Ok   bool
}

func NewResultData() *ResultData {
	return &ResultData{
		Code: 0,
		Msg:  "",
		Ok:   false,
	}
}

const (
	STATUS_SUCCESS  = "Success"
	STATUS_FAILURE  = "Failure"
	STATUS_NOTFOUND = "NotFound"
)

type IK8sCore interface {
	Create(string, []byte) (err *HttpError)
	Replace(string, string, []byte) (err *HttpError)
	Read(string, string) (*simplejson.Json, *HttpError)
	Delete(string, string) (err *HttpError)
	WriteToEtcd(string, string, []byte) *HttpError
}

type K8sCore struct {
	Config *config.Context
	Kind   string
	Urls   Urls
}

func (k *K8sCore) baseApi() string {
	return k.Config.StringDefault("k8s", "http://192.168.52.227:8080")
}

func (l *K8sCore) Create(ns string, data []byte) *HttpError {
	url := fmt.Sprintf(l.Urls.Create, ns)
	jsonData := l.post(url, data)
	httpResult := GetHttpCode(jsonData)
	err := GetHttpErr(httpResult)
	if httpResult.Kind == l.Kind {
		err.Code = 200
		err.Message = "Success"
	} else if httpResult.Code == 200 || httpResult.Status == STATUS_SUCCESS {
		err.Code = 200
		err.Message = fmt.Sprintf("status:%s", err.Status)
	}
	// 404-不存在 409-已存在
	return err
}

func (l *K8sCore) Replace(ns string, name string, data []byte) *HttpError {
	url := fmt.Sprintf(l.Urls.Read, ns, name)
	jsonData := l.put(url, data)
	httpResult := GetHttpCode(jsonData)
	err := GetHttpErr(httpResult)
	if httpResult.Kind == l.Kind {
		err.Code = 200
		err.Message = "Success"
	} else if httpResult.Code == 200 || httpResult.Status == STATUS_SUCCESS {
		err.Code = 200
		err.Message = httpResult.Status + ":" + httpResult.Message
	}
	return err
}

func (l *K8sCore) Read(nsName string, name string) (*simplejson.Json, *HttpError) {
	url := fmt.Sprintf(l.Urls.Read, nsName, name)
	jsonData := l.get(url)
	httpResult := GetHttpCode(jsonData)
	err := GetHttpErr(httpResult)
	if httpResult.Kind == l.Kind {
		err.Code = 200
		err.Message = "Success"
	} else if httpResult.Code == 200 || httpResult.Status == STATUS_SUCCESS {
		err.Code = 200
		err.Message = httpResult.Status + ":" + httpResult.Message
	}
	return jsonData, err
}

func (l *K8sCore) Delete(ns string, name string) *HttpError {
	url := fmt.Sprintf(l.Urls.Read, ns, name)
	jsonData := l.del(url)
	httpResult := GetHttpCode(jsonData)
	err := GetHttpErr(httpResult)
	if httpResult.Kind == l.Kind {
		err.Code = 200
		err.Message = "Success"
		// 404-不存在 409-已存在
	} else if httpResult.Code == 200 || httpResult.Status == STATUS_SUCCESS {
		err.Code = 200
		err.Message = fmt.Sprintf("status:%s", err.Status)
	}
	return err
}

func (l *K8sCore) WriteToEtcd(ns string, name string, data []byte) *HttpError {
	// 1. 检查是否已存在
	_, err := l.Read(ns, name)
	if err.Code == 404 {
		// 不存在，创建
		err := l.Create(ns, data)
		if err != nil {
			return err
		}
	} else {
		// 已存在，直接覆盖
		err := l.Replace(ns, name, data)
		if err != nil {
			return err
		}
	}
	return &HttpError{
		Code:    200,
		Message: "Success",
		Status:  "Unknown",
	}
}

func (k *K8sCore) post(url string, data []byte) *simplejson.Json {
	url = k.baseApi() + url
	//log.Println(url)
	//log.Println(string(data))
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Panic(err)
	}
	req.Header.Set("Content-Type", "application/yaml; charset=utf-8")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	json, _ := simplejson.NewJson([]byte(body))
	return json
}

func (k *K8sCore) get(url string) *simplejson.Json {
	url = k.baseApi() + url
	//log.Println(url)
	response, err := http.Get(url)
	exception.CheckError(err, 80)

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	json, _ := simplejson.NewJson([]byte(body))
	return json
}

func (k *K8sCore) put(url string, data []byte) *simplejson.Json {
	url = k.baseApi() + url
	//log.Println(url)
	//log.Println(string(data))
	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		log.Panic(err)
	}
	req.Header.Set("Content-Type", "application/yaml; charset=utf-8")
	resp, err := client.Do(req)
	exception.CheckError(err, 82)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	exception.CheckError(err, 82)

	json, _ := simplejson.NewJson([]byte(body))
	return json
}

func (k *K8sCore) del(url string) *simplejson.Json {
	url = k.baseApi() + url
	//log.Println(url)
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	exception.CheckError(err, 81)

	resp, err := client.Do(req)
	exception.CheckError(err, 81)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	exception.CheckError(err, 81)

	json, _ := simplejson.NewJson([]byte(body))
	return json
}

func (k *K8sCore) patch(url string, data []byte) *simplejson.Json {
	url = k.baseApi() + url
	//log.Println(url)
	//log.Println(string(data))
	client := &http.Client{}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(data))
	exception.CheckError(err, 83)
	req.Header.Set("Content-Type", "application/strategic-merge-patch+json")
	resp, err := client.Do(req)
	exception.CheckError(err, 83)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	exception.CheckError(err, 83)
	json, _ := simplejson.NewJson([]byte(body))
	return json
}

type Urls struct {
	Read   string
	Create string
}

type HttpError struct {
	Code    int64
	Message string
	Status  string
}

type HttpResult struct {
	Code       int64
	Message    string
	Reason     string
	Status     string
	ApiVersion string
	Kind       string
}

func NewHttpResult() *HttpResult {
	return &HttpResult{
		Code:       0,
		Message:    "",
		Reason:     "",
		Status:     "",
		ApiVersion: "",
		Kind:       "",
	}
}

func (h *HttpResult) Parse(jsons *simplejson.Json) {
	h.ApiVersion, _ = jsons.Get("apiVersion").String()
	h.Status, _ = jsons.Get("status").String()
	h.Reason, _ = jsons.Get("reason").String()
	h.Message, _ = jsons.Get("message").String()
	h.Code, _ = jsons.Get("code").Int64()
	h.Kind, _ = jsons.Get("kind").String()
	if h.Status == "" {
		// 特殊处理，如namespace
		h.Status, _ = jsons.Get("status").Get("phase").String()
	}
}

func GetHttpCode(jsons *simplejson.Json) *HttpResult {
	httpResult := NewHttpResult()
	httpResult.Parse(jsons)
	return httpResult
}

func GetHttpErr(result *HttpResult) *HttpError {
	return &HttpError{
		Code:    result.Code,
		Message: "reason: " + result.Reason + ", message: " + result.Message,
		Status:  result.Status,
	}
}
