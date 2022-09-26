package topotypes

type TopoDecalRef struct {
	Ref string `json:"ref"`
}

type TopoDecal struct {
	Topos
	Size    *[2]float64    `json:"size,omitempty"`
	Depth   *float64       `json:"depth,omitempty"`
	Texture string         `json:"texture"`
	Refs    []TopoDecalRef `json:"targets,omitempty"`
	Density *float64       `json:"density,omitempty"`
}

func NewTopoDecal() *TopoDecal {
	t := &TopoDecal{}
	t.Type = TopoTypeToString(TOPO_TYPE_DECAL)
	return t
}
