package main

import (
	"fmt"
	"wukong/iface"
	"wukong/net"
)

type PingRouter struct {
	net.BaseRouter
}

//func (pr *PingRouter) PreHandle(request iface.IRequest) {
//	fmt.Println("Call Router PreHandle ....")
//
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping.... \n"))
//	if err != nil {
//		fmt.Println("call back before ping error ")
//	}
//}

func (pr *PingRouter) Handle(request iface.IRequest) {
	fmt.Println("Call Router Handle ....")

	fmt.Println("recv from client:msgId = ", request.GetMsgId(),
		"data = ", string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("ping....ping....ping...."))
	if err != nil {
		fmt.Println(err)
	}
}

//func (pr *PingRouter) PostHandle(request iface.IRequest) {
//	fmt.Println("Call Router PostHandle ....")
//
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping.... \n"))
//	if err != nil {
//		fmt.Println("call back after ping error ")
//	}
//}

func main() {

	s := net.NewServer()

	s.AddRouter(&PingRouter{})

	s.Serve()

}
