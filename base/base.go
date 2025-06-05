package base

import (
	"encoding/json"
	"fmt"

	"github.com/flywave/topotypes/anchor"
	"github.com/flywave/topotypes/profile"
)

type Base struct {
	Version int    `json:"version"`
	Type    string `json:"type"`
}

func (b *Base) GetType() string {
	return b.Type
}

// Profile types
type TriangleProfile struct {
	Type string     `json:"type"` // "TRIANGLE"
	P1   [3]float64 `json:"p1"`
	P2   [3]float64 `json:"p2"`
	P3   [3]float64 `json:"p3"`
}

type RectangleProfile struct {
	Type string     `json:"type"` // "RECTANGLE"
	P1   [3]float64 `json:"p1"`
	P2   [3]float64 `json:"p2"`
}

type CircProfile struct {
	Type   string     `json:"type"` // "CIRC"
	Center [3]float64 `json:"center"`
	Norm   [3]float64 `json:"norm"`
	Radius float64    `json:"radius"`
}

type ElipsProfile struct {
	Type   string     `json:"type"` // "ELIPS"
	S1     [3]float64 `json:"s1"`
	S2     [3]float64 `json:"s2"`
	Center [3]float64 `json:"center"`
}

type PolygonProfile struct {
	Type   string         `json:"type"` // "POLYGON"
	Edges  [][3]float64   `json:"edges"`
	Inners [][][3]float64 `json:"inners,omitempty"`
}

type SegmentType string

const (
	SegmentTypeLine            SegmentType = "LINE"
	SegmentTypeThreePointArc   SegmentType = "THREE_POINT_ARC"
	SegmentTypeCircleCenterArc SegmentType = "CIRCLE_CENTER_ARC"
	SegmentTypeSpline          SegmentType = "SPLINE"
	SegmentTypeBezier          SegmentType = "BEZIER"
)

func (s SegmentType) ToInt() int {
	switch s {
	case SegmentTypeLine:
		return 1
	case SegmentTypeThreePointArc:
		return 2
	case SegmentTypeCircleCenterArc:
		return 3
	case SegmentTypeSpline:
		return 4
	case SegmentTypeBezier:
		return 5
	default:
		return 1
	}
}

// TransitionMode enum
type TransitionMode string

const (
	TransitionModeTransformed TransitionMode = "TRANSFORMED"
	TransitionModeRound       TransitionMode = "ROUND"
	TransitionModeRight       TransitionMode = "RIGHT"
)

func (t TransitionMode) ToInt() int {
	switch t {
	case TransitionModeTransformed:
		return 1
	case TransitionModeRound:
		return 2
	case TransitionModeRight:
		return 3
	default:
		return 1
	}
}

type PipeJointMode string

const (
	PipeJointModeSphere   PipeJointMode = "SPHERE"
	PipeJointModeBox      PipeJointMode = "BOX"
	PipeJointModeCylinder PipeJointMode = "CYLINDER"
)

func (p PipeJointMode) ToInt() int {
	switch p {
	case PipeJointModeSphere:
		return 0
	case PipeJointModeBox:
		return 1
	case PipeJointModeCylinder:
		return 2
	default:
		return 0
	}
}

type Axis struct {
	Location  [3]float64 `json:"location"`
	Direction [3]float64 `json:"direction"`
}

// Revol represents a revolved object
type Revol struct {
	Base
	Profile profile.Profile `json:"profile"`
	Axis    *Axis           `json:"axis"`
	Angle   float64         `json:"angle"`
}

func NewRevol() *Revol {
	return &Revol{
		Base: Base{Type: "Revol"},
	}
}

// Prism represents a prism object
type Prism struct {
	Base
	Profile   profile.Profile `json:"profile"`
	Direction [3]float64      `json:"direction"`
	Height    float64         `json:"height"`
}

func NewPrism() *Prism {
	return &Prism{
		Base: Base{Type: "Prism"},
	}
}

