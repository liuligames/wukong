package net

import (
	"fmt"
	"net"
	"wukong/iface"
)

type Connection struct {
	Conn     *net.TCPConn
	ConnID   uint32
	isClosed bool
	ExitChan chan bool
	Router   iface.IRouter
}

func NewConnection(conn *net.TCPConn, connID uint32, router iface.IRouter) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
}

func (c *Connection) StartReader() {
	fmt.Println("reader goroutine is running.....")
	defer fmt.Println("connId = ", c.ConnID, "reader is exit remote addr is ", c.GetRemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}

		req := Request{
			conn: c,
			data: buf,
		}

		go func(request iface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

func (c *Connection) Start() {
	fmt.Println("conn start ... connId  = ", c.ConnID)

	go c.StartReader()
	//todo 启动从当前连接写数据的业务

}

func (c *Connection) Stop() {
	fmt.Println("conn stop ... connId  = ", c.ConnID)

	if c.isClosed {
		return
	}
	c.isClosed = true

	_ = c.Conn.Close()

	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID

}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	return nil
}
