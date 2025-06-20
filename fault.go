package topotypes

type TopoFault struct {
	Topos
}

func NewTopoFault() *TopoFault {
	t := &TopoFault{}
	t.Type = TopoTypeToString(TOPO_TYPE_FAULT)
	return t
}