// Pipe represents a pipe object
type Pipe struct {
	Base
	Wire           [][3]float64          `json:"-,omitempty"`
	Profile        [2]profile.Profile    `json:"profile"`
	InnerProfile   *[2]profile.Profile   `json:"innerProfile,omitempty"`
	SegmentType    SegmentType           `json:"segmentType"`
	Anchors        [2]*anchor.TopoAnchor `json:"anchors"`
	TransitionMode TransitionMode        `json:"transitionMode"`
	UpDir          *[3]float64           `json:"upDir,omitempty"`
}

func NewPipe() *Pipe {
	return &Pipe{
		Base: Base{Type: "Pipe"},
	}
}

func (m *Pipe) GetAnchors() [2]*anchor.TopoAnchor {
	return m.Anchors
}

type SegmentIndex struct {
	Start int         `json:"-"`
	End   int         `json:"end"`
	Type  SegmentType `json:"type"`
}

// MultiSegmentPipePrimitive represents a multi-segment pipe
type MultiSegmentPipe struct {
	Base
	Wires          [][][3]float64        `json:"-,omitempty"`
	Profiles       []profile.Profile     `json:"profiles"`
	InnerProfiles  []profile.Profile     `json:"innerProfiles,omitempty"`
	SegmentTypes   []SegmentType         `json:"-"`
	TransitionMode TransitionMode        `json:"transitionMode"`
	UpDir          *[3]float64           `json:"upDir,omitempty"`
	Anchors        [2]*anchor.TopoAnchor `json:"anchors"`
	SegmentIndexs  []SegmentIndex        `json:"segmentIndexs"`
}

func NewMultiSegmentPipe() *MultiSegmentPipe {
	return &MultiSegmentPipe{
		Base: Base{Type: "MultiSegmentPipe"},
	}
}

func (m *MultiSegmentPipe) GetAnchors() [2]*anchor.TopoAnchor {
	return m.Anchors
}

type ProfileLayer struct {
	Name          string                 `json:"name,omitempty"`
	Profiles      []profile.Profile      `json:"profiles"`
	InnerProfiles []profile.Profile      `json:"innerProfiles,omitempty"`
	Property      map[string]interface{} `json:"property,omitempty"`
	MTL           string                 `json:"mtl,omitempty"`
}

type MultiLayerExtrusionStructure struct {
	Base
	Wires          [][][3]float64  `json:"-,omitempty"`
	SegmentTypes   []SegmentType   `json:"-"`
	TransitionMode TransitionMode  `json:"transitionMode"`
	Layers         []*ProfileLayer `json:"layers,omitempty"`
	UpDir          *[3]float64     `json:"upDir,omitempty"`
	SegmentIndexs  []SegmentIndex  `json:"segmentIndexs"`
}

func NewMultiLayerExtrusionStructure() *MultiLayerExtrusionStructure {
	return &MultiLayerExtrusionStructure{
		Base: Base{Type: "MultiLayerExtrusionStructure"},
	}
}

// PipeJointEndpoint represents a pipe joint endpoint
type PipeJointEndpoint struct {
	ID           string          `json:"id"`
	Offset       *[3]float64     `json:"offset,omitempty"`
	Normal       *[3]float64     `json:"normal,omitempty"`
	Profile      profile.Profile `json:"profile,omitempty"`
	InnerProfile profile.Profile `json:"innerProfile,omitempty"`
}

// PipeJoint represents a pipe joint
type PipeJoint struct {
	Base
	Ins          []string            `json:"ins"`
	Outs         []string            `json:"outs"`
	InsEndpoint  []PipeJointEndpoint `json:"-"`
	OutsEndpoint []PipeJointEndpoint `json:"-"`
	Mode         PipeJointMode       `json:"mode"`
	Flanged      bool                `json:"flanged"`
	UpDir        *[3]float64         `json:"upDir,omitempty"`
}

func NewPipeJoint() *PipeJoint {
	return &PipeJoint{
		Base: Base{Type: "PipeJoint"},
	}
}

