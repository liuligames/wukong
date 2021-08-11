package core

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"math/rand"
	"sync"
	"wukong/iface"
	"wukong/mmo/pb/msg"
)

type Player struct {
	PId  int32
	Conn iface.IConnection
	X    float32
	Y    float32
	Z    float32
	V    float32
}

var id int32 = 0
var idLock sync.Mutex

func NewPlayer(conn iface.IConnection) *Player {
	idLock.Lock()
	id++
	idLock.Unlock()

	return &Player{
		PId:  id,
		Conn: conn,
		X:    float32(160 + rand.Intn(10)),
		Y:    0,
		Z:    float32(140 + rand.Intn(20)),
		V:    0,
	}
}

func (p *Player) SendMsg(msgId uint32, data proto.Message) {
	marshalMsg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("Marshal msg err :", err)
		return
	}

	if p.Conn == nil {
		fmt.Println("conn in player is nil")
		return
	}

	if err := p.Conn.SendMsg(msgId, marshalMsg); err != nil {
		fmt.Println("player sendMsg error!")
		return
	}

	return
}

func (p *Player) SyncPId() {
	protoMsg := &msg.SyncPid{
		PId: p.PId,
	}

	p.SendMsg(1, protoMsg)
}

func (p *Player) BroadCastStartPosition() {
	protoMsg := &msg.BroadCast{
		PId: p.PId,
		Tp:  2,
		Data: &msg.BroadCast_P{
			P: &msg.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	p.SendMsg(200, protoMsg)
}

func (p *Player) Talk(content string) {
	protoMsg := &msg.BroadCast{
		PId: p.PId,
		Tp:  1,
		Data: &msg.BroadCast_Content{
			Content: content,
		},
	}

	players := WorldManagerObj.GetAllPlayers()

	for _, player := range players {
		player.SendMsg(200, protoMsg)
	}

}
