package topotypes

type TopoSurface struct {
	Topos
}

func (sp *TopoSurface) GetModel() string {
	return ""
}

type TopoLeveledSurface struct {
	TopoSurface
	Leveled string `json:"leveled"`
	RowSize uint32 `json:"row-size"`
}

func NewTopoLeveledSurface(lvlType int) *TopoLeveledSurface {
	t := &TopoLeveledSurface{Leveled: LeveledTypeToString(lvlType)}
	t.Type = TopoTypeToString(TOPO_TYPE_LEVELED_SURFACE)
	return t
}

type TopoMaterialSurface struct {
	TopoPrism
}
