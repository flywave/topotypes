package component

import "github.com/flywave/topotypes/material"

type Board struct {
	Decal
	Material *material.Material `json:"material,omitempty"`
}
