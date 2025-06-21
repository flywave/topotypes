package topotypes

import (
	"encoding/json"

	"github.com/flywave/topotypes/geology"
)

type TopoBorehole struct {
	Topos
	geology.Borehole
	UpDir *[3]float64 `json:"upDir,omitempty"`
}

func NewBorehole() *TopoBorehole {
	t := &TopoBorehole{}
	t.Type = TopoTypeToString(TOPO_TYPE_BOREHOLE)
	return t
}

func BoreholeUnMarshal(js []byte) (*TopoBorehole, error) {
	mp := TopoBorehole{}
	e := json.Unmarshal(js, &mp)
	if e != nil {
		return nil, e
	}

	return &mp, nil
}
