package topotypes

type TopoMakerInterface interface {
	GetMaterials() map[string]*TopoMaterial
	IsInstance() bool
	GetMaterialIds() []string
	etModel() string
}

type TopoMaker struct {
	Topos
	MeshMode   string           `json:"mode"`
	Materials  *TopoMaterialMap `json:"materials,omitempty"`
	Instanced  bool             `json:"instanced,omitempty"`
	MaterialId string           `json:"mtl_id,omitempty"`
}

func (sp *TopoMaker) GetModel() string {
	return ""
}

func (t *TopoMaker) IsInstance() bool {
	return t.Instanced
}

func (t *TopoMaker) GetMaterials() map[string]*TopoMaterial {
	if t.Materials == nil {
		return map[string]*TopoMaterial{}
	}
	return (map[string]*TopoMaterial)(*t.Materials)
}

func (t *TopoMaker) GetMaterialIds() []string {
	if t.MaterialId == "" {
		t.MaterialId = "mtl_0"
	}
	return []string{t.MaterialId}
}
