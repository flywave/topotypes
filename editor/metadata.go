package editor

import (
	"encoding/json"
	"errors"

	quatd "github.com/flywave/go3d/float64/quaternion"
	vec3d "github.com/flywave/go3d/float64/vec3"
)

type Metadata struct {
	Scale       float64     `json:"scale"`
	Rotation    quatd.T     `json:"rotation"`
	Offset      vec3d.T     `json:"offset"`
	Anchors     []*Anchor   `json:"anchors,omitempty"`
	AnchorCount int         `json:"anchorcount,omitempty"`
	Boards      []*Board    `json:"boards,omitempty"`
	BoardCount  int         `json:"boardcount,omitempty"`
	Components  []Component `json:"components,omitempty"`
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
				var err error
				c, err = ShapeUnMarshal(js)
				if err != nil {
					return nil, err
				}
				base.Components[i] = c
				continue
			case COMPONENT_TYPE_PRISM:
				var err error
				c, err = PrismUnMarshal(js)
				if err != nil {
					return nil, err
				}
				base.Components[i] = c
				continue
			case COMPONENT_TYPE_REVOL:
				var err error
				c, err = RevolUnMarshal(js)
				if err != nil {
					return nil, err
				}
				base.Components[i] = c
				continue
			case COMPONENT_TYPE_SOLID:
				c = &Solid{}
			case COMPONENT_TYPE_MODEL:
				c = &Model{}
			case COMPONENT_TYPE_STEEL_STRUCTURE:
				var err error
				c, err = SteelStructureUnMarshal(js)
				if err != nil {
					return nil, err
				}
				base.Components[i] = c
				continue
			case COMPONENT_TYPE_CATENARY:
				var err error
				c, err = CatenaryUnMarshal(js)
				if err != nil {
					return nil, err
				}
				base.Components[i] = c
				continue
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
