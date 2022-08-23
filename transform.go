package topotypes

import (
	dmat4 "github.com/flywave/go3d/float64/mat4"
	"github.com/flywave/go3d/float64/quaternion"
	dvec3 "github.com/flywave/go3d/float64/vec3"
)

type TopoTransform struct {
	Rotation  *[4]float64 `json:"rotation"`
	Translate *[3]float64 `json:"translate"`
	Scale     *[3]float64 `json:"scale"`
}

func NewTopoTransform() *TopoTransform {
	return &TopoTransform{Scale: &[3]float64{1.0, 1.0, 1.0}}
}

func (t *TopoTransform) Compose() *dmat4.T {
	if t.Translate == nil {
		t.Translate = &[3]float64{}
	}
	if t.Scale == nil {
		t.Scale = &[3]float64{1, 1, 1}
	}
	if t.Rotation == nil {
		t.Rotation = &[4]float64{0, 0, 0, 1}
	}
	return dmat4.Compose((*dvec3.T)(t.Translate), (*quaternion.T)(t.Rotation), (*dvec3.T)(t.Scale))
}
