package topotypes

type TopoCustom struct {
	TopoMaker
	CenterMode string   `json:"mode"`
	In         []string `json:"in-pipe-ids,omitempty"`
	Out        []string `json:"out-pipe-ids,omitempty"`
	SmoothEdge bool     `json:"smooth-edge"`
}

func NewTopoCustom() *TopoCustom {
	t := &TopoCustom{}
	t.Type = TopoTypeToString(TOPO_TYPE_CUSTOM)
	return t
}
