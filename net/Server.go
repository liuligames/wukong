package net

import (
	"fmt"
	"net"
	"wukong/iface"
	"wukong/utils"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	Router    iface.IRouter
}

func (s *Server) Start() {
	fmt.Printf("[liuli] Server Name : %s listenner at IP : %s ,Port : %d is starting \n",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[liuli] Version : %s MaxConn : %d ,MaxPackageSize : %d \n",
		utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPackageSize)
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

func NewServer() iface.IServer {
	return &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:    nil,
	}
}
