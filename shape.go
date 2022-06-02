package topotypes

import (
	"encoding/json"
	"errors"
)

type TopoBoundy interface {
	IsTopoBoundy() bool
}

type TopoShape struct {
	TopoMaker
	Shape      string      `json:"-"`
	ShapeModel interface{} `json:"shape"`
}

func (sp *TopoShape) IsTopoBoundy() bool {
	return true
}

type TopoShapeBox struct {
	Type   string
	Point1 [3]float64 `json:"point1"`
	Point2 [3]float64 `json:"point2"`
}

func NewTopoShapeBox() *TopoShapeBox {
	return &TopoShapeBox{
		Type: TopoTypeToString(TOPO_SHAPE_MODE_BOX),
	}
}

type TopoShapeCone struct {
	Type    string   `json:"type"`
	Radius1 float64  `json:"radius1"`
	Radius2 float64  `json:"radius2"`
	Height  float64  `json:"height"`
	Angle   *float64 `json:"angle,omitempty"`
}

func NewTopoShapeCone() *TopoShapeCone {
	return &TopoShapeCone{
		Type: TopoTypeToString(TOPO_SHAPE_MODE_CONE),
	}
}

type TopoShapeCylinder struct {
	Type   string   `json:"type"`
	Radius float64  `json:"radius"`
	Height float64  `json:"height"`
	Angle  *float64 `json:"angle,omitempty"`
}

func NewTopoShapeCylinder() *TopoShapeCylinder {
	return &TopoShapeCylinder{
		Type: TopoTypeToString(TOPO_SHAPE_MODE_CYLINDER),
	}
}

type TopoShapeRevolution struct {
	Type     string
	Meridian [][3]float64 `json:"meridian"`
	Angle    *float64     `json:"angle,omitempty"`
	Max      *float64     `json:"max,omitempty"`
	Min      *float64     `json:"min,omitempty"`
}

func NewTopoShapeRevolution() *TopoShapeRevolution {
	return &TopoShapeRevolution{
		Type: TopoTypeToString(TOPO_SHAPE_MODE_REVOLUTION),
	}
}

type TopoShapeSphere struct {
	Type   string
	Center *[3]float64 `json:"center,omitempty"`
	Radius float64     `json:"radius"`
	Angle1 *float64    `json:"angle1,omitempty"`
	Angle2 *float64    `json:"angle2,omitempty"`
	Angle  *float64    `json:"angle,omitempty"`
}

func NewTopoShapeSphere() *TopoShapeSphere {
	return &TopoShapeSphere{
		Type: TopoTypeToString(TOPO_SHAPE_MODE_SPHERE),
	}
}

type TopoShapeTorus struct {
	Type    string
	Radius1 float64  `json:"radius1"`
	Radius2 float64  `json:"radius2"`
	Angle1  *float64 `json:"angle1,omitempty"`
	Angle2  *float64 `json:"angle2,omitempty"`
	Angle   *float64 `json:"angle,omitempty"`
}

func NewTopoShapeTorus() *TopoShapeTorus {
	return &TopoShapeTorus{
		Type: TopoTypeToString(TOPO_SHAPE_MODE_TORUS),
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

type TopoShapePipe struct {
	TopoPipe
	Shape string
}

func NewTopoShapePipe() *TopoShapePipe {
	t := &TopoShapePipe{
		Shape: ShapeTypeToString(TOPO_SHAPE_MODE_PIPE),
	}
	t.Type = TopoTypeToString(TOPO_TYPE_SHAPE)
	return t
}

type TopoShapeWedge struct {
	Type  string
	Edge  [3]float64      `json:"edge"`
	Limit *WedgeFaceLimit `json:"limit,omitempty"`
	Ltx   *float64        `json:"ltx,omitempty"`
}

func NewTopoShapeWedge() *TopoShapeWedge {
	return &TopoShapeWedge{Type: TopoTypeToString(TOPO_TYPE_SHAPE)}
}

func ShapeUnMarshal(js []byte) (ToposInterface, error) {
	sp := make(map[string]interface{})
	e := json.Unmarshal(js, &sp)
	if e != nil {
		return nil, e
	}
	t, ok := sp["shape"]
	if !ok {
		return nil, errors.New("invalid json")
	}
	bt, _ := json.Marshal(t)
	e = json.Unmarshal(bt, &sp)
	if e != nil {
		return nil, e
	}
	t, ok = sp["type"]
	if !ok {
		return nil, errors.New("invalid json")
	}
	var inter interface{}
	res := &TopoShape{}
	res.Shape = t.(string)
	ty := StringToShapeType(res.Shape)
	switch ty {
	case TOPO_SHAPE_MODE_BOX:
		inter = NewTopoShapeBox()
	case TOPO_SHAPE_MODE_CYLINDER:
		inter = NewTopoShapeCylinder()
	case TOPO_SHAPE_MODE_CONE:
		inter = NewTopoShapeCone()
	case TOPO_SHAPE_MODE_SPHERE:
		inter = NewTopoShapeSphere()
	case TOPO_SHAPE_MODE_TORUS:
		inter = NewTopoShapeTorus()
	case TOPO_SHAPE_MODE_WEDGE:
		inter = NewTopoShapeWedge()
	case TOPO_SHAPE_MODE_REVOLUTION:
		inter = NewTopoShapeRevolution()
	case TOPO_SHAPE_MODE_PIPE:
		p := NewTopoShapePipe()
		tp, e := PipeUnMarshal(js)
		if e != nil {
			return nil, e
		}
		p.TopoPipe = *tp
		return p, nil
	default:
		return nil, errors.New("not support topo type")
	}

	res.ShapeModel = inter
	e = json.Unmarshal(([]byte)(js), res)
	if e != nil {
		return nil, e
	}
	return res, nil
}
