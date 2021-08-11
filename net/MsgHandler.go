package net

import (
	"fmt"
	"strconv"
	"wukong/iface"
	"wukong/utils"
)

type MsgHandler struct {
	Apis           map[uint32]iface.IRouter
	TaskQueue      []chan iface.IRequest
	WorkerPoolSize uint32
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]iface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan iface.IRequest, utils.GlobalObject.WorkerPoolSize),
	}
}

func (mh *MsgHandler) DoMsgHandler(request iface.IRequest) {
	handler, ok := mh.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println(" api msgId = ", request.GetMsgId(), "is not found Need Register")
		return
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

func (mh *MsgHandler) StartWorkerPool() {
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		mh.TaskQueue[i] = make(chan iface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}

}

func (mh *MsgHandler) StartOneWorker(workerId int, taskQueue chan iface.IRequest) {
	fmt.Println("Worker Id = ", workerId, "is started ...")

	for {
		select {
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

func (mh *MsgHandler) SendMsgToTaskQueue(request iface.IRequest) {
	workerId := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	fmt.Println("add connId = ", request.GetConnection().GetConnID(),
		"request msgId = ", request.GetMsgId(), "to workerId = ", workerId)
	mh.TaskQueue[workerId] <- request

}
