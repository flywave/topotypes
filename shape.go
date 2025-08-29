package topotypes

import (
	"encoding/json"
	"errors"

	"github.com/flywave/topotypes/shape"
)

type TopoBound interface {
	IsTopoBound() bool
}

type TopoShape struct {
	TopoParametric
	Shape      string      `json:"-"`
	ShapeModel interface{} `json:"shape"`
}

func (t *TopoShape) UnmarshalJSON(data []byte) error {
	stu := struct {
		Topos
		Materials  TopoMaterialMap `json:"materials,omitempty"`
		MaterialId string          `json:"mtl_id,omitempty"`
		ShapeModel interface{}     `json:"shape"`
	}{}

	if err := json.Unmarshal(data, &stu); err != nil {
		return err
	}
	t.Topos = stu.Topos
	t.Materials = stu.Materials
	t.MaterialId = stu.MaterialId
	t.ShapeModel = stu.ShapeModel
	return nil
}

func (sp *TopoShape) IsTopoBound() bool {
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

func TopoShapePipeUnmarshal(js []byte) (*TopoShapePipe, error) {
	pipe := TopoShapePipe{}
	e := json.Unmarshal(js, &pipe)
	if e != nil {
		return nil, e
	}
	if pipe.Profile != nil {
		prof, er := ProfileUnmarshal(pipe.Profile)
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

func ShapeUnmarshal(js []byte) (ToposInterface, error) {
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
		if inter, err = TopoShapePipeUnmarshal(js); err != nil {
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
