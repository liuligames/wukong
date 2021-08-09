package main

import (
	"fmt"
	"wukong/iface"
	"wukong/net"
)

type PingRouter struct {
	net.BaseRouter
}

func (pr *PingRouter) Handle(request iface.IRequest) {
	fmt.Println("Call Router Handle ....")

	fmt.Println("recv from client:msgId = ", request.GetMsgId(),
		"data = ", string(request.GetData()))

	err := request.GetConnection().SendMsg(200, []byte("ping....ping....ping...."))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloRouter struct {
	net.BaseRouter
}

func (hr *HelloRouter) Handle(request iface.IRequest) {
	fmt.Println("Call HelloRouter Handle ....")

	fmt.Println("recv from client:msgId = ", request.GetMsgId(),
		"data = ", string(request.GetData()))

	err := request.GetConnection().SendMsg(201, []byte("Hello Games"))
	if err != nil {
		fmt.Println(err)
	}
}

func DoConnectionBegin(conn iface.IConnection) {
	fmt.Println("DoConnectionBegin is Called ...")
	if err := conn.SendMsg(202, []byte("DoConnection Begin")); err != nil {
		fmt.Println(err)
	}
}

func DoConnectionLost(conn iface.IConnection) {
	fmt.Println("DoConnectionLast is Called ...")
	fmt.Println("conn id = ", conn.GetConnID(), "is lost")
}

func main() {

	s := net.NewServer()

	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})

	s.Serve()

}
