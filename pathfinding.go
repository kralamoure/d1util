package d1util

import (
	"fmt"
	"math"
)

func AroundCellNum(cellNum, directionModerator, index, gameMapWidth int, cells []Cell) (int, bool) {
	if len(cells)-1 < cellNum {
		return 0, false
	}

	loc6 := gameMapWidth
	loc7 := []int{1, loc6, loc6*2 - 1, loc6 - 1, -1, -loc6, -loc6*2 + 1, -(loc6 - 1)}

	loc8 := 0
	switch index % 8 {
	case 0:
		loc8 = 2
	case 1:
		loc8 = 6
	case 2:
		loc8 = 4
	case 3:
		loc8 = 0
	case 4:
		loc8 = 3
	case 5:
		loc8 = 5
	case 6:
		loc8 = 1
	case 7:
		loc8 = 7
	}
	loc8 = (loc8 + directionModerator) % 8

	loc9 := cellNum + loc7[loc8]
	loc10 := cells

	if len(loc10)-1 < loc9 {
		return 0, false
	}

	loc11 := loc10[loc9]

	if loc11.Active && math.Abs(loc11.X-loc10[cellNum].X) <= cellWidth {
		return loc9, true
	}

	return 0, false
}

func DirectionToIndex(dir int) (int, error) {
	index := 0

	switch dir {
	case 0:
		index = 3
	case 1:
		index = 6
	case 2:
		index = 0
	case 3:
		index = 4
	case 4:
		index = 2
	case 5:
		index = 5
	case 6:
		index = 1
	case 7:
		index = 7
	default:
		return 0, fmt.Errorf("invalid direction %d", dir)
	}

	return index, nil
}
