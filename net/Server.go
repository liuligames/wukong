package net

import (
	"fmt"
	"net"
	"wukong/iface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	Router    iface.IRouter
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP :%s , Port %d \n", s.IP, s.Port)

	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error :", err)
			return
		}

		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen:", s.IPVersion, "err", err)
			return
		}
		fmt.Println("start server,{", s.Name, "} Listenning...")

		var connId uint32
		connId = 0

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			dealConn := NewConnection(conn, connId, s.Router)
			connId++

			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()

	//todo 服务器额外业务

	select {}
}

func (s *Server) AddRouter(router iface.IRouter) {
	s.Router = router
	fmt.Println("Add Router !!!!")
}

func NewServer(name string) iface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      9527,
		Router:    nil,
	}
}
