package editor

import "github.com/flywave/topotypes/material"

type Edge [][3]float64

type Face []Edge

type Solid struct {
	Faces     []Face               `json:"faces,omitempty"`
	Materials []*material.Material `json:"materials,omitempty"`
}
