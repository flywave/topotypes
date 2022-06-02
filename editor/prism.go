package editor

import (
	"github.com/flywave/topotypes/material"
	"github.com/flywave/topotypes/profile"
)

type Prism struct {
	BaseComponent
	Profile      profile.Profile      `json:"profile,omitempty"`
	UntilProfile profile.Profile      `json:"until_profile,omitempty"`
	Direction    [3]float64           `json:"direction"`
	Materials    []*material.Material `json:"materials,omitempty"`
}
