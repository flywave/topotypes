package topotypes

type TopoSectionLine struct {
	Topos
}

func NewTopoSectionLine() *TopoSectionLine {
	t := &TopoSectionLine{}
	t.Type = TopoTypeToString(TOPO_TYPE_SECTION_LINE)
	return t
}
