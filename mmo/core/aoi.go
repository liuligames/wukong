package core

import (
	"fmt"
)

const(
	AOI_MIN_X int = 85
	AOI_MAX_X int = 410
	AOI_CNT_X int = 10
	AOI_MIN_Y int = 75
	AOI_MAX_Y int = 400
	AOI_CNT_Y int = 20
)

type AoiManager struct {
	MinX int
	MaxX int
	CntX int

	MinY int
	MaxY int
	CntY int

	GrIds map[int]*Grid
}

func NewAoiManager(minX, maxX, cntX, minY, maxY, cntY int) *AoiManager {
	apiManager := &AoiManager{
		MinX:  minX,
		MaxX:  maxX,
		CntX:  cntX,
		MinY:  minY,
		MaxY:  maxY,
		CntY:  cntY,
		GrIds: make(map[int]*Grid),
	}

	for y := 0; y < cntY; y++ {
		for x := 0; x < cntX; x++ {
			gId := y*cntX + x

			apiManager.GrIds[gId] = NewGrid(gId,
				apiManager.MinX+x*apiManager.GrIdWidth(),
				apiManager.MinX+(x+1)*apiManager.GrIdWidth(),
				apiManager.MinY+y*apiManager.GrIdLength(),
				apiManager.MinY+(y+1)*apiManager.GrIdLength())
		}
	}

	return apiManager
}

func (am *AoiManager) GrIdWidth() int {
	return (am.MaxX - am.MinX) / am.CntX
}

func (am *AoiManager) GrIdLength() int {
	return (am.MaxY - am.MinY) / am.CntY
}

func (am *AoiManager) String() string {
	s := fmt.Sprintf("AoiManager:\n MinX:%d,MaxX:%d,CntX:%d,MinY:%d,MaxY:%d,CntY:%d\n GrIds in AoiManager",
		am.MinX, am.MaxX, am.CntX, am.MinY, am.MaxY, am.CntY)

	for _, grId := range am.GrIds {
		s += fmt.Sprintln(grId)
	}
	return s
}

func (am *AoiManager) GetSurroundGrIdsByGid(gId int) (grids []*Grid) {
	if _, ok := am.GrIds[gId]; !ok {
		return
	}

	grids = append(grids, am.GrIds[gId])

	idx := gId % am.CntX

	if idx > 0 {
		grids = append(grids, am.GrIds[gId-1])
	}
	if idx < am.CntX-1 {
		grids = append(grids, am.GrIds[gId+1])
	}

	gIdsX := make([]int, 0, len(grids))

	for _, value := range grids {
		gIdsX = append(gIdsX, value.GId)
	}

	for _, value := range gIdsX {
		idY := value / am.CntX
		if idY > 0 {
			grids = append(grids, am.GrIds[value-am.CntX])
		}

		if idY < am.CntY-1 {
			grids = append(grids, am.GrIds[value+am.CntX])
		}
	}

	return
}

func (am *AoiManager) GetGIdByPos(x, y float32) int {
	idx := (int(x) - am.MinX) / am.GrIdWidth()
	idy := (int(y) - am.MinY) / am.GrIdLength()

	return idy*am.CntX + idx
}

func (am *AoiManager) GetPIdsByPos(x, y float32) (playerIds []int) {
	gId := am.GetGIdByPos(x, y)
	grIds := am.GetSurroundGrIdsByGid(gId)

	for _, value := range grIds {
		playerIds = append(playerIds, value.GetPlayerIds()...)
		fmt.Printf("======》 grid id : %d, pids : %v =======\n", value.GId, value.GetPlayerIds())
	}
	return
}

func (am *AoiManager) AddPIdToGrId(pId, gId int) {
	am.GrIds[gId].Add(pId)
}

func (am *AoiManager) RemovePIdFromGrId(pId, gId int) {
	am.GrIds[gId].Remove(pId)
}

func (am *AoiManager) GetPIdByGId(gId int) (playerIds []int) {
	playerIds = am.GrIds[gId].GetPlayerIds()
	return
}

func (am *AoiManager) AddToGridByPos(pId int, x, y float32) {
	gId := am.GetGIdByPos(x, y)
	grId := am.GrIds[gId]
	grId.Add(pId)
}

func (am *AoiManager) RemoveFromGridByPos(pId int, x, y float32) {
	gId := am.GetGIdByPos(x, y)
	grId := am.GrIds[gId]
	grId.Remove(pId)
}
