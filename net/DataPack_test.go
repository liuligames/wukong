package net

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {

	listener, err := net.Listen("tcp", "127.0.0.1:9527")
	if err != nil {
		fmt.Println("server listen err", err)
		return
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("server accept error", err)
			}

			go func(conn net.Conn) {
				//拆包
				dp := NewDataPack()
				for {
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read head error", err)
						break
					}

					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("server unpack err", err)
						return
					}
					if msgHead.GetMsgLen() > 0 {
						msg := msgHead.(*Message)
						msg.Data = make([]byte, msg.GetMsgLen())

						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println("server unpack err", err)
							return
						}

						fmt.Println("------> Recv MsgId", msg.Id, ",msgLen = ", msg.MsgLen, "data=", string(msg.Data))
					}
				}

			}(conn)
		}
	}()

	conn, err := net.Dial("tcp", "127.0.0.1:9527")
	if err != nil {
		fmt.Println("client dial err:", err)
	}

	dp := NewDataPack()

	msg1 := &Message{
		Id:     1,
		MsgLen: 3,
		Data:   []byte{'l', 'i', 'u'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 error", err)
	}

	msg2 := &Message{
		Id:     2,
		MsgLen: 5,
		Data:   []byte{'g', 'a', 'm', 'e', 's'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg1 error", err)
	}
	sendData1 = append(sendData1, sendData2...)

	conn.Write(sendData1)

	select {}
}
