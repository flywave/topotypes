package editor

import "github.com/flywave/topotypes/utils"

const (
	COMPONENT_TYPE_NONE = iota
	COMPONENT_TYPE_MODEL
	COMPONENT_TYPE_PRISM
	COMPONENT_TYPE_REVOL
	COMPONENT_TYPE_SHAPE
	COMPONENT_TYPE_SOLID
	COMPONENT_TYPE_CATENARY
	COMPONENT_TYPE_STEEL_STRUCTURE
)

func ComponentTypeToString(tp int) string {
	switch tp {
	case COMPONENT_TYPE_MODEL:
		return "model"
	case COMPONENT_TYPE_REVOL:
		return "revol"
	case COMPONENT_TYPE_PRISM:
		return "prism"
	case COMPONENT_TYPE_SHAPE:
		return "shape"
	case COMPONENT_TYPE_SOLID:
		return "solid"
	case COMPONENT_TYPE_STEEL_STRUCTURE:
		return "steel-structure"
	case COMPONENT_TYPE_CATENARY:
		return "catenary"
	default:
		return ""
	}
}

func StringToComponentType(tp string) int {
	if utils.StrEquals(tp, "shape") {
		return COMPONENT_TYPE_SHAPE
	} else if utils.StrEquals(tp, "prism") {
		return COMPONENT_TYPE_PRISM
	} else if utils.StrEquals(tp, "revol") {
		return COMPONENT_TYPE_REVOL
	} else if utils.StrEquals(tp, "solid") {
		return COMPONENT_TYPE_SOLID
	} else if utils.StrEquals(tp, "model") {
		return COMPONENT_TYPE_MODEL
	} else if utils.StrEquals(tp, "steel-structure") {
		return COMPONENT_TYPE_STEEL_STRUCTURE
	} else if utils.StrEquals(tp, "catenary") {
		return COMPONENT_TYPE_CATENARY
	}
	return COMPONENT_TYPE_NONE
}

type Component interface{}

type BaseComponent struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Transform *struct {
		Rotation  *[4]float64 `json:"rotation,omitempty"`
		Translate *[3]float64 `json:"translate,omitempty"`
		Scale     *[3]float64 `json:"scale,omitempty"`
	} `json:"transform,omitempty"`
	BBox *[2][3]float64 `json:"bbox,omitempty"`
}