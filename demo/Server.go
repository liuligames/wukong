package main

import (
	"fmt"
	"wukong/iface"
	"wukong/net"
)

type PingRouter struct {
	net.BaseRouter
}

func (pr *PingRouter) PreHandle(request iface.IRequest) {
	fmt.Println("Call Router PreHandle ....")

	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping.... \n"))
	if err != nil {
		fmt.Println("call back before ping error ")
	}
}

func (pr *PingRouter) Handle(request iface.IRequest) {
	fmt.Println("Call Router Handle ....")

	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping... \n"))
	if err != nil {
		fmt.Println("call back ping...ping...ping... error ")
	}
}

func (pr *PingRouter) PostHandle(request iface.IRequest) {
	fmt.Println("Call Router PostHandle ....")

	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping.... \n"))
	if err != nil {
		fmt.Println("call back after ping error ")
	}
}

func main() {

	s := net.NewServer("[liu li games V0.3]")

	s.AddRouter(&PingRouter{})

	s.Serve()

}
