package net

import "wukong/iface"

type Request struct {
	conn iface.IConnection
	msg iface.IMessage
}

func (r *Request) GetConnection() iface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgId() uint32  {
	return r.msg.GetMsgId()
}