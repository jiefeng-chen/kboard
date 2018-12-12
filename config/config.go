package config

import (
	"github.com/BurntSushi/toml"
	"sync"
	"fmt"
)

type IConfig interface {
	LoadConfigFile(path string) *Config
	SetTSL(tsl *ServerTSL) error
	GetTSL() *ServerTSL
	GetAddress() string
	IsHttps() bool
	IsLog() bool
	IsAuth() bool
	IsHttp2() bool
	GetHttpVersion() string
}

type ServerTSL struct {
	Cert string
	Key  string
}

type Config struct {
	Path string
	Data TomlConfigData
	Once sync.Once // 实现单例模式
	Lock sync.RWMutex
}

type TomlConfigData struct {
	Server struct {
		Host        string
		Port        int
		Https       bool
		Log         bool   // 日志记录
		Auth        bool   // 鉴权
		HttpVersion string `toml:"httpVersion"` // 1.0, 1.1, 2.0
		TLS         struct {
			Cert string
			Key  string
		}
	}
	Database struct {
		Host     string
		Port     int
		Username string
		Password string
	}
	Memcache struct {
		Host string
		Port int
	}
	Redis struct {
		Host string
		Port int
	}
	Kubernetes struct {
		Host string
		Port int
	}
	Etcd struct {
		Host string
		Port int
	}
}

// load config file
// singleton
func NewConfig() *Config {
	return &Config{
		Path: "",
		Data: TomlConfigData{},
		Once: sync.Once{},
		Lock: sync.RWMutex{},
	}
}

// error code 1000 ~ 1200
func (c *Config) LoadConfigFile(path string) *Config {
	fmt.Println("loading config file...")
	c.Once.Do(func() {
		if path == "" {
			CheckError(NewError("path to config file is empty"), 1000)
		}
		c.Path = path
		if _, err := toml.DecodeFile(path, &c.Data); err != nil {
			CheckError(err, 1001)
		}
	})


	return c
}

// 重新加载配置文件
func (c *Config) ReloadConfigFile() {
	fmt.Println("reloading config file...")
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	c.Once.Do(func() {
		c.Lock.Lock()
		defer c.Lock.Unlock()
		if _, err := toml.DecodeFile(c.Path, &c.Data); err != nil {
			CheckError(err, 1001)
		}
	})
}

func (c *Config) SetTSL(tsl *ServerTSL) error {
	if tsl.Cert == "" || tsl.Key == "" {
		return NewError("server tsl contain invalid value")
	}
	c.Data.Server.TLS.Key = tsl.Key
	c.Data.Server.TLS.Cert = tsl.Cert
	return nil
}

func (c *Config) GetAddress() string {
	if c.Data.Server.Host == "" {
		CheckError(NewError("server host is empty"), 1004)
	}
	port := c.Data.Server.Port
	if port <= 0 || port > 65535 {
		CheckError(NewError("server port is invalid"), 1004)
	}
	return c.Data.Server.Host + ":" + ToString(port)
}

func (c *Config) GetTSL() *ServerTSL {
	cert := c.Data.Server.TLS.Cert
	key := c.Data.Server.TLS.Key
	if cert == "" || key == "" {
		CheckError(NewError("cert or key is empty"), 1005)
	}
	return &ServerTSL{
		Cert: cert,
		Key:  key,
	}
}

func (c *Config) IsHttps() bool {
	return c.Data.Server.Https
}

func (c *Config) IsLog() bool {
	return c.Data.Server.Log
}

func (c *Config) IsAuth() bool {
	return c.Data.Server.Auth
}

func (c *Config) IsHttp2() bool {
	if c.Data.Server.HttpVersion == "2.0" {
		return true
	}
	return false
}

func (c *Config) GetHttpVersion() string {
	var httpVersion string
	switch c.Data.Server.HttpVersion {
	case "1.0":
		fallthrough
	case "1.1":
		fallthrough
	case "2.0":
		httpVersion = c.Data.Server.HttpVersion
	default:
		httpVersion = "1.1"
	}
	return "HTTP/" + httpVersion
}
