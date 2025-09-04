package component

import (
	"encoding/json"

	geom "github.com/flywave/go-geom"

	"github.com/flywave/topotypes/material"
)

type ParametricShape interface {
	GetType() string
}

type Parametric struct {
	BaseComponent
	Materials  map[string]*material.Material `json:"materials,omitempty"`
	MaterialId string                        `json:"mtl_id,omitempty"`
	Geom       geom.Geometry                 `json:"geom,omitempty"`
	Shape      ParametricShape               `json:"shape,omitempty"`
}

func ParametricUnmarshal(js []byte) (*Parametric, error) {
	r := Parametric{}
	e := json.Unmarshal(js, &r)
	if e != nil {
		return nil, e
	}
	return &r, nil
}
