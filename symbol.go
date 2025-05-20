package topotypes

type TopoSymbol struct {
	Topos
	Model     string              `json:"model"`
	Instanced bool                `json:"instanced"`
	Matrixs   map[int][16]float64 `json:"matrixs,omitempty"`
}

func (sp *TopoSymbol) GetModel() string {
	return sp.Model
}

func (sp *TopoSymbol) SetModel(fileid string) {
	sp.Model = fileid
}

func NewTopoSymbol() *TopoSymbol {
	t := &TopoSymbol{}
	t.Type = TopoTypeToString(TOPO_TYPE_SYMBOL)
	return t
}

type TopoSymbolPath struct {
	Topos
	Model   string  `json:"model"`
	Mode    string  `json:"mode"`
	Density float64 `json:"density"`
}

func (sp *TopoSymbolPath) GetModel() string {
	return sp.Model
}

func (sp *TopoSymbolPath) SetModel(fileid string) {
	sp.Model = fileid
}

func NewTopoSymbolPath(md int) *TopoSymbolPath {
	t := &TopoSymbolPath{Mode: PathModeToString(md)}
	t.Type = TopoTypeToString(TOPO_TYPE_SYMBOL_PATH)
	return t
}

type TopoSymbolSurface struct {
	TopoSurface
	Model string     `json:"model"`
	Mode  string     `json:"mode"`
	Cell  [2]float64 `json:"cell"`
}

func NewTopoSymbolSurface(md int) *TopoSymbolSurface {
	t := &TopoSymbolSurface{Mode: SurfaceModeToString(md)}
	t.Type = TopoTypeToString(TOPO_TYPE_SYMBOL_SURFACE)
	return t
}

func (sp *TopoSymbolSurface) GetModel() string {
	return sp.Model
}

func (sp *TopoSymbolSurface) SetModel(fileid string) {
	sp.Model = fileid
}
