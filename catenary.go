package topotypes

import (
	"encoding/json"

	"github.com/flywave/topotypes/anchor"
	"github.com/flywave/topotypes/catenary"
)

type TopoCatenary struct {
	TopoParametric
	Anchors [2]*anchor.TopoAnchor `json:"anchors"`
	catenary.Catenary
}

func CatenaryUnMarshal(js []byte) (*TopoCatenary, error) {
	catenary := TopoCatenary{}
	e := json.Unmarshal(js, &catenary)
	if e != nil {
		return nil, e
	}
	if catenary.Profile != nil {
		prof, er := ProfileUnMarshal(catenary.Profile)
		if er != nil {
			return nil, er
		}
		catenary.Profile = prof
	}
	return &catenary, nil
}

func (sp *TopoCatenary) GetAnchor() [2]*anchor.TopoAnchor {
	return sp.Anchors
}

func (sp *TopoCatenary) GetProfile() TopoProfile {
	return sp.Profile
}
