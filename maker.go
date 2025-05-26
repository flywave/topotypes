package topotypes

import (
	"encoding/json"
	"strings"

	"github.com/flywave/topotypes/base"
	"github.com/flywave/topotypes/gim"
	"github.com/flywave/topotypes/hydropower"
	"github.com/flywave/topotypes/material"
)

type TopoMakerInterface interface {
	GetMaterials() map[string]*material.Material
	IsInstance() bool
	GetMaterialIds() []string
	GetShape() interface{}
	SetShape(interface{})
}

type TopoMaker struct {
	Topos
	Materials  map[string]*material.Material `json:"materials,omitempty"`
	MaterialId string                        `json:"mtl_id,omitempty"`
	Shape      interface{}                   `json:"shape,omitempty"`
}

func (t *TopoMaker) IsInstance() bool {
	return t.Instanced
}

func (t *TopoMaker) GetShape() interface{} {
	return t.Shape
}

func (t *TopoMaker) GetMaterials() map[string]*material.Material {
	if t.Materials == nil {
		return map[string]*material.Material{}
	}
	return (map[string]*material.Material)(t.Materials)
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

func (t *TopoMaker) SetShape(shape interface{}) {
	t.Shape = shape
}

func (t *TopoMaker) UnmarshalJSON(data []byte) error {
	stu := struct {
		Topos
		Materials  map[string]*material.Material `json:"materials,omitempty"`
		MaterialId string                        `json:"mtl_id,omitempty"`
		Shape      map[string]interface{}        `json:"shape,omitempty"`
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
	} else {
		shp, err := base.Unmarshal(ty, dt)
		if err != nil {
			return err
		}
		t.Shape = shp
	}
	return nil
}
