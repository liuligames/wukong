package api

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"wukong/iface"
	"wukong/mmo/core"
	"wukong/mmo/pb/msg"
	"wukong/net"
)

type WorldChatApi struct {
	net.BaseRouter
}

func (wc WorldChatApi) Handle(request iface.IRequest) {
	protoMag := &msg.Talk{}
	err := proto.Unmarshal(request.GetData(), protoMag)
	if err != nil {
		fmt.Println("talk Unmarshal error : ", err)
		return
	}

	pId, err := request.GetConnection().GetProperty("PId")
	if err != nil {
		fmt.Println("get Property PId error : ", err)
		return
	}

	player := core.WorldManagerObj.GetPlayerByPId(pId.(int32))

	player.Talk(fmt.Sprintf("玩家 [%s] : %s", protoMag.Name, protoMag.Content))
}
