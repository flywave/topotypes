package base

import (
	"encoding/json"
	"fmt"

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

// Segment types
const (
	SegmentTypeLine            = "LINE"
	SegmentTypeThreePointArc   = "THREE_POINT_ARC"
	SegmentTypeCircleCenterArc = "CIRCLE_CENTER_ARC"
	SegmentTypeSpline          = "SPLINE"
)

// Transition modes
const (
	TransitionModeTransformed = "TRANSFORMED"
	TransitionModeRound       = "ROUND"
	TransitionModeRight       = "RIGHT"
)

// Joint modes
const (
	JointModeSphere   = "SPHERE"
	JointModeBox      = "BOX"
	JointModeCylinder = "CYLINDER"
)

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
	Wire           [][3]float64        `json:"wire,omitempty"`
	Profile        [2]profile.Profile  `json:"profile"`
	InnerProfile   *[2]profile.Profile `json:"innerProfile,omitempty"`
	SegmentType    string              `json:"segmentType"`
	TransitionMode string              `json:"transitionMode"`
	UpDir          *[3]float64         `json:"upDir,omitempty"`
}

func NewPipe() *Pipe {
	return &Pipe{
		Base: Base{Type: "Pipe"},
	}
}

// MultiSegmentPipePrimitive represents a multi-segment pipe
type MultiSegmentPipePrimitive struct {
	Base
	Wires          [][][3]float64    `json:"wires,omitempty"`
	Profiles       []profile.Profile `json:"profiles"`
	InnerProfiles  []profile.Profile `json:"innerProfiles,omitempty"`
	SegmentTypes   []string          `json:"segmentTypes"`
	TransitionMode string            `json:"transitionMode"`
	UpDir          *[3]float64       `json:"upDir,omitempty"`
}

func NewMultiSegmentPipePrimitive() *MultiSegmentPipePrimitive {
	return &MultiSegmentPipePrimitive{
		Base: Base{Type: "MultiSegmentPipe"},
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
	Ins     []PipeJointEndpoint `json:"ins"`
	Outs    []PipeJointEndpoint `json:"outs"`
	Mode    string              `json:"mode"`
	Flanged bool                `json:"flanged"`
	UpDir   *[3]float64         `json:"upDir,omitempty"`
}

func NewPipeJoint() *PipeJoint {
	return &PipeJoint{
		Base: Base{Type: "PipeJoint"},
	}
}

// Catenary represents a catenary object
type Catenary struct {
	Base
	P1           *[3]float64     `json:"p1,omitempty"`
	P2           *[3]float64     `json:"p2,omitempty"`
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
	Radius1 float64 `json:"radius1"`
	Radius2 float64 `json:"radius2"`
	Height  float64 `json:"height"`
	Angle   float64 `json:"angle,omitempty"`
}

func NewConeShape() *ConeShape {
	return &ConeShape{
		Base: Base{Type: "ConeShape"},
	}
}

type CylinderShape struct {
	Base
	Radius float64 `json:"radius"`
	Height float64 `json:"height"`
	Angle  float64 `json:"angle,omitempty"`
}

func NewCylinderShape() *CylinderShape {
	return &CylinderShape{
		Base: Base{Type: "CylinderShape"},
	}
}

type RevolutionShape struct {
	Base
	Meridian [][3]float64 `json:"meridian"`
	Angle    float64      `json:"angle,omitempty"`
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
	Angle1 float64     `json:"angle1,omitempty"`
	Angle2 float64     `json:"angle2,omitempty"`
	Angle  float64     `json:"angle,omitempty"`
}

func NewSphereShape() *SphereShape {
	return &SphereShape{
		Base: Base{Type: "SphereShape"},
	}
}

type TorusShape struct {
	Base
	Radius1 float64 `json:"radius1"`
	Radius2 float64 `json:"radius2"`
	Angle1  float64 `json:"angle1,omitempty"`
	Angle2  float64 `json:"angle2,omitempty"`
	Angle   float64 `json:"angle,omitempty"`
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
	Ltx   float64     `json:"ltx,omitempty"`
}

func NewWedgeShape() *WedgeShape {
	return &WedgeShape{
		Base: Base{Type: "WedgeShape"},
	}
}

type PipeShape struct {
	Base
	Wire    [][3]float64    `json:"wire"`
	Profile profile.Profile `json:"profile"`
	UpDir   *[3]float64     `json:"upDir,omitempty"`
}

func NewPipeShape() *PipeShape {
	return &PipeShape{
		Base: Base{Type: "PipeShape"},
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
		pipe := &MultiSegmentPipePrimitive{}
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
	default:
		return nil, fmt.Errorf("invalid type: %s", ty)
	}
}
