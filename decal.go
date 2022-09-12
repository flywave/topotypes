package topotypes

import (
	dquat "github.com/flywave/go3d/float64/quaternion"
)

type TopoDecalRef struct {
	Ref string `json:"ref"`
}

type TopoProjector struct {
	Orientation *[3]float64 `json:"orientation"`
	Size        *[3]float64 `json:"size"`
}

func (p *TopoProjector) GetTransform() *TopoTransform {
	quat := dquat.FromEulerAngles(p.Orientation[0], p.Orientation[1], p.Orientation[2])

	return &TopoTransform{
		Rotation:  (*[4]float64)(&quat),
		Translate: &[3]float64{0, 0, 0},
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
