package topotypes

import (
	"encoding/json"
	"errors"

	"github.com/flywave/topotypes/profile"
	"github.com/flywave/topotypes/utils"
)

const (
	TOPO_PROFILE_TYPE_NONE      = profile.TYPE_NONE
	TOPO_PROFILE_TYPE_TRIANGLE  = profile.TYPE_TRIANGLE
	TOPO_PROFILE_TYPE_RECTANGLE = profile.TYPE_RECTANGLE
	TOPO_PROFILE_TYPE_CIRC      = profile.TYPE_CIRC
	TOPO_PROFILE_TYPE_ELIPS     = profile.TYPE_ELIPS
	TOPO_PROFILE_TYPE_POLYGON   = profile.TYPE_POLYGON
	TOPO_PROFILE_TYPE_L_STEEL   = profile.TYPE_L_STEEL
	TOPO_PROFILE_TYPE_ARC       = profile.TYPE_ARC
	TOPO_PROFILE_TYPE_EDGE      = profile.TYPE_EDGE
)

func ProfileTypeToString(tp int) string {
	switch tp {
	case TOPO_PROFILE_TYPE_TRIANGLE:
		return "triangle"
	case TOPO_PROFILE_TYPE_RECTANGLE:
		return "rectangle"
	case TOPO_PROFILE_TYPE_CIRC:
		return "circ"
	case TOPO_PROFILE_TYPE_ELIPS:
		return "ellipse"
	case TOPO_PROFILE_TYPE_POLYGON:
		return "polygon"
	case TOPO_PROFILE_TYPE_L_STEEL:
		return "lsteel"
	case TOPO_PROFILE_TYPE_ARC:
		return "arc"
	case TOPO_PROFILE_TYPE_EDGE:
		return "edge"
	default:
		return ""
	}
}

func StringToProfileType(tp string) int {
	if utils.StrEquals(tp, "triangle") {
		return TOPO_PROFILE_TYPE_TRIANGLE
	} else if utils.StrEquals(tp, "rectangle") {
		return TOPO_PROFILE_TYPE_RECTANGLE
	} else if utils.StrEquals(tp, "circ") {
		return TOPO_PROFILE_TYPE_CIRC
	} else if utils.StrEquals(tp, "ellipse") {
		return TOPO_PROFILE_TYPE_ELIPS
	} else if utils.StrEquals(tp, "polygon") {
		return TOPO_PROFILE_TYPE_POLYGON
	} else if utils.StrEquals(tp, "lsteel") {
		return TOPO_PROFILE_TYPE_L_STEEL
	} else if utils.StrEquals(tp, "arc") {
		return TOPO_PROFILE_TYPE_ARC
	} else if utils.StrEquals(tp, "edge") {
		return TOPO_PROFILE_TYPE_EDGE
	}
	return TOPO_PROFILE_TYPE_NONE
}

type TopoProfile profile.Profile

type TopoTriangle profile.Triangle

func NewTopoTriangle() *TopoTriangle {
	t := TopoTriangle{}
	t.Type = ProfileTypeToString(TOPO_PROFILE_TYPE_TRIANGLE)
	return &t
}

type TopoRectangle profile.Rectangle

func NewTopoRectangle() *TopoRectangle {
	t := TopoRectangle{}
	t.Type = ProfileTypeToString(TOPO_PROFILE_TYPE_RECTANGLE)
	return &t
}

type TopoCirc profile.Circ

func NewTopoCirc() *TopoCirc {
	t := TopoCirc{}
	t.Type = ProfileTypeToString(TOPO_PROFILE_TYPE_CIRC)
	return &t
}

type TopoElips profile.Elips

func NewTopoElips() *TopoElips {
	t := TopoElips{}
	t.Type = ProfileTypeToString(TOPO_PROFILE_TYPE_ELIPS)
	return &t
}

type TopoPolygon profile.Polygon

func NewTopoPolygon() *TopoPolygon {
	t := TopoPolygon{}
	t.Type = ProfileTypeToString(TOPO_PROFILE_TYPE_POLYGON)
	return &t
}

func ProfileUnMarshal(inter interface{}) (interface{}, error) {
	switch pro := inter.(type) {
	case map[string]interface{}:
		v, ok := pro["type"]
		t, ok2 := v.(string)
		if !ok || !ok2 {
			return nil, errors.New("profile type error")
		}
		pro_t := StringToProfileType(t)
		js, er := json.Marshal(inter)
		if er != nil {
			return nil, er
		}
		var pf interface{}
		switch pro_t {
		case TOPO_PROFILE_TYPE_TRIANGLE:
			pf = NewTopoTriangle()
		case TOPO_PROFILE_TYPE_RECTANGLE:
			pf = NewTopoRectangle()
		case TOPO_PROFILE_TYPE_CIRC:
			pf = NewTopoCirc()
		case TOPO_PROFILE_TYPE_ELIPS:
			pf = NewTopoElips()
		case TOPO_PROFILE_TYPE_POLYGON:
			pf = NewTopoPolygon()
		case profile.TYPE_L_STEEL:
			pf = profile.NewLShape()
		case profile.TYPE_ARC:
			pf = profile.NewArcCircle()
		case profile.TYPE_EDGE:
			pf = profile.NewEdge()
		default:
			return nil, errors.New("profile type error")
		}
		e := json.Unmarshal(([]byte)(js), pf)
		if e != nil {
			return nil, e
		}
		return pf, nil
	case []interface{}:
		var pros []interface{}
		for _, p := range pro {
			if p != nil {
				p, err := ProfileUnMarshal(p)
				if err != nil {
					return nil, err
				}
				pros = append(pros, p)
			}
		}
		return pros, nil
	default:
		return nil, errors.New("profile type error")
	}
}
