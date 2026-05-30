package component

import (
	"encoding/json"
	"errors"
)

type Metadata struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Version   uint64 `json:"version"`
	Transform *struct {
		Rotation  *[4]float64 `json:"rotation,omitempty"`
		Translate *[3]float64 `json:"translate,omitempty"`
		Scale     *[3]float64 `json:"scale,omitempty"`
	} `json:"transform,omitempty"`
	Anchors     []*Anchor   `json:"anchors,omitempty"`
	AnchorCount int         `json:"anchorcount,omitempty"`
	Components  []Component `json:"components,omitempty"`
}

func MetadataUnmarshal(js []byte) (*Metadata, error) {
	base := Metadata{}
	e := json.Unmarshal(js, &base)
	if e != nil {
		return nil, e
	}
	base.Components, e = ComponentsUnmarshal(base.Components)
	if e != nil {
		return nil, e
	}
	return &base, nil
}

func ComponentsUnmarshalMetadata(js []byte) (*Metadata, error) {
	var cmps []Component
	err := json.Unmarshal(js, &cmps)
	if err != nil {
		return nil, err
	}
	cmps, err = ComponentsUnmarshal(cmps)
	if err != nil {
		return nil, err
	}
	return &Metadata{Components: cmps}, nil
}

func ComponentsUnmarshal(components []Component) ([]Component, error) {
	for i := range components {
		switch pro := components[i].(type) {
		case map[string]interface{}:
			v, ok := pro["type"]
			t, ok2 := v.(string)
			if !ok || !ok2 {
				return nil, errors.New("components type error")
			}
			js2, _ := json.Marshal(components[i])
			com_t := StringToComponentType(t)
			var c interface{}
			switch com_t {
			case COMPONENT_TYPE_SHAPE:
				var err error
				c, err = ShapeUnmarshal(js2)
				if err != nil {
					return nil, err
				}
				components[i] = c
				continue
			case COMPONENT_TYPE_PRISM:
				var err error
				c, err = PrismUnmarshal(js2)
				if err != nil {
					return nil, err
				}
				components[i] = c
				continue
			case COMPONENT_TYPE_REVOL:
				var err error
				c, err = RevolUnmarshal(js2)
				if err != nil {
					return nil, err
				}
				components[i] = c
				continue
			case COMPONENT_TYPE_SOLID:
				c = &Solid{}
			case COMPONENT_TYPE_MODEL:
				c = &Model{}
			case COMPONENT_TYPE_DECAL:
				c = &Decal{}
			case COMPONENT_TYPE_BOARD:
				c = &Board{}
			case COMPONENT_TYPE_PARAMETRIC:
				c = &Parametric{}
			case COMPONENT_TYPE_CATENARY:
				var err error
				c, err = CatenaryUnmarshal(js2)
				if err != nil {
					return nil, err
				}
				components[i] = c
				continue
			}
			e := json.Unmarshal(([]byte)(js2), c)
			if e != nil {
				return nil, e
			}
			components[i] = c
		default:
			return nil, errors.New("components type error")
		}
	}

	return components, nil
}
