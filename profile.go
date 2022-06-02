package topotypes

import (
	"encoding/json"
	"errors"
)

type WedgeFaceLimit struct {
	XMin float64 `json:"xmin"`
	ZMin float64 `json:"zmin"`
	XMax float64 `json:"xmax"`
	ZMax float64 `json:"zmax"`
}

type TopoProfile interface {
}

type TopoTriangle struct {
	TopoProfile
	Type string     `json:"type"`
	P1   [3]float64 `json:"p1"`
	P2   [3]float64 `json:"p2"`
	P3   [3]float64 `json:"p3"`
}

func NewTopoTriangle() *TopoTriangle {
	t := TopoTriangle{}
	t.Type = ProfileTypeToString(TOPO_PROFILE_TYPE_TRIANGLE)
	return &t
}

type TopoRectangle struct {
	TopoProfile
	Type string     `json:"type"`
	P1   [3]float64 `json:"p1"`
	P2   [3]float64 `json:"p2"`
}

func NewTopoRectangle() *TopoRectangle {
	t := TopoRectangle{}
	t.Type = ProfileTypeToString(TOPO_PROFILE_TYPE_RECTANGLE)
	return &t
}

type TopoCirc struct {
	TopoProfile
	Type   string     `json:"type"`
	Center [3]float64 `json:"center"`
	Norm   [3]float64 `json:"norm"`
	Radius float64    `json:"radius"`
}

func NewTopoCirc() *TopoCirc {
	t := TopoCirc{}
	t.Type = ProfileTypeToString(TOPO_PROFILE_TYPE_CIRC)
	return &t
}

type TopoElips struct {
	TopoProfile
	Type   string     `json:"type"`
	S1     [3]float64 `json:"s1"`
	S2     [3]float64 `json:"s2"`
	Center [3]float64 `json:"center"`
}

func NewTopoElips() *TopoElips {
	t := TopoElips{}
	t.Type = ProfileTypeToString(TOPO_PROFILE_TYPE_ELIPS)
	return &t
}

type TopoPolygon struct {
	TopoProfile
	Type  string       `json:"type"`
	Edges [][3]float64 `json:"edges,omitempty"`
}

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
		default:
			return nil, errors.New("profile type error")
		}
		e := json.Unmarshal(([]byte)(js), pf)
		if e != nil {
			return nil, e
		}
		return pf, nil
	default:
		return nil, errors.New("profile type error")
	}
}
