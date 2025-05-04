package topotypes

import "encoding/json"

type TopoMultiPipe struct {
	TopoMaker
	Wires          [][][3]float64 `json:"-"`
	Profiles       []TopoProfile  `json:"profiles"`
	InnerProfiles  []TopoProfile  `json:"inner_profiles,omitempty"`
	Smooths        []string       `json:"smooths,omitempty"`
	TransitionMode string         `json:"transition_mode"`
}

func NewTopoMultiPipe() *TopoMultiPipe {
	t := &TopoMultiPipe{Smooths: []string{}}
	t.Type = TopoTypeToString(TOPO_TYPE_MULTI_PIPE)
	return t
}

func (sp *TopoMultiPipe) IsTopoBoundy() bool {
	return true
}

func MultiPipeUnMarshal(js []byte) (*TopoMultiPipe, error) {
	mp := TopoMultiPipe{}
	e := json.Unmarshal(js, &mp)
	if e != nil {
		return nil, e
	}

	var pros []TopoProfile
	for _, prof := range mp.Profiles {
		p, er := ProfileUnMarshal(prof)
		if er != nil {
			return nil, er
		}
		pros = append(pros, p)
	}
	mp.Profiles = pros
	var inpros []TopoProfile
	for _, prof := range mp.InnerProfiles {
		p, er := ProfileUnMarshal(prof)
		if er != nil {
			return nil, er
		}
		inpros = append(inpros, p)
	}
	mp.InnerProfiles = inpros
	return &mp, nil
}
