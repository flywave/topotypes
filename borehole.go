package topotypes

import (
	"encoding/json"

	"github.com/flywave/topotypes/geology"
	"github.com/flywave/topotypes/topo4d"
)

type TopoBorehole struct {
	Topos
	geology.Borehole
	Materials map[string]*TopoMaterial `json:"materials,omitempty"`
	UpDir     *[3]float64              `json:"upDir,omitempty"`
	Generate  *topo4d.Generate4D       `json:"generate,omitempty"`
	Work      *topo4d.TopoWork         `json:"works,omitempty"`
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

func (b *TopoBorehole) IsInstance() bool {
	return false
}

func (b *TopoBorehole) GetGenerate() *topo4d.Generate4D {
	return b.Generate
}

func (b *TopoBorehole) GetTopoWork() *topo4d.TopoWork {
	return b.Work
}

func (b *TopoBorehole) SetTopoWork(w *topo4d.TopoWork) {
	b.Work = w

}

func (b *TopoBorehole) GetShape() ParametricShape {
	return &b.Borehole
}

func (b *TopoBorehole) SetShape(s ParametricShape) {
	b.Borehole = *s.(*geology.Borehole)
}

func (tp *TopoBorehole) GetType() string {
	return tp.Type
}
