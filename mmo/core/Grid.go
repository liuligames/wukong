package core

import (
	"fmt"
	"sync"
)

type Grid struct {
	GId       int
	MinX      int
	MaxX      int
	MinY      int
	MaxY      int
	PlayerIds map[int]bool
	PIdLock   sync.RWMutex
}

func NewGrid(gId, minX, maxX, minY, maxY int) *Grid {
	return &Grid{
		GId:       gId,
		MinX:      minX,
		MaxX:      maxX,
		MinY:      minY,
		MaxY:      maxY,
		PlayerIds: make(map[int]bool),
	}
}

func (g *Grid) Add(playerId int) {
	g.PIdLock.Lock()
	defer g.PIdLock.Unlock()

	g.PlayerIds[playerId] = true
}

func (g *Grid) Remove(playerId int) {
	g.PIdLock.Lock()
	defer g.PIdLock.Unlock()

	delete(g.PlayerIds, playerId)
}

func (g *Grid) GetPlayerIds() (playerIds []int) {
	g.PIdLock.RLock()
	defer g.PIdLock.RUnlock()
	for k, _ := range g.PlayerIds {
		playerIds = append(playerIds, k)
	}
	return
}

func (g *Grid) String() string {
	return fmt.Sprintf("Grid:%d, minX:%d, maxX:%d, minY:%d, maxY:%d, playerIds:%v",
		g.GId, g.MinX, g.MaxX, g.MinY, g.MaxY, g.PlayerIds)
}
