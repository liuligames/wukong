package net

import (
	"fmt"
	"net"
	"wukong/iface"
	"wukong/utils"
)

type Server struct {
	Name        string
	IPVersion   string
	IP          string
	Port        int
	MsgHandler  iface.IMsgHandler
	ConnManager iface.IConnManager
	OnConnStart func(conn iface.IConnection)
	OnConnStop  func(conn iface.IConnection)
}

func (s *Server) Start() {
	fmt.Printf("[liuli] Server Name : %s listenner at IP : %s ,Port : %d is starting \n",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[liuli] Version : %s MaxConn : %d ,MaxPackageSize : %d \n",
		utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPackageSize)
	fmt.Printf("[Start] Server Listenner at IP :%s , Port %d \n", s.IP, s.Port)

	go func() {

		s.MsgHandler.StartWorkerPool()

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

			if s.ConnManager.Len() >= utils.GlobalObject.MaxConn {
				fmt.Println("too many connections maxConn = ", utils.GlobalObject.MaxConn)
				if err := conn.Close(); err != nil {
					fmt.Println("MaxConn Close error :", err)
				}
				continue
			}

			dealConn := NewConnection(s, conn, connId, s.MsgHandler)
			connId++

			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("stop server name", s.Name)
	s.ConnManager.ClearConn()
}

func (s *Server) Serve() {
	s.Start()

	//todo 服务器额外业务

	select {}
}

func (s *Server) AddRouter(msgId uint32, router iface.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
	fmt.Println("Add Router !!!!")
}

func (s *Server) GetConnManager() iface.IConnManager {
	return s.ConnManager
}

func NewServer() iface.IServer {
	return &Server{
		Name:        utils.GlobalObject.Name,
		IPVersion:   "tcp4",
		IP:          utils.GlobalObject.Host,
		Port:        utils.GlobalObject.TcpPort,
		MsgHandler:  NewMsgHandler(),
		ConnManager: NewConnManager(),
	}
}

func (s *Server) SetOnConnStart(hookFunc func(connection iface.IConnection))  {
	s.OnConnStart = hookFunc
}

func (s *Server) SetOnConnStop(hookFunc func(connection iface.IConnection)) {
	s.OnConnStop = hookFunc
}

func (s *Server) CallOnConnStart(connection iface.IConnection) {
	if s.OnConnStart != nil{
		fmt.Println("Call OnConnStart()")
		s.OnConnStart(connection)
	}
}

func (s *Server) CallOnConnStop(connection iface.IConnection) {
	if s.OnConnStop != nil{
		fmt.Println("Call OnConnStart()")
		s.OnConnStop(connection)
	}
}



