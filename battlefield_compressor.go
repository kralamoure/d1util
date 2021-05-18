package retroutil

import (
	"errors"
	"strings"
)

func DecompressCells(data string, forced bool) ([]Cell, error) {
	if len(data)%10 != 0 {
		return nil, errors.New("invalid length of data")
	}

	cells := make([]Cell, len(data)/10)

	n := 0
	for i := 0; n < len(data); i++ {
		cell, err := decompressCell(data[n:n+10], forced, 0)
		if err != nil {
			return nil, err
		}
		cell.Id = i

		cells[i] = cell

		n += 10
	}

	return cells, nil
}

func decompressCell(data string, forced bool, permanentLevel int) (Cell, error) {
	cell := defaultGameMapCell()

	loc7 := len(data) - 1
	loc8 := make([]int, len(data))
	for loc7 >= 0 {
		loc8[loc7] = hashCodes[string(data[loc7])]
		loc7--
	}

	if (loc8[0]&32)>>5 == 0 {
		cell.Active = false
	} else {
		cell.Active = true
	}

	if cell.Active || forced {
		cell.PermanentLevel = permanentLevel

		if loc8[0]&1 == 0 {
			cell.LineOfSight = false
		} else {
			cell.LineOfSight = true
		}

		cell.LayerGroundRot = (loc8[1] & 48) >> 4
		cell.GroundLevel = loc8[1] & 15
		cell.Movement = (loc8[2] & 56) >> 3
		cell.LayerGroundNum = ((loc8[0] & 24) << 6) + ((loc8[2] & 7) << 6) + loc8[3]
		cell.GroundSlope = (loc8[4] & 60) >> 2

		if ((loc8[4] & 2) >> 1) == 0 {
			cell.LayerGroundFlip = false
		} else {
			cell.LayerGroundFlip = true
		}

		cell.LayerObject1Num = ((loc8[0] & 4) << 11) + ((loc8[4] & 1) << 12) + (loc8[5] << 6) + loc8[6]
		cell.LayerObject1Rot = (loc8[7] & 48) >> 4

		if (loc8[7]&8)>>3 == 0 {
			cell.LayerObject1Flip = false
		} else {
			cell.LayerObject1Flip = true
		}

		if (loc8[7]&4)>>2 == 0 {
			cell.LayerObject2Flip = false
		} else {
			cell.LayerObject2Flip = true
		}

		if (loc8[7]&2)>>1 == 0 {
			cell.LayerObject2Interactive = false
		} else {
			cell.LayerObject2Interactive = true
		}

		cell.LayerObject2Num = ((loc8[0] & 2) << 12) + ((loc8[7] & 1) << 12) + (loc8[8] << 6) + loc8[9]
		cell.LayerObjectExternal = ""
		cell.LayerObjectExternalInteractive = false
	}

	return cell, nil
}

func compressCells(cells []Cell) string {
	sb := &strings.Builder{}

	for _, cell := range cells {
		sb.WriteString(compressCell(cell))
	}

	return sb.String()
}

func compressCell(cell Cell) string {
	sb := &strings.Builder{}

	loc4 := make([]int, 10)

	active := 0
	if cell.Active {
		active = 1
	}
	loc4[0] = active << 5

	lineOfSight := 0
	if cell.LineOfSight {
		lineOfSight = 1
	}
	loc4[0] = loc4[0] | lineOfSight

	loc4[0] = loc4[0] | (cell.LayerGroundNum&1536)>>6
	loc4[0] = loc4[0] | (cell.LayerObject1Num&8192)>>11
	loc4[0] = loc4[0] | (cell.LayerObject2Num&8192)>>12

	loc4[1] = (cell.LayerGroundRot & 3) << 4
	loc4[1] = loc4[1] | cell.GroundLevel&15

	loc4[2] = (cell.Movement & 7) << 3
	loc4[2] = loc4[2] | cell.LayerGroundNum>>6&7

	loc4[3] = cell.LayerGroundNum & 63

	loc4[4] = (cell.GroundSlope & 15) << 2

	layerGroundFlip := 0
	if cell.LayerGroundFlip {
		layerGroundFlip = 1
	}
	loc4[4] = loc4[4] | layerGroundFlip<<1

	loc4[4] = loc4[4] | cell.LayerObject1Num>>12&1

	loc4[5] = cell.LayerObject1Num >> 6 & 63

	loc4[6] = cell.LayerObject1Num & 63

	loc4[7] = (cell.LayerObject1Rot & 3) << 4

	layerObject1Flip := 0
	if cell.LayerObject1Flip {
		layerObject1Flip = 1
	}
	loc4[7] = loc4[7] | layerObject1Flip<<3

	layerObject2Flip := 0
	if cell.LayerObject2Flip {
		layerObject2Flip = 1
	}
	loc4[7] = loc4[7] | layerObject2Flip<<2

	layerObject2Interactive := 0
	if cell.LayerObject2Interactive {
		layerObject2Interactive = 1
	}
	loc4[7] = loc4[7] | layerObject2Interactive<<1

	loc4[7] = loc4[7] | cell.LayerObject2Num>>12&1

	loc4[8] = cell.LayerObject2Num >> 6 & 63

	loc4[9] = cell.LayerObject2Num & 63

	for _, v := range loc4 {
		sb.WriteString(encode64(v))
	}

	return sb.String()
}

