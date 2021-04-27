package d1util

type Cell struct {
	Id                             int
	Active                         bool
	LineOfSight                    bool
	LayerGroundRot                 int
	GroundLevel                    int
	Movement                       int
	LayerGroundNum                 int
	GroundSlope                    int
	LayerGroundFlip                bool
	LayerObject1Num                int
	LayerObject1Rot                int
	LayerObject1Flip               bool
	LayerObject2Flip               bool
	LayerObject2Interactive        bool
	LayerObject2Num                int
	PermanentLevel                 int
	LayerObjectExternal            string
	LayerObjectExternalInteractive bool
	X                              float64
	Y                              float64
	SpriteOnId                     []int
}

func defaultGameMapCell() Cell {
	return Cell{
		Id:                      0,
		Active:                  true,
		LineOfSight:             true,
		LayerGroundRot:          0,
		GroundLevel:             7,
		Movement:                4,
		LayerGroundNum:          0,
		GroundSlope:             1,
		LayerGroundFlip:         false,
		LayerObject1Num:         0,
		LayerObject1Rot:         0,
		LayerObject1Flip:        false,
		LayerObject2Flip:        false,
		LayerObject2Interactive: false,
		LayerObject2Num:         0,
		X:                       0,
		Y:                       0,
		SpriteOnId:              nil,
	}
}

func getCellHeight(groundSlope, groundLevel int) float64 {
	loc4 := 0.5
	if groundSlope == 1 {
		loc4 = 0
	}

	loc5 := groundLevel - 7

	return float64(loc5) + loc4
}
