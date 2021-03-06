package net

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"wukong/iface"
	"wukong/utils"
)

type Connection struct {
	TcpServer    iface.IServer
	Conn         *net.TCPConn
	ConnID       uint32
	isClosed     bool
	ExitChan     chan bool
	msgChan      chan []byte
	MsgHandler   iface.IMsgHandler
	property     map[string]interface{}
	propertyLock sync.RWMutex
}

func NewConnection(tcpServer iface.IServer, conn *net.TCPConn, connID uint32, msgHandler iface.IMsgHandler) *Connection {
	c := &Connection{
		TcpServer:  tcpServer,
		Conn:       conn,
		ConnID:     connID,
		MsgHandler: msgHandler,
		isClosed:   false,
		msgChan:    make(chan []byte),
		ExitChan:   make(chan bool, 1),
		property:   make(map[string]interface{}),
	}
	c.TcpServer.GetConnManager().Add(c)
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("reader goroutine is running.....")
	defer fmt.Println("connId = ", c.ConnID, "[reader is exit] remote addr is ", c.GetRemoteAddr().String())
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

		if utils.GlobalObject.WorkerPoolSize > 0 {
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			go c.MsgHandler.DoMsgHandler(&req)
		}
	}
}

func (c *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is running]")
	defer fmt.Println(c.GetRemoteAddr().String(), " [conn Writer exit]")

	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send data error ", err)
				return
			}
		case <-c.ExitChan:
			return
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("conn start ... connId  = ", c.ConnID)

	go c.StartReader()

	go c.StartWriter()

	c.TcpServer.CallOnConnStart(c)
}

func (c *Connection) Stop() {
	fmt.Println("conn stop ... connId  = ", c.ConnID)

	if c.isClosed {
		return
	}
	c.isClosed = true

	c.TcpServer.CallOnConnStop(c)

	_ = c.Conn.Close()

	c.ExitChan <- true

	c.TcpServer.GetConnManager().Remove(c)

	close(c.ExitChan)
	close(c.msgChan)
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

	c.msgChan <- binaryMsg

	return nil
}

func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	c.property[key] = value
}

func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()
	if value, ok := c.property[key]; ok {
		return value, nil
	} else {
		return nil, errors.New("no property found")
	}
}

func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}
