package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"net"
	"time"
	pMsg "wukong/mmo/pb/msg"
	wkNet "wukong/net"
)

func main() {

	fmt.Println("--------------> 陆离频道 <--------------")
	time.Sleep(1 + time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:9527")
	if err != nil {
		fmt.Println("client start err exit! ")
		return
	}

	for i := 0; i < 2; i++ {
		dp := wkNet.NewDataPack()
		//binaryMsg, err := dp.Pack(wkNet.NewMessage(0, []byte("LiuLiGamesV0.9 client0 Test Message ")))
		//if err != nil {
		//	fmt.Println("Pack error", err)
		//}
		//
		//if _, err := conn.Write(binaryMsg); err != nil {
		//	fmt.Println("conn Write error", err)
		//	return
		//}

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

			data := &pMsg.SyncPid{}
			err := proto.Unmarshal(msg.Data, data)
			if err != nil {
				fmt.Println("Unmarshal msg data error", err)
				return
			}

			fmt.Println("-----> Recv Server msgId= ",
				msg.Id, " len= ", msg.MsgLen, " data = ", data)
		}

		time.Sleep(1 * time.Second)
	}
	fmt.Println("-----> stop for")

	for {

		time.Sleep(5 * time.Second)

		fmt.Println("-----> start talk")

		talk(conn,"我是你爸爸~")
	}


}

func talk(conn net.Conn,content string)  {

	data := &pMsg.Talk{
		Content: content,
		Name: "陆离",
	}

	marshalMsg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("Marshal msg err :", err)
		return
	}

	dp := wkNet.NewDataPack()
	binaryMsg, err := dp.Pack(wkNet.NewMessage(2, marshalMsg))
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
		return
	}

	if msgHead.GetMsgLen() > 0 {
		msg := msgHead.(*wkNet.Message)
		msg.SetData(make([]byte, msg.GetMsgLen()))

		if _, err := io.ReadFull(conn, msg.Data); err != nil {
			fmt.Println("read msg data error", err)
			return
		}

		data := &pMsg.BroadCast{}
		err := proto.Unmarshal(msg.Data, data)
		if err != nil {
			fmt.Println("Unmarshal msg data error", err)
			return
		}

		fmt.Println("-----> Recv Server msgId= ",
			msg.Id, " len= ", msg.MsgLen, " data = ", data)
	}
}
