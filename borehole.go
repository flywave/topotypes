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
	Materials map[string]*TopoMaterial `json:"materials,omitempty"`
}

func NewBorehole() *Borehole {
	t := &Borehole{}
	t.Type = TopoTypeToString(TOPO_TYPE_BOREHOLE)
	return t
}

func BoreholeUnMarshal(js []byte) (*Borehole, error) {
	mp := Borehole{}
	e := json.Unmarshal(js, &mp)
	if e != nil {
		return nil, e
	}

	return &mp, nil
}
