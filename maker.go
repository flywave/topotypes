package topotypes

type TopoMakerInterface interface {
	GetMaterials() map[string]*TopoMaterial
	IsInstance() bool
	GetMaterialId() string
}

type TopoMaker struct {
	Topos
	MeshMode   string           `json:"mode"`
	Materials  *TopoMaterialMap `json:"materials,omitempty"`
	Instanced  bool             `json:"instanced,omitempty"`
	MaterialId string           `json:"mtl_id,omitempty"`
}

func (t *TopoMaker) IsInstance() bool {
	return t.Instanced
}

func (t *TopoMaker) GetMaterials() map[string]*TopoMaterial {
	return (map[string]*TopoMaterial)(*t.Materials)
}

func (t *TopoMaker) GetMaterialId() string {
	return t.MaterialId
}