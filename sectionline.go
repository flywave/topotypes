package topotypes

import "github.com/flywave/topotypes/geology"

type TopoSectionLine struct {
	Topos
	geology.SectionLine
}

func NewTopoSectionLine() *TopoSectionLine {
	t := &TopoSectionLine{}
	t.Type = TopoTypeToString(TOPO_TYPE_SECTION_LINE)
	return t
}
