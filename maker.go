package topotypes

import (
	"encoding/json"
	"strings"

	"github.com/flywave/topotypes/base"
	"github.com/flywave/topotypes/gim"
	"github.com/flywave/topotypes/hydropower"
	"github.com/flywave/topotypes/topo4d"
)

type TopoMakerInterface interface {
	GetMaterials() map[string]*TopoMaterial
	IsInstance() bool
	GetMaterialIds() []string
	GetShape() MakerShape
	SetShape(MakerShape)
}

type MakerShape interface {
	GetType() string
}

type TopoMaker struct {
	Topos
	Materials  map[string]*TopoMaterial `json:"materials,omitempty"`
	MaterialId string                   `json:"mtl_id,omitempty"`
	Generate   *topo4d.Generate         `json:"generate,omitempty"`
	Shape      MakerShape               `json:"shape,omitempty"`
}

func (t *TopoMaker) IsInstance() bool {
	return t.Instanced
}

func (t *TopoMaker) GetShape() MakerShape {
	return t.Shape
}

func (t *TopoMaker) GetGenerate() *topo4d.Generate {
	return t.Generate
}

func (t *TopoMaker) GetMaterials() map[string]*TopoMaterial {
	if t.Materials == nil {
		return map[string]*TopoMaterial{}
	}
	return (map[string]*TopoMaterial)(t.Materials)
}

func (t *TopoMaker) GetMaterialIds() []string {
	if t.MaterialId == "" {
		t.MaterialId = "mtl_0"
	}
	return []string{t.MaterialId}
}

func (tp *TopoMaker) IsTopoBound() bool {
	return true
}

func (t *TopoMaker) SetShape(shape MakerShape) {
	t.Shape = shape
}

func (t *TopoMaker) UnmarshalJSON(data []byte) error {
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
