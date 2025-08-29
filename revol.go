package topotypes

import (
	"encoding/json"
)

type TopoRevol struct {
	TopoParametric
	Profile interface{}   `json:"profile"`
	Axis    [2][3]float64 `json:"axis"`
	Angle   float64       `json:"angle"`
}

func NewTopoRevol() *TopoRevol {
	t := &TopoRevol{}
	t.Type = TopoTypeToString(TOPO_TYPE_REVOL)
	return t
}

func (sp *TopoRevol) IsTopoBound() bool {
	return true
}

func RevolUnmarshal(js []byte) (*TopoRevol, error) {
	revol := TopoRevol{}
	e := json.Unmarshal(js, &revol)
	if e != nil {
		return nil, e
	}
	prof, er := ProfileUnmarshal(revol.Profile)
	if er != nil {
		return nil, er
	}
	revol.Profile = prof
	return &revol, nil
}

func (t *TopoRevol) UnmarshalJSON(data []byte) error {
	stu := struct {
		Topos
		Materials  TopoMaterialMap `json:"materials,omitempty"`
		MaterialId string          `json:"mtl_id,omitempty"`
		Profile    interface{}     `json:"profile"`
		Axis       [2][3]float64   `json:"axis"`
		Angle      float64         `json:"angle"`
	}{}

	if err := json.Unmarshal(data, &stu); err != nil {
		return err
	}
	t.Topos = stu.Topos
	t.Materials = stu.Materials
	t.MaterialId = stu.MaterialId
	t.Profile = stu.Profile
	t.Axis = stu.Axis
	t.Angle = stu.Angle
	return nil
}
