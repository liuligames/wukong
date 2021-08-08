package main

import (
	"fmt"
	"io"
	"net"
	"time"
	wkNet "wukong/net"
)

func main() {

	fmt.Println("client1 start...")
	time.Sleep(1 + time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:9527")
	if err != nil {
		fmt.Println("client start err exit! ")
		return
	}

	for {
		dp := wkNet.NewDataPack()
		binaryMsg, err := dp.Pack(wkNet.NewMessage(1, []byte("LiuLiGamesV0.7 client1 Test Message ")))
		if err != nil {
			fmt.Println("Pack error", err)
		}

		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("conn Write error", err)
			return
		}

		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head error", err)
			return
		}

		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("client unpack msgHead error", err)
			break
		}

		if msgHead.GetMsgLen() > 0 {
			msg := msgHead.(*wkNet.Message)
			msg.SetData(make([]byte, msg.GetMsgLen()))

			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg data error", err)
				return
			}

			fmt.Println("-----> Recv Server msgId= ",
				msg.Id, " len= ", msg.MsgLen, " data = ", string(msg.Data))
		}

		time.Sleep(1 * time.Second)
	}

}
