package material

import (
	"github.com/flywave/go-mst"
	"github.com/flywave/topotypes/utils"
)

const (
	TYPE_NONE = iota
	TYPE_PBR
	TYPE_LAMBERT
	TYPE_PHONG
	TYPE_BASE
)

func MaterialTypeToString(tp int) string {
	switch tp {
	case TYPE_PBR:
		return "pbr"
	case TYPE_LAMBERT:
		return "lambert"
	case TYPE_PHONG:
		return "phong"
	case TYPE_BASE:
		return "base"
	default:
		return ""
	}
}

func StringToMaterialType(tp string) int {
	if utils.StrEquals(tp, "pbr") {
		return TYPE_PBR
	} else if utils.StrEquals(tp, "lambert") {
		return TYPE_LAMBERT
	} else if utils.StrEquals(tp, "phong") {
		return TYPE_PHONG
	} else if utils.StrEquals(tp, "base") {
		return TYPE_BASE
	}
	return TYPE_NONE
}

const (
	TextureCube = iota
	TextureNormal
	TextureNormalAutoScale
)

func TxtureMapToString(tp int) string {
	switch tp {
	case TextureCube:
		return "cube"
	case TextureNormal:
		return "normal"
	case TextureNormalAutoScale:
		return "normal-auto-scale"
	default:
		return ""
	}
}

func StringToTxtureMap(tp string) int {
	if utils.StrEquals(tp, "cube") {
		return TextureCube
	} else if utils.StrEquals(tp, "normal") {
		return TextureNormal
	} else if utils.StrEquals(tp, "normal-auto-scale") {
		return TextureNormalAutoScale
	}
	return TYPE_NONE
}

type Material struct {
	Name             string      `json:"name,omitempty"`
	Type             string      `json:"type"`
	Color            [3]byte     `json:"color"`
	Transparency     float64     `json:"transparency"`
	Ambient          [3]byte     `json:"ambient"`
	Emissive         [3]byte     `json:"emissive"`
	Specular         [3]byte     `json:"specular"`
	Shininess        *float64    `json:"shininess,omitempty"`
	Specularity      *float64    `json:"specularity,omitempty"`
	Roughness        *float64    `json:"roughness,omitempty"`
	Metallic         *float64    `json:"metallic,omitempty"`
	Reflectance      *float64    `json:"reflectance,omitempty"`
	AmbientOcclusion *float64    `json:"ambient-occlusion,omitempty"`
	Texture          string      `json:"texture,omitempty"`
	TextureScale     *[2]float64 `json:"texture-scale,omitempty"`
	TextureOrigin    *[2]float64 `json:"texture-origin,omitempty"`
	TextureRepeat    *[2]float64 `json:"texture-repeat,omitempty"`
	TextureAutoScale *[2]float64 `json:"texture-auto-scale,omitempty"`
	TxtureMap        *string     `json:"texture-map,omitempty"`
	TextureAngle     *float64    `json:"texture-angle,omitempty"`
}

func (m *Material) HasTexture() bool {
	return m.Texture != ""
}

func MtlToMeshMtl(mtl *Material) mst.MeshMaterial {
	ty := StringToMaterialType(mtl.Type)
	switch ty {
	case TYPE_BASE:
		mt := &mst.TextureMaterial{}
		mt.Color = mtl.Color
		return mt
	case TYPE_PBR:
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
	case TYPE_LAMBERT:
		mt := &mst.LambertMaterial{}
		mt.Ambient = mtl.Ambient
		mt.Emissive = mtl.Emissive
		mt.Color = mtl.Color
		return mt
	case TYPE_PHONG:
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
