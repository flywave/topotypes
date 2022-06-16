package editor

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
			js2, _ := json.Marshal(base.Components[i])
			com_t := StringToComponentType(t)
			var c interface{}
			switch com_t {
			case COMPONENT_TYPE_SHAPE:
				var err error
				c, err = ShapeUnMarshal(js2)
				if err != nil {
					return nil, err
				}
				base.Components[i] = c
				continue
			case COMPONENT_TYPE_PRISM:
				var err error
				c, err = PrismUnMarshal(js2)
				if err != nil {
					return nil, err
				}
				base.Components[i] = c
				continue
			case COMPONENT_TYPE_REVOL:
				var err error
				c, err = RevolUnMarshal(js2)
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
				c, err = SteelStructureUnMarshal(js2)
				if err != nil {
					return nil, err
				}
				base.Components[i] = c
				continue
			case COMPONENT_TYPE_CATENARY:
				var err error
				c, err = CatenaryUnMarshal(js2)
				if err != nil {
					return nil, err
				}
				base.Components[i] = c
				continue
			}
			e := json.Unmarshal(([]byte)(js2), c)
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
