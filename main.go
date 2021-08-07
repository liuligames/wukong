package main

import (
	"github.com/gin-gonic/gin"
	"github.com/panjf2000/gnet"
	"log"
	"net/http"
)

func main() {
	//// 初始化引擎
	//engine := gin.Default()
	//// 注册一个路由和处理函数
	//engine.GET("/", middleware1, middleware2, WebRoot)
	//
	//engine.GET("/user/:name", func(c *gin.Context) {
	//	// 使用 c.Param(key) 获取 url 参数
	//	name := c.Param("name")
	//	c.String(http.StatusOK, "Hello %s", name)
	//})
	//
	//// 绑定端口，然后启动应用
	//engine.Run(":9205")

	echo := new(echoServer)
	log.Fatal(gnet.Serve(echo, "tcp://:9000", gnet.WithMulticore(true)))
}

/**
* 根请求处理函数
* 所有本次请求相关的方法都在 context 中，完美
* 输出响应 hello, world
 */
func WebRoot(context *gin.Context) {
	firstname := context.DefaultQuery("firstname", "Guest")
	lastname := context.Query("lastname")
	context.String(http.StatusOK, "hello %s %s", firstname, lastname)
}

func middleware1(context *gin.Context) {
	log.Println("exec middleware1")
	a := 1

	log.Println(a)

}
func middleware2(context *gin.Context) {
	log.Println("exec middleware2")
}

type echoServer struct {
	*gnet.EventServer
}

func (es *echoServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	out = frame
	return
}
