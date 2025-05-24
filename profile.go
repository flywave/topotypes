package topotypes

import (
	"github.com/flywave/topotypes/profile"
)

const (
	TOPO_PROFILE_TYPE_NONE      = profile.TYPE_NONE
	TOPO_PROFILE_TYPE_TRIANGLE  = profile.TYPE_TRIANGLE
	TOPO_PROFILE_TYPE_RECTANGLE = profile.TYPE_RECTANGLE
	TOPO_PROFILE_TYPE_CIRC      = profile.TYPE_CIRC
	TOPO_PROFILE_TYPE_ELIPS     = profile.TYPE_ELIPS
	TOPO_PROFILE_TYPE_POLYGON   = profile.TYPE_POLYGON
	TOPO_PROFILE_TYPE_L_STEEL   = profile.TYPE_L_STEEL
)

type TopoProfile profile.Profile

type TopoTriangle profile.Triangle

func NewTopoTriangle() *TopoTriangle {
	t := TopoTriangle{}
	t.Type = profile.ProfileTypeToString(TOPO_PROFILE_TYPE_TRIANGLE)
	return &t
}

type TopoRectangle profile.Rectangle

func NewTopoRectangle() *TopoRectangle {
	t := TopoRectangle{}
	t.Type = profile.ProfileTypeToString(TOPO_PROFILE_TYPE_RECTANGLE)
	return &t
}

type TopoCirc profile.Circ

func NewTopoCirc() *TopoCirc {
	t := TopoCirc{}
	t.Type = profile.ProfileTypeToString(TOPO_PROFILE_TYPE_CIRC)
	return &t
}

type TopoElips profile.Elips

func NewTopoElips() *TopoElips {
	t := TopoElips{}
	t.Type = profile.ProfileTypeToString(TOPO_PROFILE_TYPE_ELIPS)
	return &t
}

type TopoPolygon profile.Polygon

func NewTopoPolygon() *TopoPolygon {
	t := TopoPolygon{}
	t.Type = profile.ProfileTypeToString(TOPO_PROFILE_TYPE_POLYGON)
	return &t
}

func ProfileUnMarshal(inter interface{}) (interface{}, error) {
	return profile.ProfileUnMarshal(inter)
}
