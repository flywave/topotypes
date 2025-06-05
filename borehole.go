package topotypes

import "encoding/json"

type BoreholeSample struct {
	Name      string                 `json:"name,omitempty"`
	DepthFrom float64                `json:"depth_from"`
	DepthTo   float64                `json:"depth_to"`
	MTL       string                 `json:"mtl,omitempty"`
	Property  map[string]interface{} `json:"property,omitempty"`
}

type Borehole struct {
	Topos
	Samples   []*BoreholeSample        `json:"samples"`
	Diameter  *float64                 `json:"diameter,omitempty"`
	Materials map[string]*TopoMaterial `json:"materials,omitempty"`
}

func NewBorehole() *Borehole {
	t := &Borehole{}
	t.Type = TopoTypeToString(TOPO_TYPE_BOREHOLE)
	return t
}

func (b *Borehole) GetMaterials() map[string]*TopoMaterial {
	return b.Materials
}
func (b *Borehole) GetMaterialIds() []string {
	mtlids := []string{}
	for _, m := range b.Samples {
		mtlids = append(mtlids, m.MTL)
	}
	return mtlids
}

func BoreholeUnMarshal(js []byte) (*Borehole, error) {
	mp := Borehole{}
	e := json.Unmarshal(js, &mp)
	if e != nil {
		return nil, e
	}

	return &mp, nil
}
