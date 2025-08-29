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

func CatenaryUnmarshal(js []byte) (*TopoCatenary, error) {
	catenary := TopoCatenary{}
	e := json.Unmarshal(js, &catenary)
	if e != nil {
		return nil, e
	}
	if catenary.Profile != nil {
		prof, er := ProfileUnmarshal(catenary.Profile)
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

func (t *TopoCatenary) UnmarshalJSON(data []byte) error {
	stu := struct {
		Topos
		Materials  TopoMaterialMap       `json:"materials,omitempty"`
		MaterialId string                `json:"mtl_id,omitempty"`
		Anchors    [2]*anchor.TopoAnchor `json:"anchors"`
		catenary.Catenary
	}{}

	if err := json.Unmarshal(data, &stu); err != nil {
		return err
	}
	t.Topos = stu.Topos
	t.Materials = stu.Materials
	t.MaterialId = stu.MaterialId
	t.Anchors = stu.Anchors
	t.Catenary = stu.Catenary
	return nil
}
