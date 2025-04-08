package topotypes

import (
	"encoding/json"
	"errors"

	"github.com/flywave/topotypes/shape"
)

type TopoBoundy interface {
	IsTopoBoundy() bool
}

type TopoShape struct {
	TopoMaker
	Shape      string      `json:"-"`
	ShapeModel interface{} `json:"shape"`
}

func (sp *TopoShape) GetModel() string {
	return ""
}

func (sp *TopoShape) IsTopoBoundy() bool {
	return true
}

type TopoShapeBox shape.Box

func NewTopoShapeBox() *TopoShapeBox {
	return &TopoShapeBox{
		Type: ShapeTypeToString(TOPO_SHAPE_MODE_BOX),
	}
}

type TopoShapeCone shape.Cone

func NewTopoShapeCone() *TopoShapeCone {
	return &TopoShapeCone{
		Type: ShapeTypeToString(TOPO_SHAPE_MODE_CONE),
	}
}

type TopoShapeCylinder shape.Cylinder

func NewTopoShapeCylinder() *TopoShapeCylinder {
	return &TopoShapeCylinder{
		Type: ShapeTypeToString(TOPO_SHAPE_MODE_CYLINDER),
	}
}

type TopoShapeRevolution shape.Revolution

func NewTopoShapeRevolution() *TopoShapeRevolution {
	return &TopoShapeRevolution{
		Type: ShapeTypeToString(TOPO_SHAPE_MODE_REVOLUTION),
	}
}

type TopoShapeSphere shape.Sphere

func NewTopoShapeSphere() *TopoShapeSphere {
	return &TopoShapeSphere{
		Type: ShapeTypeToString(TOPO_SHAPE_MODE_SPHERE),
	}
}

type TopoShapeTorus shape.Torus

func NewTopoShapeTorus() *TopoShapeTorus {
	return &TopoShapeTorus{
		Type: ShapeTypeToString(TOPO_SHAPE_MODE_TORUS),
	}
}

const (
	TransitionModeRight = shape.TransitionModeRight
	TransitionModeRound = shape.TransitionModeRound
	TransitionModeTrans = shape.TransitionModeTrans
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

type TopoShapePipe shape.Pipe

func NewTopoShapePipe() *TopoShapePipe {
	t := &TopoShapePipe{
		Type: ShapeTypeToString(TOPO_SHAPE_MODE_PIPE),
	}
	return t
}

func TopoShapePipeUnMarshal(js []byte) (*TopoShapePipe, error) {
	pipe := TopoShapePipe{}
	e := json.Unmarshal(js, &pipe)
	if e != nil {
		return nil, e
	}
	if pipe.Profile != nil {
		prof, er := ProfileUnMarshal(pipe.Profile)
		if er != nil {
			return nil, er
		}
		pipe.Profile = prof
	}
	return &pipe, nil
}

type TopoShapeWedge shape.Wedge

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
		var err error
		if inter, err = TopoShapePipeUnMarshal(js); err != nil {
			return nil, err
		}
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
