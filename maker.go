package topotypes

type TopoMakerInterface interface {
	GetMaterials() []*TopoMaterial
	IsInstance() bool
}

type TopoMaker struct {
	Topos
	MeshMode  string          `json:"mode"`
	Materials []*TopoMaterial `json:"materials,omitempty"`
	Instanced bool            `json:"instanced,omitempty"`
}

func (t *TopoMaker) IsInstance() bool {
	return t.Instanced
}

func (t *TopoMaker) GetMaterials() []*TopoMaterial {
	return t.Materials
}
