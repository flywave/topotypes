package topotypes

type TopoDecalRef struct {
	Ref string `json:"ref"`
}

type TopoDecal struct {
	Topos
	Size    *[2]float64    `json:"size"`
	Depth   *float64       `json:"depth"`
	Texture string         `json:"texture"`
	Refs    []TopoDecalRef `json:"targets,omitempty"`
}