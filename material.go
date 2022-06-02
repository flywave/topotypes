package topotypes

import mst "github.com/flywave/go-mst"

type TopoMaterial struct {
	Type             string   `json:"type"`
	Color            [3]byte  `json:"color"`
	Transparency     float64  `json:"transparency"`
	Ambient          [3]byte  `json:"ambient"`
	Emissive         [3]byte  `json:"emissive"`
	Specular         [3]byte  `json:"specular"`
	Shininess        *float64 `json:"shininess,omitempty"`
	Specularity      *float64 `json:"specularity,omitempty"`
	Roughness        *float64 `json:"roughness,omitempty"`
	Metallic         *float64 `json:"metallic,omitempty"`
	Reflectance      *float64 `json:"reflectance,omitempty"`
	AmbientOcclusion *float64 `json:"ambient-occlusion,omitempty"`
}

func (m *TopoMaterial) HasTexture() bool {
	return false
}

func TopoMtlToMeshMtl(mtl *TopoMaterial) mst.MeshMaterial {
	ty := StringToMaterialType(mtl.Type)
	switch ty {
	case TOPO_MATERIAL_TYPE_PBR:
		mt := &mst.PbrMaterial{}
		mt.Color = mtl.Color
		if mtl.Metallic != nil {
			mt.Metallic = float32(*mtl.Metallic)
		}
		mt.Emissive[0] = mtl.Emissive[0]
		mt.Emissive[1] = mtl.Emissive[1]
		mt.Emissive[2] = mtl.Emissive[2]
		mt.Emissive[3] = 1.0
		if mtl.Roughness != nil {
			mt.Roughness = float32(*mtl.Roughness)
		}
		if mtl.Reflectance != nil {
			mt.Reflectance = float32(*mtl.Reflectance)
		}
		if mtl.AmbientOcclusion != nil {
			mt.AmbientOcclusion = float32(*mtl.AmbientOcclusion)
		}
		return mt
	case TOPO_MATERIAL_TYPE_LAMBERT:
		mt := &mst.LambertMaterial{}
		mt.Ambient = mtl.Ambient
		mt.Emissive = mtl.Emissive
		mt.Color = mtl.Color
		return mt
	case TOPO_MATERIAL_TYPE_PHONG:
		mt := &mst.LambertMaterial{}
		mt.Ambient = mtl.Ambient
		mt.Emissive = mtl.Emissive
		mt.Color = mtl.Color

		mtp := &mst.PhongMaterial{LambertMaterial: *mt}
		mtp.Specular = mtl.Specular
		if mtl.Shininess != nil {
			mtp.Shininess = float32(*mtl.Shininess)
		}
		if mtl.Specularity != nil {
			mtp.Specularity = float32(*mtl.Specularity)
		}
		return mtp
	}
	return nil
}
