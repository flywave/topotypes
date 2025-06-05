package topotypes

import (
	"encoding/json"
	"strings"

	"github.com/flywave/topotypes/base"
	"github.com/flywave/topotypes/gim"
	"github.com/flywave/topotypes/hydropower"
	"github.com/flywave/topotypes/topo4d"
)

type TopoParametricInterface interface {
	TopoMaterialModelInterface
	IsInstance() bool
	GetGenerate() *topo4d.Generate4D
	GetTopoWork() *topo4d.TopoWork
	SetTopoWork(*topo4d.TopoWork)
	GetShape() ParametricShape
	SetShape(ParametricShape)
}

type ParametricShape interface {
	GetType() string
}

type TopoParametric struct {
	Topos
	Materials  map[string]*TopoMaterial `json:"materials,omitempty"`
	MaterialId string                   `json:"mtl_id,omitempty"`
	Generate   *topo4d.Generate4D       `json:"generate,omitempty"`
	Work       *topo4d.TopoWork         `json:"works,omitempty"`
	Shape      ParametricShape          `json:"shape,omitempty"`
}

func NewTopoParametric() *TopoParametric {
	return &TopoParametric{
		Topos: Topos{
			Type: TopoTypeToString(TOPO_TYPE_PARAMETRIC),
		},
		Materials: map[string]*TopoMaterial{},
	}
}

func (t *TopoParametric) IsInstance() bool {
	return t.Instanced
}

func (t *TopoParametric) GetShape() ParametricShape {
	return t.Shape
}

func (t *TopoParametric) GetGenerate() *topo4d.Generate4D {
	return t.Generate
}

func (t *TopoParametric) GetMaterials() map[string]*TopoMaterial {
	if t.Materials == nil {
		return map[string]*TopoMaterial{}
	}
	return (map[string]*TopoMaterial)(t.Materials)
}

func (t *TopoParametric) GetMaterialIds() []string {
	if t.MaterialId == "" {
		t.MaterialId = "mtl_0"
	}
	return []string{t.MaterialId}
}

func (tp *TopoParametric) IsTopoBound() bool {
	return true
}

func (t *TopoParametric) SetShape(shape ParametricShape) {
	t.Shape = shape
}

func (t *TopoParametric) GetTopoWork() *topo4d.TopoWork {
	return t.Work
}
func (t *TopoParametric) SetTopoWork(work *topo4d.TopoWork) {
	t.Work = work
}

func (t *TopoParametric) UnmarshalJSON(data []byte) error {
	stu := struct {
		Topos
		Materials  map[string]*TopoMaterial `json:"materials,omitempty"`
		MaterialId string                   `json:"mtl_id,omitempty"`
		Shape      map[string]interface{}   `json:"shape,omitempty"`
	}{}

	if err := json.Unmarshal(data, &stu); err != nil {
		return err
	}
	t.Topos = stu.Topos
	t.Materials = stu.Materials
	t.MaterialId = stu.MaterialId

	ty := stu.Shape["type"].(string)
	tys := strings.Split(ty, "/")
	dt, _ := json.Marshal(stu.Shape)
	if len(tys) > 1 {
		if tys[0] == gim.Major {
			shp, err := gim.Unmarshal(ty, dt)
			if err != nil {
				return err
			}
			t.Shape = shp
		} else if tys[0] == hydropower.Major {
			shp, err := hydropower.Unmarshal(ty, dt)
			if err != nil {
				return err
			}
			t.Shape = shp
		}
	} else {
		shp, err := base.Unmarshal(ty, dt)
		if err != nil {
			return err
		}
		t.Shape = shp
	}
	return nil
}
