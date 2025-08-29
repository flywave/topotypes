package shape

import (
	"encoding/json"
	"errors"

	"github.com/flywave/topotypes/profile"
	"github.com/flywave/topotypes/utils"
)

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
		return "BoxShape"
	case MODE_CYLINDER:
		return "CylinderShape"
	case MODE_CONE:
		return "ConeShape"
	case MODE_SPHERE:
		return "SphereShape"
	case MODE_TORUS:
		return "TorusShape"
	case MODE_WEDGE:
		return "WedgeShape"
	case MODE_REVOLUTION:
		return "RevolutionShape"
	case MODE_PIPE:
		return "PipeShape"
	default:
		return ""
	}
}

func StringToShapeType(tp string) int {
	if utils.StrEquals(tp, "BoxShape") {
		return MODE_BOX
	} else if utils.StrEquals(tp, "CylinderShape") {
		return MODE_CYLINDER
	} else if utils.StrEquals(tp, "ConeShape") {
		return MODE_CONE
	} else if utils.StrEquals(tp, "SphereShape") {
		return MODE_SPHERE
	} else if utils.StrEquals(tp, "TorusShape") {
		return MODE_TORUS
	} else if utils.StrEquals(tp, "WedgeShape") {
		return MODE_WEDGE
	} else if utils.StrEquals(tp, "RevolutionShape") {
		return MODE_REVOLUTION
	} else if utils.StrEquals(tp, "PipeShape") {
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
	SegmentType    string       `json:"segment_type,omitempty"`
	TransitionMode string       `json:"transition_mode"`
}

func NewPipe() *Pipe {
	t := &Pipe{
		Type: ShapeTypeToString(MODE_PIPE),
	}
	return t
}

func PipeUnmarshal(js []byte) (*Pipe, error) {
	pipe := Pipe{}
	e := json.Unmarshal(js, &pipe)
	if e != nil {
		return nil, e
	}
	if pipe.Profile != nil {
		prof, er := profile.ProfileUnmarshal(pipe.Profile)
		if er != nil {
			return nil, er
		}
		pipe.Profile = prof
	}
	return &pipe, nil
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

func ShapeUnmarshal(inter interface{}) (interface{}, string, error) {
	switch pro := inter.(type) {
	case map[string]interface{}:
		v, ok := pro["type"]
		t, ok2 := v.(string)
		if !ok || !ok2 {
			return nil, "", errors.New("profile type error")
		}
		pro_t := StringToShapeType(t)
		js, er := json.Marshal(inter)
		if er != nil {
			return nil, "", er
		}
		switch pro_t {
		case MODE_BOX:
			inter = NewBox()
		case MODE_CYLINDER:
			inter = NewCylinder()
		case MODE_CONE:
			inter = NewCone()
		case MODE_SPHERE:
			inter = NewSphere()
		case MODE_TORUS:
			inter = NewTorus()
		case MODE_WEDGE:
			inter = NewWedge()
		case MODE_REVOLUTION:
			inter = NewRevolution()
		case MODE_PIPE:
			var err error
			if inter, err = PipeUnmarshal(js); err != nil {
				return nil, "", err
			}
			return inter, ShapeTypeToString(pro_t), nil
		default:
			return nil, "", errors.New("profile type error")
		}
		e := json.Unmarshal(([]byte)(js), inter)
		if e != nil {
			return nil, "", e
		}
		return inter, ShapeTypeToString(pro_t), nil
	default:
		return nil, "", errors.New("profile type error")
	}
}
