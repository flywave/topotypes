package topotypes

import "encoding/json"

type TopoPipe struct {
	TopoMaker
	Wire           [][3]float64   `json:"-"`
	Profile        TopoProfile    `json:"profile"`
	Anchors        [2]*TopoAnchor `json:"anchors"`
	Customs        [2]*string     `json:"customs"`
	Smooth         string         `json:"smooth,omitempty"`
	TransitionMode string         `json:"transition_mode"`
}

func NewTopoPipe() *TopoPipe {
	t := &TopoPipe{}
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
	return &pipe, nil
}
