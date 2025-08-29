package topotypes

import "github.com/flywave/topotypes/material"

type TopoBoard struct {
	Topos
	Size    *[2]float64 `json:"size,omitempty"`
	Depth   *float64    `json:"depth,omitempty"`
	Texture string      `json:"texture"`
}

func NewTopoBoard() *TopoBoard {
	t := &TopoBoard{}
	t.Type = TopoTypeToString(TOPO_TYPE_BOARD)
	return t
}

func (b *TopoBoard) GetMaterials() map[string]*material.Material {
	return map[string]*material.Material{}
}

func (b *TopoBoard) GetMaterialIds() []string {
	mtlids := []string{}
	return mtlids
}
