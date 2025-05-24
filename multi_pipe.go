package topotypes

import "encoding/json"

type TopoMultiSegmentPipe struct {
	TopoMaker
	Wires          [][][3]float64 `json:"-"`
	Profiles       []TopoProfile  `json:"profiles"`
	InnerProfiles  []TopoProfile  `json:"innerProfiles,omitempty"`
	SegmentTypes   []SegmentType  `json:"segmentTypes,omitempty"`
	TransitionMode string         `json:"transitionMode"`
	UpDir          *[3]float64    `json:"upDir,omitempty"`
}

func NewTopoMultiSegmentPipe() *TopoMultiSegmentPipe {
	t := &TopoMultiSegmentPipe{SegmentTypes: []SegmentType{}}
	t.Type = TopoTypeToString(TOPO_TYPE_MULTI_PIPE)
	return t
}

func (sp *TopoMultiSegmentPipe) IsTopoBound() bool {
	return true
}

func MultiPipeUnMarshal(js []byte) (*TopoMultiSegmentPipe, error) {
	mp := TopoMultiSegmentPipe{}
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
