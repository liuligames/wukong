package main

import (
	"fmt"
	"wukong/iface"
	"wukong/mmo/api"
	"wukong/mmo/core"
	"wukong/net"
)

func OnConnectionAdd(conn iface.IConnection) {
	player := core.NewPlayer(conn)

	player.SyncPId()

	player.BroadCastStartPosition()

	core.WorldManagerObj.AddPlayer(player)

	conn.SetProperty("PId", player.PId)

	fmt.Println("===========》 Player id = ", player.PId, "is arrived ========")
}

func main() {
	s := net.NewServer()

	s.SetOnConnStart(OnConnectionAdd)

	s.AddRouter(2, &api.WorldChatApi{})

	s.Serve()
}
