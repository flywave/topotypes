package topotypes

type TopoBoard struct {
	TopoMaker
	Size    *[2]float64 `json:"size,omitempty"`
	Depth   *float64    `json:"depth,omitempty"`
	Texture string      `json:"texture"`
}

func NewTopoBoard() *TopoBoard {
	t := &TopoBoard{}
	t.Type = TopoTypeToString(TOPO_TYPE_BOARD)
	return t
}