// Catenary represents a catenary object
type Catenary struct {
	Base
	P1           *[3]float64     `json:"-"`
	P2           *[3]float64     `json:"-"`
	Profile      profile.Profile `json:"profile"`
	Slack        float64         `json:"slack"`
	MaxSag       float64         `json:"maxSag"`
	Tessellation int             `json:"tessellation"`
	UpDir        *[3]float64     `json:"upDir,omitempty"`
}

func NewCatenary() *Catenary {
	return &Catenary{
		Base: Base{Type: "Catenary"},
	}
}

// Shape objects
type BoxShape struct {
	Base
	Point1 [3]float64 `json:"point1"`
	Point2 [3]float64 `json:"point2"`
}

func NewBoxShape() *BoxShape {
	return &BoxShape{
		Base: Base{Type: "BoxShape"},
	}
}

type ConeShape struct {
	Base
	Radius1 float64  `json:"radius1"`
	Radius2 float64  `json:"radius2"`
	Height  float64  `json:"height"`
	Angle   *float64 `json:"angle,omitempty"`
}

func NewConeShape() *ConeShape {
	return &ConeShape{
		Base: Base{Type: "ConeShape"},
	}
}

type CylinderShape struct {
	Base
	Radius float64  `json:"radius"`
	Height float64  `json:"height"`
	Angle  *float64 `json:"angle,omitempty"`
}

func NewCylinderShape() *CylinderShape {
	return &CylinderShape{
		Base: Base{Type: "CylinderShape"},
	}
}

type RevolutionShape struct {
	Base
	Meridian [][3]float64 `json:"meridian"`
	Angle    *float64     `json:"angle,omitempty"`
	Max      *float64     `json:"max,omitempty"`
	Min      *float64     `json:"min,omitempty"`
}

func NewRevolutionShape() *RevolutionShape {
	return &RevolutionShape{
		Base: Base{Type: "RevolutionShape"},
	}
}

type SphereShape struct {
	Base
	Center *[3]float64 `json:"center,omitempty"`
	Radius float64     `json:"radius"`
	Angle1 *float64    `json:"angle1,omitempty"`
	Angle2 *float64    `json:"angle2,omitempty"`
	Angle  *float64    `json:"angle,omitempty"`
}

func NewSphereShape() *SphereShape {
	return &SphereShape{
		Base: Base{Type: "SphereShape"},
	}
}

type TorusShape struct {
	Base
	Radius1 float64  `json:"radius1"`
	Radius2 float64  `json:"radius2"`
	Angle1  *float64 `json:"angle1,omitempty"`
	Angle2  *float64 `json:"angle2,omitempty"`
	Angle   *float64 `json:"angle,omitempty"`
}

func NewTorusShape() *TorusShape {
	return &TorusShape{
		Base: Base{Type: "TorusShape"},
	}
}

type WedgeShape struct {
	Base
	Edge  [3]float64  `json:"edge"`
	Limit *[4]float64 `json:"limit,omitempty"`
	Ltx   *float64    `json:"ltx,omitempty"`
}

func NewWedgeShape() *WedgeShape {
	return &WedgeShape{
		Base: Base{Type: "WedgeShape"},
	}
}

type PipeShape struct {
	Base
	Wire    [2][3]float64   `json:"wire"`
	Profile profile.Profile `json:"profile"`
	UpDir   *[3]float64     `json:"upDir,omitempty"`
}

func NewPipeShape() *PipeShape {
	return &PipeShape{
		Base: Base{Type: "PipeShape"},
	}
}

type StepShape struct {
	Base
	Name string `json:"name"`
	Step string `json:"step"`
}

func NewStepShape() *StepShape {
	return &StepShape{
		Base: Base{Type: "StepShape"},
	}
}

type Shape interface {
	GetType() string
}

