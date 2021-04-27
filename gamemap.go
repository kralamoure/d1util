package d1util

func BuiltCells(cellNum *int, buildAll bool, gameMapWidth int, cells []Cell) []Cell {
	out := make([]Cell, len(cells))
	copy(out, cells)

	loc9 := -1
	loc10 := 0
	loc11 := 0.0
	loc12 := out
	loc13 := len(loc12)
	loc14 := gameMapWidth - 1

	loc16 := false
	if cellNum != nil {
		loc16 = true
	}

	for loc20 := 0; loc20 < loc13; loc20++ {
		if loc9 == loc14 {
			loc9 = 0
			loc10++
			if loc11 == 0 {
				loc11 = cellHalfWidth
				loc14--
			} else {
				loc11 = 0
				loc14++
			}
		} else {
			loc9++
		}

		if loc16 {
			if loc20 < *cellNum {
				continue
			} else if loc20 > *cellNum {
				return out
			}
		}

		loc21 := loc12[loc20]
		if loc21.Active {
			loc22 := float64(loc9)*cellWidth + loc11
			loc23 := float64(loc10)*cellHalfHeight - cellHeight*float64(loc21.GroundLevel-7)
			loc21.X = loc22
			loc21.Y = loc23

			loc12[loc20] = loc21

			continue
		}

		if buildAll {
			var loc32 = float64(loc9)*cellWidth + loc11
			var loc33 = float64(loc10) * cellHalfHeight
			loc21.X = loc32
			loc21.Y = loc33

			loc12[loc20] = loc21
		}
	}

	return out
}
