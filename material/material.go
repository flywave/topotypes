package material

import (
	"encoding/json"
	"fmt"

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
	Color            *[3]byte    `json:"color,omitempty"`
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
	TextureMap       *string     `json:"texture-map,omitempty"`
	TextureAngle     *float64    `json:"texture-angle,omitempty"`
}

func (m *Material) HasTexture() bool {
	return m.Texture != ""
}

func (m *Material) GetColor() [3]byte {
	if m.Color != nil {
		return *m.Color
	}
	return [3]byte{255, 255, 255}
}

func MtlToMeshMtl(mtl *Material) mst.MeshMaterial {
	if mtl == nil {
		return nil
	}
	ty := StringToMaterialType(mtl.Type)
	switch ty {
	case TYPE_BASE:
		mt := &mst.TextureMaterial{}
		mt.Color = mtl.GetColor()
		return mt
	case TYPE_PBR:
		mt := &mst.PbrMaterial{}
		mt.Color = mtl.GetColor()
		if mtl.Metallic != nil {
			mt.Metallic = float32(*mtl.Metallic)
		}
		mt.Emissive[0] = mtl.Emissive[0]
		mt.Emissive[1] = mtl.Emissive[1]
		mt.Emissive[2] = mtl.Emissive[2]
		if mtl.Roughness != nil {
			mt.Roughness = float32(*mtl.Roughness)
		} else {
			mt.Roughness = 1
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
		mt.Color = mtl.GetColor()
		return mt
	case TYPE_PHONG:
		mt := &mst.LambertMaterial{}
		mt.Ambient = mtl.Ambient
		mt.Emissive = mtl.Emissive
		mt.Color = mtl.GetColor()

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

func TopoMtlToMeshMtl(mtl *Material) mst.MeshMaterial {
	return MtlToMeshMtl((*Material)(mtl))
}

type TopoMaterialMap map[string]*Material

func (mk *TopoMaterialMap) UnmarshalJSON(data []byte) error {
	*mk = make(map[string]*Material)
	var tmp interface{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	switch t := tmp.(type) {
	case map[string]interface{}:
		dt, _ := json.Marshal(t)
		mtls := make(map[string]*Material)
		json.Unmarshal(dt, &mtls)
		*mk = mtls
	case []interface{}:
		if len(t) == 0 {
			return nil
		}
		mtls := make(map[string]*Material)
		dt, _ := json.Marshal(t)
		ary := []*Material{}
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
