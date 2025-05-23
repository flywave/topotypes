package topotypes

import "encoding/json"

type TopoRevol struct {
	TopoMaker
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

func RevolUnMarshal(js []byte) (*TopoRevol, error) {
	revol := TopoRevol{}
	e := json.Unmarshal(js, &revol)
	if e != nil {
		return nil, e
	}
	prof, er := ProfileUnMarshal(revol.Profile)
	if er != nil {
		return nil, er
	}
	revol.Profile = prof
	return &revol, nil
}
