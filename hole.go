package topotypes

type TopoHole struct {
	TopoMaker
	Depth  *float64    `json:"depth"`
	Fillet *[4]float64 `json:"fillet"`
}

func NewTopoHole() *TopoHole {
	t := &TopoHole{}
	t.Type = TopoTypeToString(TOPO_TYPE_HOLE)
	return t
}
