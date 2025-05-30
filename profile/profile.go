package profile

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/flywave/topotypes/utils"
)

const (
	TYPE_NONE = iota
	TYPE_TRIANGLE
	TYPE_RECTANGLE
	TYPE_CIRC
	TYPE_ELIPS
	TYPE_POLYGON
)

func ProfileTypeToString(tp int) string {
	switch tp {
	case TYPE_TRIANGLE:
		return "TRIANGLE"
	case TYPE_RECTANGLE:
		return "RECTANGLE"
	case TYPE_CIRC:
		return "CIRC"
	case TYPE_ELIPS:
		return "ELIPS"
	case TYPE_POLYGON:
		return "POLYGON"
	default:
		return ""
	}
}

func StringToProfileType(tp string) int {
	tp = strings.ToLower(tp)
	if utils.StrEquals(tp, "triangle") {
		return TYPE_TRIANGLE
	} else if utils.StrEquals(tp, "rectangle") {
		return TYPE_RECTANGLE
	} else if utils.StrEquals(tp, "circ") {
		return TYPE_CIRC
	} else if utils.StrEquals(tp, "ellipse") {
		return TYPE_ELIPS
	} else if utils.StrEquals(tp, "polygon") {
		return TYPE_POLYGON
	}
	return TYPE_NONE
}

type Profile interface {
}

type Triangle struct {
	Profile
	Type string     `json:"type"`
	P1   [3]float64 `json:"p1"`
	P2   [3]float64 `json:"p2"`
	P3   [3]float64 `json:"p3"`
}

func NewTriangle() *Triangle {
	t := Triangle{}
	t.Type = ProfileTypeToString(TYPE_TRIANGLE)
	return &t
}

type Rectangle struct {
	Profile
	Type string     `json:"type"`
	P1   [3]float64 `json:"p1"`
	P2   [3]float64 `json:"p2"`
}

func NewRectangle() *Rectangle {
	t := Rectangle{}
	t.Type = ProfileTypeToString(TYPE_RECTANGLE)
	return &t
}

type Circ struct {
	Profile
	Type   string     `json:"type"`
	Center [3]float64 `json:"center"`
	Norm   [3]float64 `json:"norm"`
	Radius float64    `json:"radius"`
}

func NewCirc() *Circ {
	t := Circ{}
	t.Type = ProfileTypeToString(TYPE_CIRC)
	return &t
}

type Elips struct {
	Profile
	Type   string     `json:"type"`
	S1     [3]float64 `json:"s1"`
	S2     [3]float64 `json:"s2"`
	Center [3]float64 `json:"center"`
}

func NewElips() *Elips {
	t := Elips{}
	t.Type = ProfileTypeToString(TYPE_ELIPS)
	return &t
}

type Polygon struct {
	Profile
	Type  string       `json:"type"`
	Edges [][3]float64 `json:"edges,omitempty"`
}

func NewPolygon() *Polygon {
	t := Polygon{}
	t.Type = ProfileTypeToString(TYPE_POLYGON)
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
		case TYPE_TRIANGLE:
			pf = NewTriangle()
		case TYPE_RECTANGLE:
			pf = NewRectangle()
		case TYPE_CIRC:
			cc := NewCirc()
			if w, ok := pro["width"]; ok {
				cc.Radius = w.(float64)
				return cc, nil
			}
			pf = cc
		case TYPE_ELIPS:
			pf = NewElips()
		case TYPE_POLYGON:
			pf = NewPolygon()
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
