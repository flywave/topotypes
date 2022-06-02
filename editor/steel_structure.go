package editor

import (
	"github.com/flywave/topotypes/material"
	"github.com/flywave/topotypes/profile"
)

type Structure struct {
	Wire      [][3]float64         `json:"wire"`
	Profile   profile.Profile      `json:"profile"`
	Materials []*material.Material `json:"materials,omitempty"`
}

type SteelStructure struct {
	BaseComponent
	Structures []Structure `json:"structures"`
}
