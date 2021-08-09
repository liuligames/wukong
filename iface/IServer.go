package iface

type IServer interface {
	Start()
	Stop()
	Serve()

	AddRouter(msgId uint32, router IRouter)
	GetConnManager() IConnManager

	SetOnConnStart(hookFunc func(connection IConnection))

	SetOnConnStop(hookFunc func(connection IConnection))

	CallOnConnStart(connection IConnection)

	CallOnConnStop(connection IConnection)
}
