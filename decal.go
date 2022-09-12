package topotypes

import (
	dmat4 "github.com/flywave/go3d/float64/mat4"
	dquat "github.com/flywave/go3d/float64/quaternion"
	dvec3 "github.com/flywave/go3d/float64/vec3"
)

type TopoDecalRef struct {
	Ref string `json:"ref"`
}

type TopoProjector struct {
	Orientation *[3]float64 `json:"orientation"`
	Position    *[3]float64 `json:"position"`
	Size        *[3]float64 `json:"size"`
}

func (p *TopoProjector) GetTransform() *TopoTransform {
	posMat := dmat4.Ident
	pos := dvec3.T(*p.Position)
	posMat.SetTranslation(&pos)

	quatMat := dmat4.Ident
	quat := dquat.FromEulerAngles(p.Orientation[0], p.Orientation[1], p.Orientation[2])
	quatMat.AssignQuaternion(&quat)

	result := dmat4.Ident
	result.AssignMul(&posMat, &quatMat)

	t, q, _ := dmat4.Decompose(&result)

	return &TopoTransform{
		Rotation:  (*[4]float64)(q),
		Translate: (*[3]float64)(t),
		Scale:     &[3]float64{1.0, 1.0, 1.0},
	}
}

type TopoDecal struct {
	Projector TopoProjector  `json:"projector"`
	Type      string         `json:"type"`
	Texture   string         `json:"texture"`
	Refs      []TopoDecalRef `json:"targets,omitempty"`
	Fusion    bool           `json:"fusion"`
}

func (tp *TopoDecal) GetTransform() *TopoTransform {
	return tp.Projector.GetTransform()
}

func (tp *TopoDecal) GetType() string {
	return tp.Type
}

func (tp *TopoDecal) GetFusion() bool {
	return tp.Fusion
}

func (tp *TopoDecal) ResetTransform() {
	tp.Projector = TopoProjector{}
}
