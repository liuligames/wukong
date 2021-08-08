package utils

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"wukong/iface"
)

type GlobalObj struct {
	TcpServer iface.IServer `yaml:"TcpServer"`
	Host      string `yaml:"Host"`
	TcpPort   int `yaml:"TcpPort"`
	Name      string `yaml:"Name"`

	Version        string `yaml:"Version"`
	MaxConn        int `yaml:"MaxConn"`
	MaxPackageSize uint32 `yaml:"MaxPackageSize"`
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/conf.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	GlobalObject = &GlobalObj{
		Name:           "LiuLIServerApp",
		Version:        "V0.7",
		TcpPort:        9527,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	GlobalObject.Reload()
}
