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

	err := request.GetConnection().SendMsg(200, []byte("ping....ping....ping...."))
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

func main() {

	s := net.NewServer()

	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})

	s.Serve()

}