func Unmarshal(ty string, dt []byte) (Shape, error) {
	switch ty {
	case "Revol":
		rel := &Revol{}
		err := json.Unmarshal(dt, rel)
		if err != nil {
			return nil, err
		}
		if rel.Profile != nil {
			rel.Profile, _ = profile.ProfileUnMarshal(rel.Profile)
		}
		return rel, nil
	case "Prism":
		prism := &Prism{}
		err := json.Unmarshal(dt, prism)
		if err != nil {
			return nil, err
		}
		if prism.Profile != nil {
			prism.Profile, _ = profile.ProfileUnMarshal(prism.Profile)
		}
		return prism, nil
	case "Pipe":
		pipe := &Pipe{}
		err := json.Unmarshal(dt, pipe)
		if err != nil {
			return nil, err
		}

		if pipe.Profile[0] != nil {
			pipe.Profile[0], _ = profile.ProfileUnMarshal(pipe.Profile[0])
		}
		if pipe.Profile[1] != nil {
			pipe.Profile[1], _ = profile.ProfileUnMarshal(pipe.Profile[1])
		}
		if pipe.InnerProfile != nil {
			if pipe.InnerProfile[0] != nil {
				pipe.InnerProfile[0], _ = profile.ProfileUnMarshal(pipe.InnerProfile[0])
			}
			if pipe.InnerProfile[1] != nil {
				pipe.InnerProfile[1], _ = profile.ProfileUnMarshal(pipe.InnerProfile[1])
			}
		}
		return pipe, nil
	case "MultiSegmentPipe":
		pipe := &MultiSegmentPipe{}
		err := json.Unmarshal(dt, pipe)
		if err != nil {
			return nil, err
		}
		for i := range pipe.Profiles {
			pipe.Profiles[i], _ = profile.ProfileUnMarshal(pipe.Profiles[i])
		}
		for i := range pipe.InnerProfiles {
			pipe.InnerProfiles[i], _ = profile.ProfileUnMarshal(pipe.InnerProfiles[i])
		}
		return pipe, nil
	case "MultiLayerExtrusionStructure":
		extrusion := &MultiLayerExtrusionStructure{}
		err := json.Unmarshal(dt, extrusion)
		if err != nil {
			return nil, err
		}
		for i := range extrusion.Layers {
			for j := range extrusion.Layers[i].Profiles {
				extrusion.Layers[i].Profiles[j], _ = profile.ProfileUnMarshal(extrusion.Layers[i].Profiles[j])
			}
			for j := range extrusion.Layers[i].InnerProfiles {
				extrusion.Layers[i].InnerProfiles[j], _ = profile.ProfileUnMarshal(extrusion.Layers[i].InnerProfiles[j])
			}
		}
		return extrusion, nil
	case "PipeJoint":
		joint := &PipeJoint{}
		err := json.Unmarshal(dt, joint)
		return joint, err
	case "Catenary":
		catenary := &Catenary{}
		err := json.Unmarshal(dt, catenary)
		if err != nil {
			return nil, err
		}
		if catenary.Profile != nil {
			catenary.Profile, _ = profile.ProfileUnMarshal(catenary.Profile)
		}
		return catenary, nil
	case "BoxShape":
		shape := &BoxShape{}
		err := json.Unmarshal(dt, shape)
		return shape, err
	case "ConeShape":
		shape := &ConeShape{}
		err := json.Unmarshal(dt, shape)
		return shape, err
	case "CylinderShape":
		shape := &CylinderShape{}
		err := json.Unmarshal(dt, shape)
		return shape, err
	case "RevolutionShape":
		shape := &RevolutionShape{}
		err := json.Unmarshal(dt, shape)
		return shape, err
	case "SphereShape":
		shape := &SphereShape{}
		err := json.Unmarshal(dt, shape)
		return shape, err
	case "TorusShape":
		shape := &TorusShape{}
		err := json.Unmarshal(dt, shape)
		return shape, err
	case "WedgeShape":
		shape := &WedgeShape{}
		err := json.Unmarshal(dt, shape)
		return shape, err
	case "PipeShape":
		shape := &PipeShape{}
		err := json.Unmarshal(dt, shape)
		if err != nil {
			return nil, err
		}
		if shape.Profile != nil {
			shape.Profile, _ = profile.ProfileUnMarshal(shape.Profile)
		}
		return shape, nil
	case "StepShape":
		shape := &StepShape{}
		err := json.Unmarshal(dt, shape)
		return shape, err
	default:
		return nil, fmt.Errorf("invalid type: %s", ty)
	}
}
