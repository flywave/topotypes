package topotypes

import (
	mat4d "github.com/flywave/go3d/float64/mat4"
	quatd "github.com/flywave/go3d/float64/quaternion"
	vec3d "github.com/flywave/go3d/float64/vec3"
)

type TopoTransform struct {
	Rotation  *[4]float64 `json:"rotation"`
	Translate *[3]float64 `json:"translate"`
	Scale     *[3]float64 `json:"scale"`
}

func NewTopoTransform() *TopoTransform {
	return &TopoTransform{Scale: &[3]float64{1.0, 1.0, 1.0}}
}

func (t *TopoTransform) Compose() *mat4d.T {
	if t.Translate == nil {
		t.Translate = &[3]float64{}
	}
	if t.Scale == nil {
		t.Scale = &[3]float64{1, 1, 1}
	}
	if t.Rotation == nil {
		t.Rotation = &[4]float64{0, 0, 0, 1}
	}
	return mat4d.Compose((*vec3d.T)(t.Translate), (*quatd.T)(t.Rotation), (*vec3d.T)(t.Scale))
}
