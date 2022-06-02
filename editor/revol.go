package editor

import (
	"github.com/flywave/topotypes/material"
	"github.com/flywave/topotypes/profile"
)

type Revol struct {
	BaseComponent
	Profile   profile.Profile      `json:"profile,omitempty"`
	Axis      [2][3]float64        `json:"axis"`
	Angle     float64              `json:"angle"`
	Materials []*material.Material `json:"materials,omitempty"`
}
