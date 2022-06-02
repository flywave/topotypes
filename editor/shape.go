package editor

import (
	"github.com/flywave/topotypes/material"
	"github.com/flywave/topotypes/shape"
)

type Shape struct {
	BaseComponent
	Shape      string               `json:"-"`
	ShapeModel shape.Shape          `json:"shape"`
	Materials  []*material.Material `json:"materials,omitempty"`
}
