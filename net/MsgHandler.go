package net

import (
	"fmt"
	"strconv"
	"wukong/iface"
)

type MsgHandler struct {
	Apis map[uint32]iface.IRouter
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis: make(map[uint32]iface.IRouter),
	}
}

func (mh *MsgHandler) DoMsgHandler(request iface.IRequest) {
	handler, ok := mh.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println(" api msgId = ", request.GetMsgId(), "is not found Need Register")
	}

	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

func (mh *MsgHandler) AddRouter(msgId uint32, router iface.IRouter) {
	if _, ok := mh.Apis[msgId]; ok {
		panic("repeat api , msgId = " + strconv.Itoa(int(msgId)))
	}

	mh.Apis[msgId] = router
	fmt.Println("Add api msgId = ", msgId)
}
