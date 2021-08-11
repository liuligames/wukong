package core

import (
	"fmt"
	"testing"
)

func TestNewAoiManager(t *testing.T) {
	aoi := NewAoiManager(0, 300, 5, 0, 300, 5)

	fmt.Println(aoi)
}

func TestAoiManager_GetSurroundGrIdsByGid(t *testing.T) {
	aoi := NewAoiManager(0, 300, 5, 0, 300, 5)

	for k, _ := range aoi.GrIds {
		grids := aoi.GetSurroundGrIdsByGid(k)
		fmt.Println("gid = ", k, "grids len = ", len(grids))

		gIds := make([]int, 0, len(grids))

		for _, value := range grids {
			gIds = append(gIds, value.GId)
		}
		fmt.Println("grId = Ids are", gIds)
	}

}
