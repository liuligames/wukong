package net

import (
	"errors"
	"fmt"
	"io"
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
		//	buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		//	_, err := c.Conn.Read(buf)
		//	if err != nil {
		//		fmt.Println("recv buf err", err)
		//		continue
		//	}

		dp := NewDataPack()

		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error", err)
			break
		}

		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error", err)
			break
		}

		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error", err)
				break
			}
		}

		msg.SetData(data)

		req := Request{
			conn: c,
			msg:  msg,
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

func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("Connection Closed when send msg ")
	}

	dp := NewDataPack()

	binaryMsg, err := dp.Pack(NewMessage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId)
		return errors.New("Pack error msg ")
	}

	if _, err := c.Conn.Write(binaryMsg); err != nil {
		fmt.Println("Write error msg id = ", msgId, "error :", err)
		return errors.New("conn Write error ")
	}

	return nil
}
