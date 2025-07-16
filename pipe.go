package topotypes

import (
	"encoding/json"

	"github.com/flywave/topotypes/anchor"
)

type SegmentType string

const (
	SEGMENT_LINE              SegmentType = "LINE"
	SEGMENT_SPLINE            SegmentType = "SPLINE"
	SEGMENT_THREE_POINT_ARC   SegmentType = "THREE_POINT_ARC"
	SEGMENT_CIRCLE_CENTER_ARC SegmentType = "CIRCLE_CENTER_ARC"
)

type TopoPipe struct {
	TopoParametric
	Wire           [][3]float64          `json:"-"`
	Profile        TopoProfile           `json:"profile"`
	InnerProfile   TopoProfile           `json:"inner_profile,omitempty"`
	Anchors        [2]*anchor.TopoAnchor `json:"anchors"`
	SegmentType    SegmentType           `json:"segment_type,omitempty"`
	TransitionMode string                `json:"transition_mode"`
}

func NewTopoPipe() *TopoPipe {
	t := &TopoPipe{SegmentType: SEGMENT_LINE}
	t.Type = TopoTypeToString(TOPO_TYPE_PIPE)
	return t
}

func (sp *TopoPipe) IsTopoBound() bool {
	return true
}

func (sp *TopoPipe) GetAnchor() [2]*anchor.TopoAnchor {
	return sp.Anchors
}

func (sp *TopoPipe) GetProfile() TopoProfile {
	return sp.Profile
}

func PipeUnMarshal(js []byte) (*TopoPipe, error) {
	pipe := TopoPipe{}
	e := json.Unmarshal(js, &pipe)
	if e != nil {
		return nil, e
	}
	if pipe.Profile != nil {
		prof, er := ProfileUnMarshal(pipe.Profile)
		if er != nil {
			return nil, er
		}
		pipe.Profile = prof
	}

	if pipe.InnerProfile != nil {
		prof, er := ProfileUnMarshal(pipe.InnerProfile)
		if er != nil {
			return nil, er
		}
		pipe.InnerProfile = prof
	}
	return &pipe, nil
}

func (t *TopoPipe) UnmarshalJSON(data []byte) error {
	stu := struct {
		Topos
		Materials      TopoMaterialMap       `json:"materials,omitempty"`
		MaterialId     string                `json:"mtl_id,omitempty"`
		Anchors        [2]*anchor.TopoAnchor `json:"anchors"`
		SegmentType    SegmentType           `json:"segment_type,omitempty"`
		TransitionMode string                `json:"transition_mode"`
		Profile        interface{}           `json:"profile"`
		InnerProfile   interface{}           `json:"inner_profile,omitempty"`
	}{}

	if err := json.Unmarshal(data, &stu); err != nil {
		return err
	}
	t.Topos = stu.Topos
	t.Materials = stu.Materials
	t.MaterialId = stu.MaterialId
	t.SegmentType = stu.SegmentType
	t.TransitionMode = stu.TransitionMode
	t.Profile = stu.Profile
	t.InnerProfile = stu.InnerProfile
	return nil
}
