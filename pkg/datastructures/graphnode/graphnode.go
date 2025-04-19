package graphnode

type GNode struct {
	Val       int
	Neighbors []*GNode
}

type Coordinate struct {
	Row int
	Col int
}

var Steps = [4]Coordinate{
	{Row: 1, Col: 0},
	{Row: -1, Col: 0},
	{Row: 0, Col: 1},
	{Row: 0, Col: -1},
}
