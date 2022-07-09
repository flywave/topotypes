package component

import (
	"encoding/json"

	"github.com/flywave/topotypes/material"
	"github.com/flywave/topotypes/shape"
)

type Shape struct {
	BaseComponent
	Shape      string               `json:"-"`
	ShapeModel shape.Shape          `json:"shape"`
	Materials  []*material.Material `json:"materials,omitempty"`
}

func ShapeUnMarshal(js []byte) (*Shape, error) {
	r := Shape{}
	e := json.Unmarshal(js, &r)
	if e != nil {
		return nil, e
	}
	if r.ShapeModel != nil {
		shp, tp, er := shape.ShapeUnMarshal(r.ShapeModel)
		if er != nil {
			return nil, er
		}
		r.ShapeModel = shp
		r.Shape = tp
	}
	return &r, nil
}
