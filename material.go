package topotypes

import (
	mst "github.com/flywave/go-mst"
	"github.com/flywave/topotypes/material"
)

type TopoMaterial material.Material

func TopoMtlToMeshMtl(mtl *TopoMaterial) mst.MeshMaterial {
	return material.MtlToMeshMtl((*material.Material)(mtl))
}
