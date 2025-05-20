package topotypes

type TopoMask struct {
	Topos
	Model string `json:"model"`
}

func (sp *TopoMask) GetModel() string {
	return sp.Model
}

func (sp *TopoMask) SetModel(fileid string) {
	sp.Model = fileid
}
