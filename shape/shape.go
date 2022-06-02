package shape

import "github.com/flywave/topotypes/utils"

const (
	MODE_NONE = iota
	MODE_BOX
	MODE_CYLINDER
	MODE_CONE
	MODE_SPHERE
	MODE_TORUS
	MODE_WEDGE
	MODE_REVOLUTION
	MODE_PIPE
)

func ShapeTypeToString(tp int) string {
	switch tp {
	case MODE_BOX:
		return "box"
	case MODE_CYLINDER:
		return "cylinder"
	case MODE_CONE:
		return "cone"
	case MODE_SPHERE:
		return "sphere"
	case MODE_TORUS:
		return "torus"
	case MODE_WEDGE:
		return "wedge"
	case MODE_REVOLUTION:
		return "revolution"
	case MODE_PIPE:
		return "pipe"
	default:
		return ""
	}
}

func StringToShapeType(tp string) int {
	if utils.StrEquals(tp, "box") {
		return MODE_BOX
	} else if utils.StrEquals(tp, "cylinder") {
		return MODE_CYLINDER
	} else if utils.StrEquals(tp, "cone") {
		return MODE_CONE
	} else if utils.StrEquals(tp, "sphere") {
		return MODE_SPHERE
	} else if utils.StrEquals(tp, "torus") {
		return MODE_TORUS
	} else if utils.StrEquals(tp, "wedge") {
		return MODE_WEDGE
	} else if utils.StrEquals(tp, "revolution") {
		return MODE_REVOLUTION
	} else if utils.StrEquals(tp, "pipe") {
		return MODE_PIPE
	}
	return MODE_NONE
}

type Shape interface {
}

type Box struct {
	Type   string
	Point1 [3]float64 `json:"point1"`
	Point2 [3]float64 `json:"point2"`
}

func NewBox() *Box {
	return &Box{
		Type: ShapeTypeToString(MODE_BOX),
	}
}

type Cone struct {
	Type    string   `json:"type"`
	Radius1 float64  `json:"radius1"`
	Radius2 float64  `json:"radius2"`
	Height  float64  `json:"height"`
	Angle   *float64 `json:"angle,omitempty"`
}

func NewCone() *Cone {
	return &Cone{
		Type: ShapeTypeToString(MODE_CONE),
	}
}

type Cylinder struct {
	Type   string   `json:"type"`
	Radius float64  `json:"radius"`
	Height float64  `json:"height"`
	Angle  *float64 `json:"angle,omitempty"`
}

func NewCylinder() *Cylinder {
	return &Cylinder{
		Type: ShapeTypeToString(MODE_CYLINDER),
	}
}

type Revolution struct {
	Type     string
	Meridian [][3]float64 `json:"meridian"`
	Angle    *float64     `json:"angle,omitempty"`
	Max      *float64     `json:"max,omitempty"`
	Min      *float64     `json:"min,omitempty"`
}

func NewRevolution() *Revolution {
	return &Revolution{
		Type: ShapeTypeToString(MODE_REVOLUTION),
	}
}

type Sphere struct {
	Type   string
	Center *[3]float64 `json:"center,omitempty"`
	Radius float64     `json:"radius"`
	Angle1 *float64    `json:"angle1,omitempty"`
	Angle2 *float64    `json:"angle2,omitempty"`
	Angle  *float64    `json:"angle,omitempty"`
}

func NewSphere() *Sphere {
	return &Sphere{
		Type: ShapeTypeToString(MODE_SPHERE),
	}
}

type Torus struct {
	Type    string
	Radius1 float64  `json:"radius1"`
	Radius2 float64  `json:"radius2"`
	Angle1  *float64 `json:"angle1,omitempty"`
	Angle2  *float64 `json:"angle2,omitempty"`
	Angle   *float64 `json:"angle,omitempty"`
}

func NewTorus() *Torus {
	return &Torus{
		Type: ShapeTypeToString(MODE_TORUS),
	}
}

const (
	TransitionModeRight = "right"
	TransitionModeRound = "round"
	TransitionModeTrans = "trans"
)

func TransitionModeToInt(m string) int {
	switch m {
	case TransitionModeRight:
		return 1
	case TransitionModeRound:
		return 2
	case TransitionModeTrans:
		return 0
	default:
		return 1
	}
}

func TransitionModeToString(m int) string {
	switch m {
	case 1:
		return TransitionModeRight
	case 2:
		return TransitionModeRound
	}
	return TransitionModeTrans
}

type Pipe struct {
	Type           string
	Wire           [][3]float64 `json:"wire"`
	Profile        interface{}  `json:"profile"`
	UntilProfile   interface{}  `json:"until_profile"`
	Smooth         string       `json:"smooth,omitempty"`
	TransitionMode string       `json:"transition_mode"`
}

func NewPipe() *Pipe {
	t := &Pipe{
		Type: ShapeTypeToString(MODE_PIPE),
	}
	return t
}

type WedgeFaceLimit struct {
	XMin float64 `json:"xmin"`
	ZMin float64 `json:"zmin"`
	XMax float64 `json:"xmax"`
	ZMax float64 `json:"zmax"`
}

type Wedge struct {
	Type  string
	Edge  [3]float64      `json:"edge"`
	Limit *WedgeFaceLimit `json:"limit,omitempty"`
	Ltx   *float64        `json:"ltx,omitempty"`
}

func NewWedge() *Wedge {
	return &Wedge{Type: ShapeTypeToString(MODE_WEDGE)}
}
