package topotypes

import (
	"encoding/json"
)

type PrismInterface interface {
	GetDirection() *[3]float64
}

type TopoPrism struct {
	TopoMaker
	Profile      interface{} `json:"profile,omitempty"`
	UntilProfile interface{} `json:"until_profile,omitempty"`
	Direction    [3]float64  `json:"direction"`
}

func NewTopoPrism() *TopoPrism {
	t := &TopoPrism{}
	t.Type = TopoTypeToString(TOPO_TYPE_PRISM)
	return t
}

func (sp *TopoPrism) GetDirection() *[3]float64 {
	return &sp.Direction
}

func (sp *TopoPrism) IsTopoBoundy() bool {
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
	if pris.UntilProfile != nil {
		prof, er := ProfileUnMarshal(pris.UntilProfile)
		if er != nil {
			return nil, er
		}
		pris.UntilProfile = prof
	}
	return &pris, nil
}
