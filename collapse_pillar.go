package topotypes

type TopoCollapsePillar struct {
	Topos
}

func NewTopoCollapsePillar() *TopoCollapsePillar {
	t := &TopoCollapsePillar{}
	t.Type = TopoTypeToString(TOPO_TYPE_COLLAPSE_PILLAR)
	return t
}
