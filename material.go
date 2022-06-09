package topotypes

import (
	"encoding/json"
	"fmt"

	mst "github.com/flywave/go-mst"
	"github.com/flywave/topotypes/material"
)

type TopoMaterial material.Material

func TopoMtlToMeshMtl(mtl *TopoMaterial) mst.MeshMaterial {
	return material.MtlToMeshMtl((*material.Material)(mtl))
}

type TopoMaterialMap map[string]*TopoMaterial

func (mk *TopoMaterialMap) UnmarshalJSON(data []byte) error {
	*mk = make(map[string]*TopoMaterial)
	var tmp interface{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	switch t := tmp.(type) {
	case map[string]interface{}:
		dt, _ := json.Marshal(t)
		mtls := make(map[string]*TopoMaterial)
		json.Unmarshal(dt, &mtls)
		*mk = mtls
	case []interface{}:
		if len(t) == 0 {
			return nil
		}
		mtls := make(map[string]*TopoMaterial)
		dt, _ := json.Marshal(t)
		ary := []*TopoMaterial{}
		json.Unmarshal(dt, &ary)

		for i, m := range ary {
			if m.Name == "" {
				m.Name = fmt.Sprintf("mtl_%d", i)
			}
			mtls[m.Name] = m
		}
		*mk = mtls
	}
	return nil
}
