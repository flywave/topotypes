package topotypes

type TopoHole struct {
	TopoMaker
	Depth           *float64 `json:"depth"`
	Texture         string   `json:"texture"`
	InnerTopFillet  *float64 `json:"inner-top-fillet"`
	InnerDownFillet *float64 `json:"inner-down-fillet"`
}

func NewTopoHole() *TopoHole {
	t := &TopoHole{}
	t.Type = TopoTypeToString(TOPO_TYPE_HOLE)
	return t
}
