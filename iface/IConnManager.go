package iface

type IConnManager interface {
	Add(conn IConnection)

	Remove(conn IConnection)

	Get(connId uint32) (IConnection, error)

	Len() int

	ClearConn()
}
