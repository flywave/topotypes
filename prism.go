package topotypes

import (
	"encoding/json"
)

type PrismInterface interface {
	GetTopoPrism() *TopoPrism
}

type TopoPrism struct {
	TopoMaker
	Profile   interface{} `json:"profile,omitempty"`
	Direction [3]float64  `json:"dir"`
}

func NewTopoPrism() *TopoPrism {
	t := &TopoPrism{}
	t.Type = TopoTypeToString(TOPO_TYPE_PRISM)
	return t
}

func (sp *TopoPrism) GetTopoPrism() *TopoPrism {
	return sp
}

func (sp *TopoPrism) IsTopoBound() bool {
	return true
}

func PrismUnMarshal(js []byte) (*TopoPrism, error) {
	pris := TopoPrism{}
	e := json.Unmarshal(js, &pris)
	if e != nil {
		return nil, e
	}
	if pris.Profile != nil {
		prof, er := ProfileUnMarshal(pris.Profile)
		if er != nil {
			return nil, er
		}
		pris.Profile = prof
	}
	return &pris, nil
}
