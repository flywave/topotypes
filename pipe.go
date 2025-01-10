package topotypes

import "encoding/json"

const (
	SMOOTH_NONE   = ""
	SMOOTH_SPLINE = "spline"
	SMOOTH_BEZIER = "bezier"
	SMOOTH_ARC    = "arc"
)

type TopoPipe struct {
	TopoMaker
	Wire           [][3]float64   `json:"-"`
	Profile        TopoProfile    `json:"profile"`
	InnerProfile   TopoProfile    `json:"inner_profile,omitempty"`
	Anchors        [2]*TopoAnchor `json:"anchors"`
	Smooth         string         `json:"smooth,omitempty"`
	TransitionMode string         `json:"transition_mode"`
}

func NewTopoPipe() *TopoPipe {
	t := &TopoPipe{Smooth: SMOOTH_NONE}
	t.Type = TopoTypeToString(TOPO_TYPE_PIPE)
	return t
}

func (sp *TopoPipe) IsTopoBoundy() bool {
	return true
}

func (sp *TopoPipe) GetAnchor() [2]*TopoAnchor {
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
