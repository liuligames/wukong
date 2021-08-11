package utils

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net"
	"net/http"
	"time"
	"wukong/iface"
)

type GlobalObj struct {
	TcpServer iface.IServer `yaml:"TcpServer"`
	Host      string        `yaml:"Host"`
	TcpPort   int           `yaml:"TcpPort"`
	Name      string        `yaml:"Name"`

	Version          string `yaml:"Version"`
	MaxConn          int    `yaml:"MaxConn"`
	MaxPackageSize   uint32 `yaml:"MaxPackageSize"`
	WorkerPoolSize   uint32 `yaml:"WorkerPoolSize"`
	MaxWorkerTaskLen uint32 `yaml:"MaxWorkerTaskLen"`
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
		Name:             "LiuLIServerApp",
		Version:          "V0.9",
		TcpPort:          9527,
		Host:             "0.0.0.0",
		MaxConn:          1000,
		MaxPackageSize:   4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
	}

	GlobalObject.Reload()
}

var HTTPTransport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second, // 连接超时时间
		KeepAlive: 60 * time.Second, // 保持长连接的时间
	}).DialContext, // 设置连接的参数
	MaxIdleConns:          500, // 最大空闲连接
	IdleConnTimeout:       60 * time.Second, // 空闲连接的超时时间
	ExpectContinueTimeout: 30 * time.Second, // 等待服务第一个响应的超时时间
	MaxIdleConnsPerHost:   100, // 每个host保持的空闲连接数
}