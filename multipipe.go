package topotypes

import (
	"encoding/json"

	"github.com/flywave/topotypes/anchor"
)

type TopoMultiSegmentPipe struct {
	TopoParametric
	Wires          [][][3]float64        `json:"-"`
	Profiles       []TopoProfile         `json:"profiles"`
	InnerProfiles  []TopoProfile         `json:"innerProfiles,omitempty"`
	SegmentTypes   []SegmentType         `json:"segmentTypes,omitempty"`
	TransitionMode string                `json:"transitionMode"`
	UpDir          *[3]float64           `json:"upDir,omitempty"`
	Anchors        [2]*anchor.TopoAnchor `json:"anchors"`
}

func NewTopoMultiSegmentPipe() *TopoMultiSegmentPipe {
	t := &TopoMultiSegmentPipe{SegmentTypes: []SegmentType{}}
	t.Type = TopoTypeToString(TOPO_TYPE_MULTI_PIPE)
	return t
}

func (sp *TopoMultiSegmentPipe) IsTopoBound() bool {
	return true
}

func (sp *TopoMultiSegmentPipe) GetAnchor() [2]*anchor.TopoAnchor {
	return sp.Anchors
}

func MultiPipeUnmarshal(js []byte) (*TopoMultiSegmentPipe, error) {
	mp := TopoMultiSegmentPipe{}
	e := json.Unmarshal(js, &mp)
	if e != nil {
		return nil, e
	}

	var pros []TopoProfile
	for _, prof := range mp.Profiles {
		p, er := ProfileUnmarshal(prof)
		if er != nil {
			return nil, er
		}
		pros = append(pros, p)
	}
	mp.Profiles = pros
	var inpros []TopoProfile
	for _, prof := range mp.InnerProfiles {
		p, er := ProfileUnmarshal(prof)
		if er != nil {
			return nil, er
		}
		inpros = append(inpros, p)
	}
	mp.InnerProfiles = inpros
	return &mp, nil
}

func (t *TopoMultiSegmentPipe) UnmarshalJSON(data []byte) error {
	stu := struct {
		Topos
		Materials      TopoMaterialMap       `json:"materials,omitempty"`
		MaterialId     string                `json:"mtl_id,omitempty"`
		Profiles       []TopoProfile         `json:"profiles"`
		InnerProfiles  []TopoProfile         `json:"innerProfiles,omitempty"`
		SegmentTypes   []SegmentType         `json:"segmentTypes,omitempty"`
		TransitionMode string                `json:"transitionMode"`
		UpDir          *[3]float64           `json:"upDir,omitempty"`
		Anchors        [2]*anchor.TopoAnchor `json:"anchors"`
	}{}

	if err := json.Unmarshal(data, &stu); err != nil {
		return err
	}
	t.Topos = stu.Topos
	t.Materials = stu.Materials
	t.MaterialId = stu.MaterialId
	t.Anchors = stu.Anchors
	t.Profiles = stu.Profiles
	t.InnerProfiles = stu.InnerProfiles
	t.SegmentTypes = stu.SegmentTypes
	t.TransitionMode = stu.TransitionMode
	t.UpDir = stu.UpDir
	return nil
}
