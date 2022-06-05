package editor

import (
	"encoding/json"
	"errors"

	"github.com/flywave/go3d/quaternion"
	"github.com/flywave/go3d/vec3"
)

type Metadata struct {
	Scale       float64      `json:"scale"`
	Rotation    quaternion.T `json:"rotation"`
	Offset      vec3.T       `json:"offset"`
	Anchors     []*Anchor    `json:"anchors"`
	AnchorCount int          `json:"anchorcount"`
	Boards      []*Board     `json:"boards"`
	BoardCount  int          `json:"boardcount"`
	Components  []Component  `json:"components,omitempty"`
}

func MetadataUnMarshal(js []byte) (*Metadata, error) {
	base := Metadata{}
	e := json.Unmarshal(js, &base)
	if e != nil {
		return nil, e
	}
	for i := range base.Components {
		switch pro := base.Components[i].(type) {
		case map[string]interface{}:
			v, ok := pro["type"]
			t, ok2 := v.(string)
			if !ok || !ok2 {
				return nil, errors.New("components type error")
			}
			com_t := StringToComponentType(t)
			var c interface{}
			switch com_t {
			case COMPONENT_TYPE_SHAPE:
				c = &Shape{}
			case COMPONENT_TYPE_PRISM:
				c = &Prism{}
			case COMPONENT_TYPE_REVOL:
				c = &Revol{}
			case COMPONENT_TYPE_SOLID:
				c = &Solid{}
			case COMPONENT_TYPE_MODEL:
				c = &Model{}
			case COMPONENT_TYPE_STEEL_STRUCTURE:
				c = &SteelStructure{}
			case COMPONENT_TYPE_CATENARY:
				c = &Catenary{}
			}
			e := json.Unmarshal(([]byte)(js), c)
			if e != nil {
				return nil, e
			}
			base.Components[i] = c
		default:
			return nil, errors.New("components type error")
		}
	}
	return &base, nil
}
