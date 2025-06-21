package topotypes

import "github.com/flywave/topotypes/geology"

type TopoCollapsePillar struct {
	Topos
	geology.CollapsePillar
}

func NewTopoCollapsePillar() *TopoCollapsePillar {
	t := &TopoCollapsePillar{}
	t.Type = TopoTypeToString(TOPO_TYPE_COLLAPSE_PILLAR)
	return t
}
