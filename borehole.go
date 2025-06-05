package topotypes

import (
	"encoding/json"

	"github.com/flywave/topotypes/geology"
)

type TopoBorehole struct {
	Topos
	geology.Borehole
	Materials map[string]*TopoMaterial `json:"materials,omitempty"`
}

func NewBorehole() *TopoBorehole {
	t := &TopoBorehole{}
	t.Type = TopoTypeToString(TOPO_TYPE_BOREHOLE)
	return t
}

func (b *TopoBorehole) GetMaterials() map[string]*TopoMaterial {
	return b.Materials
}
func (b *TopoBorehole) GetMaterialIds() []string {
	mtlids := []string{}
	for _, m := range b.Samples {
		mtlids = append(mtlids, m.MTL)
	}
	return mtlids
}

func BoreholeUnMarshal(js []byte) (*TopoBorehole, error) {
	mp := TopoBorehole{}
	e := json.Unmarshal(js, &mp)
	if e != nil {
		return nil, e
	}

	return &mp, nil
}
