package core

import (
	"github.com/bitly/go-simplejson"
	"github.com/revel/config"
	"io/ioutil"
	"net/http"
	"log"
	"bytes"
)

type K8s struct {
	Config    *config.Context
}

func (k *K8s) baseApi() string {
	return k.Config.StringDefault("k8s", "http://192.168.37.150:8080")
}

func (*K8s) Status(json *simplejson.Json) string {
	kind, _ := json.Get("kind").String()
	if kind == "Status" {
		status, _ := json.Get("status").String();
		message, _ := json.Get("message").String();
		return status + ":" +  message
	}
	return "Success"
}


func (k *K8s) Post(url string, data []byte) *simplejson.Json {
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

func (k *K8s) Get(url string) *simplejson.Json {
	url = k.baseApi() + url
	//log.Println(url)
	response, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	json, _ := simplejson.NewJson([]byte(body))
	return json
}

func (k *K8s) Put(url string, data []byte) *simplejson.Json {
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
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	json, _ := simplejson.NewJson([]byte(body))
	return json
}

func (k *K8s) Del(url string) *simplejson.Json {
	url = k.baseApi() + url
	//log.Println(url)
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Panic(err)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	json, _ := simplejson.NewJson([]byte(body))
	return json
}

func (k *K8s) Patch(url string, data []byte) *simplejson.Json {
	url = k.baseApi() + url
	//log.Println(url)
	//log.Println(string(data))
	client := &http.Client{}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(data))
	if err != nil {
		log.Panic(err)
	}
	req.Header.Set("Content-Type", "application/strategic-merge-patch+json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	json, _ := simplejson.NewJson([]byte(body))
	return json
}
