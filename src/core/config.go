package core

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type IConfig interface {
	LoadConfigFile(path string) error
	SetTSL(tsl *ServerTSL) error
	GetAddress() (string, error)
}

type ServerTSL struct {
	Cert string
	Key string
}

type Config struct {
	Path string
	Data *YamlConfigData
}

type YamlConfigData struct {
	Server struct{
		Address string
		TLS struct{
			Cert string
			Key string
		}
	}
	Database struct{
		Host string
		Port int
		Username string
		Password string
	}
	Memcache struct{
		Host string
		Port int
	}
	Redis struct{
		Host string
		Port int
	}
}

func NewConfig() *Config {
	return &Config{
		Path: "",
		Data: &YamlConfigData{},
	}
}


// error code 1000 ~ 1200
func (c *Config) LoadConfigFile(path string) error {
	if path == "" {
		return NewError("path to config file is empty")
	}
	c.Path = path

	yamlFile, err := ioutil.ReadFile(path)
	CheckError(err, 1001)

	err = yaml.Unmarshal(yamlFile, c.Data)
	CheckError(err, 1002)

	return nil
}

func (c *Config) SetTSL(tsl *ServerTSL) error {
	if tsl.Cert == "" || tsl.Key == "" {
		return NewError("server tsl contain invalid value")
	}
	c.Data.Server.TLS.Key = tsl.Key
	c.Data.Server.TLS.Cert = tsl.Cert
	return nil
}

func (c *Config) GetAddress() (string, error) {
	if c.Data.Server.Address == "" {
		return "", NewError("server address is empty")
	}
	return c.Data.Server.Address, nil
}