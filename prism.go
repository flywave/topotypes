package topotypes

import (
	"encoding/json"
)

type PrismInterface interface {
	GetTopoPrism() *TopoPrism
}

type TopoPrism struct {
	TopoParametric
	Profile   interface{} `json:"profile,omitempty"`
	Direction [3]float64  `json:"direction"`
	Height    float64     `json:"height"`
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

func PrismUnmarshal(js []byte) (*TopoPrism, error) {
	pris := TopoPrism{}
	e := json.Unmarshal(js, &pris)
	if e != nil {
		return nil, e
	}
	if pris.Profile != nil {
		prof, er := ProfileUnmarshal(pris.Profile)
		if er != nil {
			return nil, er
		}
		pris.Profile = prof
	}
	return &pris, nil
}

func (t *TopoPrism) UnmarshalJSON(data []byte) error {
	stu := struct {
		Topos
		Materials  TopoMaterialMap `json:"materials,omitempty"`
		MaterialId string          `json:"mtl_id,omitempty"`
		Profile    interface{}     `json:"profile,omitempty"`
		Direction  [3]float64      `json:"direction"`
		Height     float64         `json:"height"`
	}{}

	if err := json.Unmarshal(data, &stu); err != nil {
		return err
	}
	t.Topos = stu.Topos
	t.Materials = stu.Materials
	t.MaterialId = stu.MaterialId
	t.Profile = stu.Profile
	t.Direction = stu.Direction
	t.Height = stu.Height
	return nil
}
