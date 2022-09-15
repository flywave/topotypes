package topotypes

type TopoHole struct {
	TopoMaker
	Depth     *float64    `json:"depth,omitempty"`
	Fillet    *[4]float64 `json:"fillet,omitempty"`
	Offset    *[2]float64 `json:"offset,omitempty"`
	Border    *[2]float64 `json:"border,omitempty"`
	HasBottom *bool       `json:"has-bottom,omitempty"`
}

func NewTopoHole() *TopoHole {
	t := &TopoHole{}
	t.Type = TopoTypeToString(TOPO_TYPE_HOLE)
	return t
}
