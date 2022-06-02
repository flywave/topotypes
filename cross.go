package topotypes

import (
	"github.com/flywave/go3d/float64/vec3"
	"github.com/flywave/go3d/quaternion"
)

type CrossTopology struct {
	Scale       float64      `json:"scale"`
	Rotation    quaternion.T `json:"rotation"`
	Offset      vec3.T       `json:"offset"`
	Anchors     []*Anchor    `json:"anchors"`
	AnchorCount int          `json:"anchorcount"`
}

type TopoCrossMultiPoint struct {
	Topos
	Refs    []TopoAnchorRef   `json:"links,omitempty"`
	Objects []*TopoCrossPoint `json:"objects,omitempty"`
}

func NewTopoCrossMultiPoint() *TopoCrossMultiPoint {
	t := &TopoCrossMultiPoint{}
	t.Type = TopoTypeToString(TOPO_TYPE_CROSS_MULTI_POINT)
	return t
}

type TopoCrossPoint struct {
	Topos
	Model string           `json:"model"`
	Links []TopoAnchorLink `json:"links,omitempty"`
}

func NewTopoCrossPoint() *TopoCrossPoint {
	t := &TopoCrossPoint{}
	t.Type = TopoTypeToString(TOPO_TYPE_CROSS_POINT)
	return t
}
