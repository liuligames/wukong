package api

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"wukong/iface"
	"wukong/mmo/core"
	"wukong/mmo/pb/msg"
	"wukong/net"
)

type MoveApi struct {
	net.BaseRouter
}

func (m *MoveApi) Handle(request iface.IRequest) {
	protoMag := &msg.Position{}
	err := proto.Unmarshal(request.GetData(), protoMag)
	if err != nil {
		fmt.Println("Move : Position Unmarshal error", err)
		return
	}

	pId, err := request.GetConnection().GetProperty("PId")
	if err != nil {
		fmt.Println("GetProperty pId error", err)
		return
	}

	player := core.WorldManagerObj.GetPlayerByPId(pId.(int32))
	player.UpdatePos(protoMag.X, protoMag.Y, protoMag.Z, protoMag.V)
}
