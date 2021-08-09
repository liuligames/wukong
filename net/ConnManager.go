package net

import (
	"errors"
	"fmt"
	"sync"
	"wukong/iface"
)

type ConnManager struct {
	connections map[uint32]iface.IConnection
	connLock    sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]iface.IConnection),
	}
}

func (cm *ConnManager) Add(conn iface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	cm.connections[conn.GetConnID()] = conn
	fmt.Println("connId = ", conn.GetConnID(), " add to ConnManager conn num = ", cm.Len())

}

func (cm *ConnManager) Remove(conn iface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	delete(cm.connections, conn.GetConnID())
	fmt.Println("connId = ", conn.GetConnID(), " remove to ConnManager conn num = ", cm.Len())

}

func (cm *ConnManager) Get(connId uint32) (iface.IConnection, error) {
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()

	if conn, ok := cm.connections[connId]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

func (cm *ConnManager) Len() int {
	return len(cm.connections)
}

func (cm *ConnManager) ClearConn() {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	for connId, conn := range cm.connections {
		conn.Stop()

		delete(cm.connections, connId)
	}

	fmt.Println("clear all connection conn num = ", cm.Len())

}