/*func compressPath(fullPathData []GameMapPath, withFirst bool) (string, error) {
	sb := &strings.Builder{}

	loc5, err := makeLightPath(fullPathData, withFirst)
	if err != nil {
		return "", err
	}

	loc11 := len(loc5)

	loc6 := 0
	for loc6 < loc11 {
		loc7 := loc5[loc6]
		loc8 := loc7.Dir & 7
		loc9 := (loc7.CellId & 4032) >> 6
		loc10 := loc7.CellId & 63

		sb.WriteRune(Encode64(loc8))
		sb.WriteRune(Encode64(loc9))
		sb.WriteRune(Encode64(loc10))

		loc6++
	}

	return sb.String(), nil
}

func makeLightPath(fullPath []GameMapPath, withFirst bool) ([]GameMapPath, error) {
	if len(fullPath) < 1 {
		return nil, errors.New("the path is empty")
	}

	var lightPath []GameMapPath

	if withFirst {
		lightPath = append(lightPath, fullPath[0])
	}

	loc6 := len(fullPath) - 1
	// TODO: not sure about "loc5 := -1" being right, and about its placement
	loc5 := -1
	for loc6 >= 0 {
		if fullPath[loc6].Dir != loc5 {
			lightPath = append(lightPath, fullPath[loc6])
			loc5 = fullPath[loc6].Dir
		}

		loc6--
	}
	return lightPath, nil
}

func extractFullPath(gameMap GameMap, compressedData string) ([]GameMapPath, error) {
	var gameMapPath []GameMapPath

	loc5 := make([]int, len(compressedData))

	loc7 := len(compressedData)
	loc8 := len(gameMap.Data)

	loc6 := 0
	for loc6 < loc7 {
		loc5[loc6] = Decode64(rune(compressedData[loc6]))
		loc5[loc6+1] = Decode64(rune(compressedData[loc6+1]))
		loc5[loc6+2] = Decode64(rune(compressedData[loc6+2]))

		loc9 := (loc5[loc6+1]&15)<<6 | loc5[loc6+2]

		if loc9 < 0 || loc9 > loc8 {
			return nil, errors.New("case not in game map")
		}

		gameMapPath = append(gameMapPath, GameMapPath{
			Dir:    loc5[loc6],
			CellId: loc9,
		})

		loc6 += 3
	}

	return gameMapPath, nil
}

func makeFullPath(gameMap GameMap, lightPath []GameMapPath) ([]int, error) {
	var gameMapPathCells []int

	loc6 := 0

	loc7 := gameMap.Width

	loc8 := []int{1, loc7, loc7*2 - 1, loc7 - 1, -1, -loc7, -loc7*2 + 1, -loc7 - 1}

	loc5 := lightPath[0].CellId

	gameMapPathCells = append(gameMapPathCells, loc5)

	loc9 := 1
	for loc9 < len(lightPath) {
		loc10 := lightPath[loc9].CellId
		loc11 := lightPath[loc9].Dir

		loc12 := 2*loc7 + 1
		// TODO: gameMapPathCells[loc6] seems like it will panic
		for gameMapPathCells[loc6] != loc10 {
			loc5 += loc8[loc11]
			// TODO: not sure about "loc6++" placement
			loc6++
			gameMapPathCells[loc6] = loc5

			// TODO: not sure about "loc12 -= 1" placement
			loc12 -= 1
			if loc12 < 0 {
				return nil, errors.New("impossible to walk")
			}
		}

		loc5 = loc10

		loc9++
	}

	return gameMapPathCells, nil
}
*/
