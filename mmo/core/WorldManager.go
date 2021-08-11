package core

import "sync"

type WorldManager struct {
	AoiManager *AoiManager

	Players map[int32]*Player

	pLock sync.RWMutex
}

var WorldManagerObj *WorldManager

func init() {
	WorldManagerObj = &WorldManager{
		AoiManager: NewAoiManager(AOI_MIN_X, AOI_MAX_X, AOI_CNT_X, AOI_MIN_Y, AOI_MAX_Y, AOI_CNT_Y),
		Players:    make(map[int32]*Player),
	}
}

func (wm *WorldManager) AddPlayer(player *Player) {
	wm.pLock.Lock()
	wm.Players[player.PId] = player
	wm.pLock.Unlock()

	wm.AoiManager.AddToGridByPos(int(player.PId), player.X, player.Z)
}

func (wm *WorldManager) RemovePlayer(pId int32) {
	player := wm.Players[pId]
	wm.AoiManager.RemoveFromGridByPos(int(player.PId), player.X, player.Z)

	wm.pLock.Lock()
	delete(wm.Players, pId)
	wm.pLock.Unlock()
}

func (wm *WorldManager) GetPlayerByPId(pId int32) *Player {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()
	return wm.Players[pId]
}

func (wm *WorldManager) GetAllPlayers() []*Player {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()

	players := make([]*Player, 0)

	for _, value := range wm.Players {
		players = append(players, value)
	}
	return players
}
