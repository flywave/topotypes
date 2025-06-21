package topotypes

import "github.com/flywave/topotypes/geology"

type TopoFault struct {
	Topos
	geology.Fault
}

func NewTopoFault() *TopoFault {
	t := &TopoFault{}
	t.Type = TopoTypeToString(TOPO_TYPE_FAULT)
	return t
}
