package topotypes

type TopoSurface struct {
	Topos
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

type TopoTextureSurface struct {
	TopoPrism
	WarpS   string `json:"warp-s"`
	WarpT   string `json:"warp-t"`
	Zoom    uint8  `json:"zoom"`
	Texture string `json:"texture"`
}
