package topotypes

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	mst "github.com/flywave/go-mst"
)

const (
	TOPO_TYPE_NONE = iota
	TOPO_TYPE_SHAPE
	TOPO_TYPE_PRISM
	TOPO_TYPE_REVOL
	TOPO_TYPE_PIPE
	TOPO_TYPE_COMPOUND
	TOPO_TYPE_CROSS_POINT
	TOPO_TYPE_SYMBOL
	TOPO_TYPE_SYMBOL_PATH
	TOPO_TYPE_SYMBOL_SURFACE
	TOPO_TYPE_TEXTURE_SURFACE
	TOPO_TYPE_MATERIAL_SURFACE
	TOPO_TYPE_MASK
	TOPO_TYPE_LIGHT
	TOPO_TYPE_LEVELED_SURFACE
	TOPO_TYPE_CAMERA
	TOPO_TYPE_CUSTOM
	TOPO_TYPE_CROSS_MULTI_POINT
	TOPO_TYPE_FEATURE
	TOPO_TYPE_PYRAMID
	TOPO_TYPE_SWEEP_LAYERS
)

func TopoTypeToString(tp int) string {
	switch tp {
	case TOPO_TYPE_SHAPE:
		return "shape"
	case TOPO_TYPE_REVOL:
		return "revol"
	case TOPO_TYPE_PRISM:
		return "prism"
	case TOPO_TYPE_CROSS_POINT:
		return "cross-point"
	case TOPO_TYPE_CROSS_MULTI_POINT:
		return "cross-multi-point"
	case TOPO_TYPE_PIPE:
		return "pipe"
	case TOPO_TYPE_SYMBOL_PATH:
		return "symbol-path"
	case TOPO_TYPE_MATERIAL_SURFACE:
		return "material-surface"
	case TOPO_TYPE_SYMBOL_SURFACE:
		return "symbol-surface"
	case TOPO_TYPE_TEXTURE_SURFACE:
		return "texture-surface"
	case TOPO_TYPE_LEVELED_SURFACE:
		return "leveled-surface"
	case TOPO_TYPE_SYMBOL:
		return "symbol"
	case TOPO_TYPE_MASK:
		return "mask"
	case TOPO_TYPE_LIGHT:
		return "light"
	case TOPO_TYPE_COMPOUND:
		return "compound"
	case TOPO_TYPE_CUSTOM:
		return "custom"
	case TOPO_TYPE_FEATURE:
		return "feature"
	case TOPO_TYPE_PYRAMID:
		return "pyramid"
	case TOPO_TYPE_SWEEP_LAYERS:
		return "sweep-layers"
	default:
		return ""
	}
}

func strEquals(t1, t2 string) bool {
	return strings.ToLower(t1) == strings.ToLower(t2)
}

func StringToTopoType(tp string) int {
	if strEquals(tp, "shape") {
		return TOPO_TYPE_SHAPE
	} else if strEquals(tp, "prism") {
		return TOPO_TYPE_PRISM
	} else if strEquals(tp, "revol") {
		return TOPO_TYPE_REVOL
	} else if strEquals(tp, "cross-point") {
		return TOPO_TYPE_CROSS_POINT
	} else if strEquals(tp, "pipe") {
		return TOPO_TYPE_PIPE
	} else if strEquals(tp, "symbol-path") {
		return TOPO_TYPE_SYMBOL_PATH
	} else if strEquals(tp, "material-surface") {
		return TOPO_TYPE_MATERIAL_SURFACE
	} else if strEquals(tp, "symbol-surface") {
		return TOPO_TYPE_SYMBOL_SURFACE
	} else if strEquals(tp, "texture-surface") {
		return TOPO_TYPE_TEXTURE_SURFACE
	} else if strEquals(tp, "leveled-surface") {
		return TOPO_TYPE_LEVELED_SURFACE
	} else if strEquals(tp, "symbol") {
		return TOPO_TYPE_SYMBOL
	} else if strEquals(tp, "mask") {
		return TOPO_TYPE_MASK
	} else if strEquals(tp, "light") {
		return TOPO_TYPE_LIGHT
	} else if strEquals(tp, "compound") {
		return TOPO_TYPE_COMPOUND
	} else if strEquals(tp, "custom") {
		return TOPO_TYPE_CUSTOM
	} else if strEquals(tp, "feature") {
		return TOPO_TYPE_FEATURE
	} else if strEquals(tp, "pyramid") {
		return TOPO_TYPE_PYRAMID
	} else if strEquals(tp, "sweep-layers") {
		return TOPO_TYPE_SWEEP_LAYERS
	}
	return TOPO_TYPE_NONE
}

const (
	TOPO_MESH_MODE_SOLID = iota
	TOPO_MESH_MODE_SHELL
)

func MeshModeToString(tp int) string {
	switch tp {
	case TOPO_MESH_MODE_SOLID:
		return "solid"
	case TOPO_MESH_MODE_SHELL:
		return "shell"
	default:
		return "solid"
	}
}

func StringToMeshMode(tp string) int {
	if strEquals(tp, "solid") {
		return TOPO_MESH_MODE_SOLID
	} else if strEquals(tp, "shell") {
		return TOPO_MESH_MODE_SHELL
	}
	return TOPO_MESH_MODE_SOLID
}

const (
	TOPO_QUAD_COORD_MODE_LOCAl = iota
	TOPO_QUAD_COORD_MODE_MERCATOR
)

func QuadCoordModeToString(tp int) string {
	switch tp {
	case TOPO_QUAD_COORD_MODE_LOCAl:
		return "local"
	case TOPO_QUAD_COORD_MODE_MERCATOR:
		return "mercator"
	default:
		return ""
	}
}

func StringToQuadCoordMode(tp string) int {
	if strEquals(tp, "local") {
		return TOPO_QUAD_COORD_MODE_LOCAl
	} else if strEquals(tp, "mercator") {
		return TOPO_QUAD_COORD_MODE_MERCATOR
	}
	return TOPO_QUAD_COORD_MODE_MERCATOR
}

const (
	TOPO_PATH_MODE_NONE = iota
	TOPO_PATH_MODE_REPEAT
	TOPO_PATH_MODE_SPACING
	TOPO_PATH_MODE_RANDOM
	TOPO_PATH_MODE_CENTER
)

func PathModeToString(tp int) string {
	switch tp {
	case TOPO_PATH_MODE_REPEAT:
		return "repeat"
	case TOPO_PATH_MODE_SPACING:
		return "spacing"
	case TOPO_PATH_MODE_RANDOM:
		return "random"
	case TOPO_PATH_MODE_CENTER:
		return "center"
	default:
		return ""
	}
}

func StringToPathMode(tp string) int {
	if strEquals(tp, "repeat") {
		return TOPO_PATH_MODE_REPEAT
	} else if strEquals(tp, "spacing") {
		return TOPO_PATH_MODE_SPACING
	} else if strEquals(tp, "random") {
		return TOPO_PATH_MODE_RANDOM
	} else if strEquals(tp, "center") {
		return TOPO_PATH_MODE_CENTER
	}
	return TOPO_PATH_MODE_NONE
}

const (
	TOPO_SURFACE_MODE_NONE = iota
	TOPO_SURFACE_MODE_GRID
	TOPO_SURFACE_MODE_RANDOM
	TOPO_SURFACE_MODE_CENTER
)

func SurfaceModeToString(tp int) string {
	switch tp {
	case TOPO_SURFACE_MODE_GRID:
		return "grid"
	case TOPO_SURFACE_MODE_RANDOM:
		return "random"
	case TOPO_SURFACE_MODE_CENTER:
		return "center"
	default:
		return ""
	}
}

func StringToSurfaceMode(tp string) int {
	if strEquals(tp, "grid") {
		return TOPO_SURFACE_MODE_GRID
	} else if strEquals(tp, "random") {
		return TOPO_SURFACE_MODE_RANDOM
	} else if strEquals(tp, "center") {
		return TOPO_SURFACE_MODE_CENTER
	}
	return TOPO_SURFACE_MODE_NONE
}

const (
	TOPO_SHAPE_MODE_NONE = iota
	TOPO_SHAPE_MODE_BOX
	TOPO_SHAPE_MODE_CYLINDER
	TOPO_SHAPE_MODE_CONE
	TOPO_SHAPE_MODE_SPHERE
	TOPO_SHAPE_MODE_TORUS
	TOPO_SHAPE_MODE_WEDGE
	TOPO_SHAPE_MODE_REVOLUTION
	TOPO_SHAPE_MODE_PIPE
)

func ShapeTypeToString(tp int) string {
	switch tp {
	case TOPO_SHAPE_MODE_BOX:
		return "box"
	case TOPO_SHAPE_MODE_CYLINDER:
		return "cylinder"
	case TOPO_SHAPE_MODE_CONE:
		return "cone"
	case TOPO_SHAPE_MODE_SPHERE:
		return "sphere"
	case TOPO_SHAPE_MODE_TORUS:
		return "torus"
	case TOPO_SHAPE_MODE_WEDGE:
		return "wedge"
	case TOPO_SHAPE_MODE_REVOLUTION:
		return "revolution"
	case TOPO_SHAPE_MODE_PIPE:
		return "pipe"
	default:
		return ""
	}
}

func StringToShapeType(tp string) int {
	if strEquals(tp, "box") {
		return TOPO_SHAPE_MODE_BOX
	} else if strEquals(tp, "cylinder") {
		return TOPO_SHAPE_MODE_CYLINDER
	} else if strEquals(tp, "cone") {
		return TOPO_SHAPE_MODE_CONE
	} else if strEquals(tp, "sphere") {
		return TOPO_SHAPE_MODE_SPHERE
	} else if strEquals(tp, "torus") {
		return TOPO_SHAPE_MODE_TORUS
	} else if strEquals(tp, "wedge") {
		return TOPO_SHAPE_MODE_WEDGE
	} else if strEquals(tp, "revolution") {
		return TOPO_SHAPE_MODE_REVOLUTION
	} else if strEquals(tp, "pipe") {
		return TOPO_SHAPE_MODE_PIPE
	}
	return TOPO_SHAPE_MODE_NONE
}

const (
	TOPO_SMOOTH_TYPE_BSPLINE = iota
)

func SmoothTypeToString(tp int) string {
	switch tp {
	case TOPO_SMOOTH_TYPE_BSPLINE:
		return "bspline"
	default:
		return "bspline"
	}
}

func StringToSmoothType(tp string) int {
	if strEquals(tp, "bspline") {
		return TOPO_SMOOTH_TYPE_BSPLINE
	}
	return TOPO_SMOOTH_TYPE_BSPLINE
}

const (
	TOPO_PROFILE_TYPE_NONE = iota
	TOPO_PROFILE_TYPE_TRIANGLE
	TOPO_PROFILE_TYPE_RECTANGLE
	TOPO_PROFILE_TYPE_CIRC
	TOPO_PROFILE_TYPE_ELIPS
	TOPO_PROFILE_TYPE_POLYGON
)

func ProfileTypeToString(tp int) string {
	switch tp {
	case TOPO_PROFILE_TYPE_TRIANGLE:
		return "triangle"
	case TOPO_PROFILE_TYPE_RECTANGLE:
		return "rectangle"
	case TOPO_PROFILE_TYPE_CIRC:
		return "circ"
	case TOPO_PROFILE_TYPE_ELIPS:
		return "ellipse"
	case TOPO_PROFILE_TYPE_POLYGON:
		return "polygon"
	default:
		return ""
	}
}

func StringToProfileType(tp string) int {
	if strEquals(tp, "triangle") {
		return TOPO_PROFILE_TYPE_TRIANGLE
	} else if strEquals(tp, "rectangle") {
		return TOPO_PROFILE_TYPE_RECTANGLE
	} else if strEquals(tp, "circ") {
		return TOPO_PROFILE_TYPE_CIRC
	} else if strEquals(tp, "ellipse") {
		return TOPO_PROFILE_TYPE_ELIPS
	} else if strEquals(tp, "polygon") {
		return TOPO_PROFILE_TYPE_POLYGON
	}
	return TOPO_PROFILE_TYPE_NONE
}

const (
	TOPO_COMPOUND_MODE_NONE = iota
	TOPO_COMPOUND_MODE_FUSION
	TOPO_COMPOUND_MODE_COMMON
	TOPO_COMPOUND_MODE_CUT
)

func CompoundModeToString(tp int) string {
	switch tp {
	case TOPO_COMPOUND_MODE_FUSION:
		return "fusion"
	case TOPO_COMPOUND_MODE_COMMON:
		return "common"
	case TOPO_COMPOUND_MODE_CUT:
		return "cut"
	default:
		return ""
	}
}

func StringToCompoundMode(tp string) int {
	if strEquals(tp, "fusion") {
		return TOPO_COMPOUND_MODE_FUSION
	} else if strEquals(tp, "common") {
		return TOPO_COMPOUND_MODE_COMMON
	} else if strEquals(tp, "cut") {
		return TOPO_COMPOUND_MODE_CUT
	}
	return TOPO_COMPOUND_MODE_FUSION
}

const (
	TOPO_LIGHT_MODE_NONE = iota
	TOPO_LIGHT_MODE_SPOT
	TOPO_LIGHT_MODE_POINT
	TOPO_LIGHT_MODE_DIRECTIONAL
	TOPO_LIGHT_MODE_AREA
)

func LightTypeToString(tp int) string {
	switch tp {
	case TOPO_LIGHT_MODE_SPOT:
		return "spot"
	case TOPO_LIGHT_MODE_POINT:
		return "point"
	case TOPO_LIGHT_MODE_DIRECTIONAL:
		return "directional"
	case TOPO_LIGHT_MODE_AREA:
		return "area"
	default:
		return ""
	}
}

func StringToLightType(tp string) int {
	if strEquals(tp, "spot") {
		return TOPO_LIGHT_MODE_SPOT
	} else if strEquals(tp, "point") {
		return TOPO_LIGHT_MODE_POINT
	} else if strEquals(tp, "directional") {
		return TOPO_LIGHT_MODE_DIRECTIONAL
	} else if strEquals(tp, "area") {
		return TOPO_LIGHT_MODE_AREA
	}
	return TOPO_LIGHT_MODE_NONE
}

const (
	TOPO_SHADOW_TYPE_NONE = iota
	TOPO_SHADOW_TYPE_HARD
	TOPO_SHADOW_TYPE_SOFT
)

func ShadowTypeToString(tp int) string {
	switch tp {
	case TOPO_SHADOW_TYPE_HARD:
		return "hard"
	case TOPO_SHADOW_TYPE_SOFT:
		return "soft"
	case TOPO_SHADOW_TYPE_NONE:
		return "none"
	default:
		return ""
	}
}

func StringToShadowType(tp string) int {
	if strEquals(tp, "hard") {
		return TOPO_SHADOW_TYPE_HARD
	} else if strEquals(tp, "soft") {
		return TOPO_SHADOW_TYPE_SOFT
	} else if strEquals(tp, "none") {
		return TOPO_SHADOW_TYPE_NONE
	}
	return TOPO_SHADOW_TYPE_NONE
}

const (
	TOPO_LEVELED_TYPE_ABSOLUTE = iota
	TOPO_LEVELED_TYPE_AVG
	TOPO_LEVELED_TYPE_MAX
	TOPO_LEVELED_TYPE_MIN
)

func LeveledTypeToString(tp int) string {
	switch tp {
	case TOPO_LEVELED_TYPE_ABSOLUTE:
		return "absolute"
	case TOPO_LEVELED_TYPE_AVG:
		return "avg"
	case TOPO_LEVELED_TYPE_MAX:
		return "max"
	case TOPO_LEVELED_TYPE_MIN:
		return "min"
	default:
		return ""
	}
}

func StringToLeveledType(tp string) int {
	if strEquals(tp, "absolute") {
		return TOPO_LEVELED_TYPE_ABSOLUTE
	} else if strEquals(tp, "avg") {
		return TOPO_LEVELED_TYPE_AVG
	} else if strEquals(tp, "max") {
		return TOPO_LEVELED_TYPE_MAX
	} else if strEquals(tp, "min") {
		return TOPO_LEVELED_TYPE_MIN
	}
	return TOPO_LEVELED_TYPE_ABSOLUTE
}

const (
	TOPO_MATERIAL_TYPE_NONE = iota
	TOPO_MATERIAL_TYPE_PBR
	TOPO_MATERIAL_TYPE_LAMBERT
	TOPO_MATERIAL_TYPE_PHONG
)

func MaterialTypeToString(tp int) string {
	switch tp {
	case TOPO_MATERIAL_TYPE_PBR:
		return "pbr"
	case TOPO_MATERIAL_TYPE_LAMBERT:
		return "lambert"
	case TOPO_MATERIAL_TYPE_PHONG:
		return "phong"
	default:
		return ""
	}
}

func StringToMaterialType(tp string) int {
	if strEquals(tp, "pbr") {
		return TOPO_MATERIAL_TYPE_PBR
	} else if strEquals(tp, "lambert") {
		return TOPO_MATERIAL_TYPE_LAMBERT
	} else if strEquals(tp, "phong") {
		return TOPO_MATERIAL_TYPE_PHONG
	}
	return TOPO_MATERIAL_TYPE_NONE
}

const (
	TOPO_WARP_TYPE_NONE = iota
	TOPO_WARP_TYPE_REPEAT
	TOPO_WARP_TYPE_CLAMP
)

func WarpTypeToString(tp int) string {
	switch tp {
	case TOPO_WARP_TYPE_REPEAT:
		return "repeat"
	case TOPO_WARP_TYPE_CLAMP:
		return "clamp"
	default:
		return ""
	}
}

func StringToWarpType(tp string) int {
	if strEquals(tp, "repeat") {
		return TOPO_WARP_TYPE_REPEAT
	} else if strEquals(tp, "clamp") {
		return TOPO_WARP_TYPE_CLAMP
	}
	return TOPO_WARP_TYPE_NONE
}

const (
	TOPO_FEATURE_BOUND_TYPE_NONE = iota
	TOPO_FEATURE_BOUND_TYPE_BBOX
	TOPO_FEATURE_BOUND_TYPE_BSPHERE
	TOPO_FEATURE_BOUND_TYPE_BPATH
	TOPO_FEATURE_BOUND_TYPE_BPOLYGON
)

func FeatureBoundTypeToString(tp int) string {
	switch tp {
	case TOPO_FEATURE_BOUND_TYPE_BBOX:
		return "bbox"
	case TOPO_FEATURE_BOUND_TYPE_BSPHERE:
		return "bsphere"
	case TOPO_FEATURE_BOUND_TYPE_BPATH:
		return "bpath"
	case TOPO_FEATURE_BOUND_TYPE_BPOLYGON:
		return "bpolygon"
	default:
		return ""
	}
}

func StringToFeatureBoundType(tp string) int {
	if strEquals(tp, "bbox") {
		return TOPO_FEATURE_BOUND_TYPE_BBOX
	} else if strEquals(tp, "bsphere") {
		return TOPO_FEATURE_BOUND_TYPE_BSPHERE
	} else if strEquals(tp, "bpath") {
		return TOPO_FEATURE_BOUND_TYPE_BPATH
	} else if strEquals(tp, "bpolygon") {
		return TOPO_FEATURE_BOUND_TYPE_BPOLYGON
	}
	return TOPO_FEATURE_BOUND_TYPE_NONE
}

const (
	TOPO_CENTER_MODE_SPHERE = iota
	TOPO_CENTER_MODE_BOX
)

func CenterModeToString(tp int) string {
	switch tp {
	case TOPO_CENTER_MODE_BOX:
		return "box"
	case TOPO_CENTER_MODE_SPHERE:
		return "sphere"
	default:
		return ""
	}
}

func StringToCenterMode(tp string) int {
	if strEquals(tp, "box") {
		return TOPO_CENTER_MODE_BOX
	} else if strEquals(tp, "sphere") {
		return TOPO_CENTER_MODE_SPHERE
	}
	return TOPO_CENTER_MODE_SPHERE
}

const (
	TOPO_PYRAMID_MODE_TILED = iota
	TOPO_PYRAMID_MODE_FLAT
)

func PyramidModeToString(tp int) string {
	switch tp {
	case TOPO_PYRAMID_MODE_TILED:
		return "tiled"
	case TOPO_PYRAMID_MODE_FLAT:
		return "flat"
	default:
		return ""
	}
}

func StringToPyramidMode(tp string) int {
	if strEquals(tp, "tiled") {
		return TOPO_PYRAMID_MODE_TILED
	} else if strEquals(tp, "flat") {
		return TOPO_PYRAMID_MODE_FLAT
	}
	return TOPO_PYRAMID_MODE_FLAT
}

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

type TopoTransform struct {
	Rotation  *[4]float64 `json:"rotation"`
	Translate *[3]float64 `json:"translate"`
	Scale     *[3]float64 `json:"scale"`
}

func NewTopoTransform() *TopoTransform {
	return &TopoTransform{Scale: &[3]float64{1.0, 1.0, 1.0}}
}

type ToposInterface interface {
	GetType() string
	GetTransform() *TopoTransform
	GetFusion() bool
	ResetTransform()
}

const (
	TOPO_SIMPLIFY_LOW    = "low"
	TOPO_SIMPLIFY_NORMAL = "normal"
	TOPO_SIMPLIFY_HIGH   = "high"
)

type TopoZoom struct {
	Simplify  string    `json:"simplify,omitempty"`
	ZoomRange [2]uint32 `json:"zoom_range,omitempty"`
}

type Topos struct {
	Type      string         `json:"type"`
	Transform *TopoTransform `json:"transform,omitempty"`
	BBox      *[2][3]float64 `json:"bbox,omitempty"`
	Fusion    bool           `json:"fusion"`
	Zooms     []*TopoZoom    `json:"zooms"`
}

func (tp *Topos) GetZooms() []*TopoZoom {
	return tp.Zooms
}

func (tp *Topos) GetType() string {
	return tp.Type
}

func (tp *Topos) GetTransform() *TopoTransform {
	return tp.Transform
}

func (tp *Topos) ResetTransform() {
	tp.Transform = NewTopoTransform()
}

func (tp *Topos) GetFusion() bool {
	return tp.Fusion
}

type TopoMakerInterface interface {
	GetMaterials() []*TopoMaterial
	IsInstance() bool
}

type TopoMaker struct {
	Topos
	MeshMode  string          `json:"mode"`
	Materials []*TopoMaterial `json:"materials,omitempty"`
	Instanced bool            `json:"instanced,omitempty"`
}

func (t *TopoMaker) IsInstance() bool {
	return t.Instanced
}

func (t *TopoMaker) GetMaterials() []*TopoMaterial {
	return t.Materials
}

type TopoMaterialSurface struct {
	TopoPrism
}

type TopoTextureSurface struct {
	TopoPrism
	WarpS   string `json:"warp-s"`
	WarpT   string `json:"warp-t"`
	Zoom    uint8  `json:"zoom"`
	Texture string `json:"texture"`
}

type TopoBoundy interface {
	IsTopoBoundy() bool
}

type TopoShape struct {
	TopoMaker
	Shape      string      `json:"-"`
	ShapeModel interface{} `json:"shape"`
}

func (sp *TopoShape) IsTopoBoundy() bool {
	return true
}

type TopoShapeBox struct {
	Type   string
	Point1 [3]float64 `json:"point1"`
	Point2 [3]float64 `json:"point2"`
}

func NewTopoShapeBox() *TopoShapeBox {
	return &TopoShapeBox{
		Type: TopoTypeToString(TOPO_SHAPE_MODE_BOX),
	}
}

type TopoShapeCone struct {
	Type    string   `json:"type"`
	Radius1 float64  `json:"radius1"`
	Radius2 float64  `json:"radius2"`
	Height  float64  `json:"height"`
	Angle   *float64 `json:"angle,omitempty"`
}

func NewTopoShapeCone() *TopoShapeCone {
	return &TopoShapeCone{
		Type: TopoTypeToString(TOPO_SHAPE_MODE_CONE),
	}
}

type TopoShapeCylinder struct {
	Type   string   `json:"type"`
	Radius float64  `json:"radius"`
	Height float64  `json:"height"`
	Angle  *float64 `json:"angle,omitempty"`
}

func NewTopoShapeCylinder() *TopoShapeCylinder {
	return &TopoShapeCylinder{
		Type: TopoTypeToString(TOPO_SHAPE_MODE_CYLINDER),
	}
}

type TopoShapeRevolution struct {
	Type     string
	Meridian [][3]float64 `json:"meridian"`
	Angle    *float64     `json:"angle,omitempty"`
	Max      *float64     `json:"max,omitempty"`
	Min      *float64     `json:"min,omitempty"`
}

func NewTopoShapeRevolution() *TopoShapeRevolution {
	return &TopoShapeRevolution{
		Type: TopoTypeToString(TOPO_SHAPE_MODE_REVOLUTION),
	}
}

type TopoShapeSphere struct {
	Type   string
	Center *[3]float64 `json:"center,omitempty"`
	Radius float64     `json:"radius"`
	Angle1 *float64    `json:"angle1,omitempty"`
	Angle2 *float64    `json:"angle2,omitempty"`
	Angle  *float64    `json:"angle,omitempty"`
}

func NewTopoShapeSphere() *TopoShapeSphere {
	return &TopoShapeSphere{
		Type: TopoTypeToString(TOPO_SHAPE_MODE_SPHERE),
	}
}

type TopoShapeTorus struct {
	Type    string
	Radius1 float64  `json:"radius1"`
	Radius2 float64  `json:"radius2"`
	Angle1  *float64 `json:"angle1,omitempty"`
	Angle2  *float64 `json:"angle2,omitempty"`
	Angle   *float64 `json:"angle,omitempty"`
}

func NewTopoShapeTorus() *TopoShapeTorus {
	return &TopoShapeTorus{
		Type: TopoTypeToString(TOPO_SHAPE_MODE_TORUS),
	}
}

// profiles
type WedgeFaceLimit struct {
	XMin float64 `json:"xmin"`
	ZMin float64 `json:"zmin"`
	XMax float64 `json:"xmax"`
	ZMax float64 `json:"zmax"`
}

type TopoProfile interface {
}

type TopoTriangle struct {
	TopoProfile
	Type string     `json:"type"`
	P1   [3]float64 `json:"p1"`
	P2   [3]float64 `json:"p2"`
	P3   [3]float64 `json:"p3"`
}

func NewTopoTriangle() *TopoTriangle {
	t := TopoTriangle{}
	t.Type = ProfileTypeToString(TOPO_PROFILE_TYPE_TRIANGLE)
	return &t
}

type TopoRectangle struct {
	TopoProfile
	Type string     `json:"type"`
	P1   [3]float64 `json:"p1"`
	P2   [3]float64 `json:"p2"`
}

func NewTopoRectangle() *TopoRectangle {
	t := TopoRectangle{}
	t.Type = ProfileTypeToString(TOPO_PROFILE_TYPE_RECTANGLE)
	return &t
}

type TopoCirc struct {
	TopoProfile
	Type   string     `json:"type"`
	Center [3]float64 `json:"center"`
	Norm   [3]float64 `json:"norm"`
	Radius float64    `json:"radius"`
}

func NewTopoCirc() *TopoCirc {
	t := TopoCirc{}
	t.Type = ProfileTypeToString(TOPO_PROFILE_TYPE_CIRC)
	return &t
}

type TopoElips struct {
	TopoProfile
	Type   string     `json:"type"`
	S1     [3]float64 `json:"s1"`
	S2     [3]float64 `json:"s2"`
	Center [3]float64 `json:"center"`
}

func NewTopoElips() *TopoElips {
	t := TopoElips{}
	t.Type = ProfileTypeToString(TOPO_PROFILE_TYPE_ELIPS)
	return &t
}

type TopoPolygon struct {
	TopoProfile
	Type  string       `json:"type"`
	Edges [][3]float64 `json:"edges,omitempty"`
}

func NewTopoPolygon() *TopoPolygon {
	t := TopoPolygon{}
	t.Type = ProfileTypeToString(TOPO_PROFILE_TYPE_POLYGON)
	return &t
}

//profile end

type TopoAnchorRef struct {
	Ref string `json:"ref"`
}

type TopoAnchorLink struct {
	Link           string `json:"link"`
	AnchorName     string `json:"anchor_name"`
	DestAnchorName string `json:"dest_anchor_name"`
}

type TopoAnchor struct {
	Name     string      `json:"name"`
	Position *[3]float64 `json:"position,omitempty"`
	Link     string      `json:"link,omitempty"`
}

const (
	TransitionModeRight = "right"
	TransitionModeRound = "round"
	TransitionModeTrans = "trans"
)

func TransitionModeToInt(m string) int {
	switch m {
	case TransitionModeRight:
		return 1
	case TransitionModeRound:
		return 2
	case TransitionModeTrans:
		return 0
	default:
		return 1
	}
}

func TransitionModeToString(m int) string {
	switch m {
	case 1:
		return TransitionModeRight
	case 2:
		return TransitionModeRound
	}
	return TransitionModeTrans
}

type TopoPipe struct {
	TopoMaker
	Wire           [][3]float64   `json:"-"`
	Profile        interface{}    `json:"profile"`
	Anchors        [2]*TopoAnchor `json:"anchors"`
	Customs        [2]*string     `json:"customs"`
	Smooth         string         `json:"smooth,omitempty"`
	TransitionMode string         `json:"transition_mode"`
}

func NewTopoPipe() *TopoPipe {
	t := &TopoPipe{}
	t.Type = TopoTypeToString(TOPO_TYPE_PIPE)
	return t
}

func (sp *TopoPipe) IsTopoBoundy() bool {
	return true
}

type TopoShapePipe struct {
	TopoPipe
	Shape string
}

func NewTopoShapePipe() *TopoShapePipe {
	t := &TopoShapePipe{
		Shape: ShapeTypeToString(TOPO_SHAPE_MODE_PIPE),
	}
	t.Type = TopoTypeToString(TOPO_TYPE_SHAPE)
	return t
}

type TopoShapeWedge struct {
	Type  string
	Edge  [3]float64      `json:"edge"`
	Limit *WedgeFaceLimit `json:"limit,omitempty"`
	Ltx   *float64        `json:"ltx,omitempty"`
}

func NewTopoShapeWedge() *TopoShapeWedge {
	return &TopoShapeWedge{Type: TopoTypeToString(TOPO_TYPE_SHAPE)}
}

type TopoDraft struct {
	Angle float64 `json:"angle"`
}

type LightInterface interface {
	GetLight() string
}

type TopoLight struct {
	Topos
	Light      string  `json:"light"`
	Color      []uint8 `json:"color"`
	Instensity float64 `json:"instensity"`
}

func (lt *TopoLight) GetLight() string {
	return lt.Light
}

type TopoAreaLight struct {
	TopoLight
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func NewTopoAreaLight() *TopoAreaLight {
	l := TopoAreaLight{}
	l.Light = LightTypeToString(TOPO_LIGHT_MODE_AREA)
	l.Type = TopoTypeToString(TOPO_TYPE_LIGHT)
	return &l
}

type TopoDirectionalLight struct {
	TopoLight
	Dir          [3]float64 `json:"dir"`
	Shadow       int        `json:"shadow"`
	Strength     *float64   `json:"strength,omitempty"`
	Bias         *float64   `json:"bias,omitempty"`
	Softness     *float64   `json:"softness,omitempty"`
	SoftnessFade *float64   `json:"softness-fade,omitempty"`
}

func NewTopoDirectionalLight() *TopoDirectionalLight {
	l := TopoDirectionalLight{}
	l.Light = LightTypeToString(TOPO_LIGHT_MODE_DIRECTIONAL)
	l.Type = TopoTypeToString(TOPO_TYPE_LIGHT)
	return &l
}

type TopoPointLight struct {
	TopoLight
	Distance float64 `json:"distance"`
}

func NewTopoPointLight() *TopoPointLight {
	l := TopoPointLight{}
	l.Light = LightTypeToString(TOPO_LIGHT_MODE_POINT)
	l.Type = TopoTypeToString(TOPO_TYPE_LIGHT)
	return &l
}

type TopoSpotLight struct {
	TopoLight
	Dir      [3]float64 `json:"dir"`
	Angle    float64    `json:"angle"`
	Exponent float64    `json:"exponent"`
	Distance float64    `json:"distance"`
}

func NewTopoSpotLight() *TopoSpotLight {
	l := TopoSpotLight{}
	l.Light = LightTypeToString(TOPO_LIGHT_MODE_SPOT)
	l.Type = TopoTypeToString(TOPO_TYPE_LIGHT)
	return &l
}

type TopoCrossMultiPoint struct {
	Topos
	Refs    []TopoAnchorRef   `json:"links,omitempty"`
	Objects []*TopoCrossPoint `json:"objects,omitempty"`
}

func NewTopoCrossMultiPoint() *TopoCrossMultiPoint {
	t := &TopoCrossMultiPoint{}
	t.Type = TopoTypeToString(TOPO_TYPE_CROSS_MULTI_POINT)
	return t
}

type TopoCrossPoint struct {
	Topos
	Model string           `json:"model"`
	Links []TopoAnchorLink `json:"links,omitempty"`
}

func NewTopoCrossPoint() *TopoCrossPoint {
	t := &TopoCrossPoint{}
	t.Type = TopoTypeToString(TOPO_TYPE_CROSS_POINT)
	return t
}

type CompoundObject struct {
	Name      string         `json:"name"`
	Shape     ToposInterface `json:"shape"`
	Transform *TopoTransform `json:"transform"`
}

type CompoundGroup struct {
	Objects []string `json:"objects"`
	Type    string   `json:"type"`
}

func NewCompoundGroup(ty int) *CompoundGroup {
	return &CompoundGroup{Type: CompoundModeToString(ty)}
}

type TopoCompound struct {
	TopoMaker
	Objects []CompoundObject `json:"objects,omitempty"`
	Groups  []CompoundGroup  `json:"groups,omitempty"`
}

func NewCompound() *TopoCompound {
	t := &TopoCompound{}
	t.Type = TopoTypeToString(TOPO_TYPE_COMPOUND)
	return t
}

type TopoCustom struct {
	TopoMaker
	CenterMode string   `json:"mode"`
	In         []string `json:"in-pipe-ids,omitempty"`
	Out        []string `json:"out-pipe-ids,omitempty"`
}

func NewTopoCustom() *TopoCustom {
	t := &TopoCustom{}
	t.Type = TopoTypeToString(TOPO_TYPE_CUSTOM)
	return t
}

type TopoMask struct {
	Topos
	Model string `json:"model"`
}

type PrismInterface interface {
	GetDirection() *[3]float64
}

type TopoPrism struct {
	TopoMaker
	Profile   interface{} `json:"profile,omitempty"`
	Direction [3]float64  `json:"direction"`
}

func NewTopoPrism() *TopoPrism {
	t := &TopoPrism{}
	t.Type = TopoTypeToString(TOPO_TYPE_PRISM)
	return t
}

func (sp *TopoPrism) GetDirection() *[3]float64 {
	return &sp.Direction
}

func (sp *TopoPrism) IsTopoBoundy() bool {
	return true
}

type TopoRevol struct {
	TopoMaker
	Profile interface{}   `json:"profile"`
	Axis    [2][3]float64 `json:"axis"`
	Angle   float64       `json:"angle"`
}

func NewTopoRevol() *TopoRevol {
	t := &TopoRevol{}
	t.Type = TopoTypeToString(TOPO_TYPE_REVOL)
	return t
}

func (sp *TopoRevol) IsTopoBoundy() bool {
	return true
}

type TopoSurface struct {
	Topos
}

type TopoLeveledSurface struct {
	TopoSurface
	Leveled string `json:"leveled"`
	RowSize uint32 `json:"row-size"`
}

func NewTopoLeveledSurface(lvlType int) *TopoLeveledSurface {
	t := &TopoLeveledSurface{Leveled: LeveledTypeToString(lvlType)}
	t.Type = TopoTypeToString(TOPO_TYPE_LEVELED_SURFACE)
	return t
}

type TopoSymbol struct {
	Topos
	Model     string `json:"model"`
	Instanced bool   `json:"instanced"`
}

func NewTopoSymbol() *TopoSymbol {
	t := &TopoSymbol{}
	t.Type = TopoTypeToString(TOPO_TYPE_SYMBOL)
	return t
}

type TopoSymbolPath struct {
	Topos
	Model   string  `json:"model"`
	Mode    string  `json:"mode"`
	Density float64 `json:"density"`
}

func NewTopoSymbolPath(md int) *TopoSymbolPath {
	t := &TopoSymbolPath{Mode: PathModeToString(md)}
	t.Type = TopoTypeToString(TOPO_TYPE_SYMBOL_PATH)
	return t
}

type TopoSymbolSurface struct {
	TopoSurface
	Model string     `json:"model"`
	Mode  string     `json:"mode"`
	Cell  [2]float64 `json:"cell"`
}

func NewTopoSymbolSurface(md int) *TopoSymbolSurface {
	t := &TopoSymbolSurface{Mode: SurfaceModeToString(md)}
	t.Type = TopoTypeToString(TOPO_TYPE_SYMBOL_SURFACE)
	return t
}

type TopoCamera struct {
	Topos
}

type TopoPyramid struct {
	Topos
	Zoom  uint32 `json:"zoom"`
	Model string `json:"model"`
	Mode  string `json:"mode"`
}

func NewTopoPyramid(md int) *TopoPyramid {
	t := &TopoPyramid{Mode: PyramidModeToString(md)}
	t.Type = TopoTypeToString(TOPO_TYPE_PYRAMID)
	return t
}

type BoundPath struct {
	Radius float32 `json:"radius"`
}

type BoundPolygon struct {
	Height float32 `json:"height"`
}

type BoundSphere struct {
	Center [3]float32 `json:"center"`
	Radius float32    `json:"radius"`
}

type FeatureBound struct {
	BBox     *[2][3]float32 `json:"bbox3d,omitempty"`
	BSphere  *BoundSphere   `json:"sphere3d,omitempty"`
	BPath    *BoundPath     `json:"bpath,omitempty"`
	BPolygon *BoundPolygon  `json:"bpolygon,omitempty"`
}

type TopoFeature struct {
	Topos
	Bounds []FeatureBound `json:"bounds,omitempty"`
}

type TopoLayer struct {
	Name    string      `json:"name,omitempty"`
	Profile TopoProfile `json:"profile"`
	Offset  [3]float64  `json:"offset,omitempty"`
	Texture string      `json:"texture"`
	// Boolean string      `json:"boolean,omitempty"`
}

type TopoSweepLayers struct {
	TopoMaker
	Layers []*TopoLayer `json:"layers,omitempty"`
}

func ProfileUnMarshal(inter interface{}) (interface{}, error) {
	switch pro := inter.(type) {
	case map[string]interface{}:
		v, ok := pro["type"]
		t, ok2 := v.(string)
		if !ok || !ok2 {
			return nil, errors.New("profile type error")
		}
		pro_t := StringToProfileType(t)
		js, er := json.Marshal(inter)
		if er != nil {
			return nil, er
		}
		var pf interface{}
		switch pro_t {
		case TOPO_PROFILE_TYPE_TRIANGLE:
			pf = NewTopoTriangle()
		case TOPO_PROFILE_TYPE_RECTANGLE:
			pf = NewTopoRectangle()
		case TOPO_PROFILE_TYPE_CIRC:
			pf = NewTopoCirc()
		case TOPO_PROFILE_TYPE_ELIPS:
			pf = NewTopoElips()
		case TOPO_PROFILE_TYPE_POLYGON:
			pf = NewTopoPolygon()
		default:
			return nil, errors.New("profile type error")
		}
		e := json.Unmarshal(([]byte)(js), pf)
		if e != nil {
			return nil, e
		}
		return pf, nil
	default:
		return nil, errors.New("profile type error")
	}
}

func ShapeUnMarshal(js []byte) (ToposInterface, error) {
	sp := make(map[string]interface{})
	e := json.Unmarshal(js, &sp)
	if e != nil {
		return nil, e
	}
	t, ok := sp["shape"]
	if !ok {
		return nil, errors.New("invalid json")
	}
	bt, _ := json.Marshal(t)
	e = json.Unmarshal(bt, &sp)
	if e != nil {
		return nil, e
	}
	t, ok = sp["type"]
	if !ok {
		return nil, errors.New("invalid json")
	}
	var inter interface{}
	res := &TopoShape{}
	res.Shape = t.(string)
	ty := StringToShapeType(res.Shape)
	switch ty {
	case TOPO_SHAPE_MODE_BOX:
		inter = NewTopoShapeBox()
	case TOPO_SHAPE_MODE_CYLINDER:
		inter = NewTopoShapeCylinder()
	case TOPO_SHAPE_MODE_CONE:
		inter = NewTopoShapeCone()
	case TOPO_SHAPE_MODE_SPHERE:
		inter = NewTopoShapeSphere()
	case TOPO_SHAPE_MODE_TORUS:
		inter = NewTopoShapeTorus()
	case TOPO_SHAPE_MODE_WEDGE:
		inter = NewTopoShapeWedge()
	case TOPO_SHAPE_MODE_REVOLUTION:
		inter = NewTopoShapeRevolution()
	case TOPO_SHAPE_MODE_PIPE:
		p := NewTopoShapePipe()
		tp, e := PipeUnMarshal(js)
		if e != nil {
			return nil, e
		}
		p.TopoPipe = *tp
		return p, nil
	default:
		return nil, errors.New("not support topo type")
	}

	res.ShapeModel = inter
	e = json.Unmarshal(([]byte)(js), res)
	if e != nil {
		return nil, e
	}
	return res, nil
}

func LightUnMarshal(js []byte) (ToposInterface, error) {
	lt := TopoLight{}
	e := json.Unmarshal(js, &lt)
	if e != nil {
		return nil, e
	}
	var inter ToposInterface
	ty := StringToLightType(lt.Light)
	switch ty {
	case TOPO_LIGHT_MODE_SPOT:
		inter = NewTopoSpotLight()
	case TOPO_LIGHT_MODE_POINT:
		inter = NewTopoPointLight()
	case TOPO_LIGHT_MODE_DIRECTIONAL:
		inter = NewTopoDirectionalLight()
	case TOPO_LIGHT_MODE_AREA:
		inter = NewTopoAreaLight()
	default:
		return nil, errors.New("not support topo type")
	}
	e = json.Unmarshal(([]byte)(js), inter)
	if e != nil {
		return nil, e
	}
	return inter, nil
}

func PrismUnMarshal(js []byte) (*TopoPrism, error) {
	pris := TopoPrism{}
	e := json.Unmarshal(js, &pris)
	if e != nil {
		return nil, e
	}
	if pris.Profile != nil {
		prof, er := ProfileUnMarshal(pris.Profile)
		if er != nil {
			return nil, er
		}
		pris.Profile = prof
	}
	return &pris, nil
}

func RevolUnMarshal(js []byte) (*TopoRevol, error) {
	revol := TopoRevol{}
	e := json.Unmarshal(js, &revol)
	if e != nil {
		return nil, e
	}
	prof, er := ProfileUnMarshal(revol.Profile)
	if er != nil {
		return nil, er
	}
	revol.Profile = prof
	return &revol, nil
}

func PipeUnMarshal(js []byte) (*TopoPipe, error) {
	pipe := TopoPipe{}
	e := json.Unmarshal(js, &pipe)
	if e != nil {
		return nil, e
	}
	if pipe.Profile != nil {
		prof, er := ProfileUnMarshal(pipe.Profile)
		if er != nil {
			return nil, er
		}
		pipe.Profile = prof
	}
	return &pipe, nil
}

func SweepLayersUnMarshal(js []byte) (*TopoSweepLayers, error) {
	sl := TopoSweepLayers{}
	e := json.Unmarshal(js, &sl)
	if e != nil {
		return nil, e
	}
	for _, l := range sl.Layers {
		prof, er := ProfileUnMarshal(l.Profile)
		if er != nil {
			return nil, er
		}
		l.Profile = prof
	}
	return &sl, nil
}

func CompoundUnMarshal(js []byte) (*TopoCompound, error) {
	tc := &TopoCompound{}
	e := json.Unmarshal(js, tc)
	if e != nil {
		return nil, e
	}
	if len(tc.Objects) == 0 {
		return tc, nil
	}

	for index, obj := range tc.Objects {
		bt, e := json.Marshal(obj.Shape)
		if e != nil {
			return nil, e
		}
		inter, e := TopoUnMarshal(bt)
		if e != nil {
			return nil, e
		}
		_, ok := inter.(TopoBoundy)
		if !ok {
			return nil, fmt.Errorf("index %d compound object must be Interface CompoundTopo", index)
		}
		tc.Objects[index].Shape = inter
	}
	return tc, nil
}

func TopoUnMarshal(js []byte) (ToposInterface, error) {
	base := Topos{}
	e := json.Unmarshal(js, &base)
	if e != nil {
		return nil, e
	}
	var inter ToposInterface
	ty := StringToTopoType(base.Type)
	switch ty {
	case TOPO_TYPE_SHAPE:
		return ShapeUnMarshal(js)
	case TOPO_TYPE_PRISM:
		return PrismUnMarshal(js)
	case TOPO_TYPE_REVOL:
		return RevolUnMarshal(js)
	case TOPO_TYPE_PIPE:
		return PipeUnMarshal(js)
	case TOPO_TYPE_SWEEP_LAYERS:
		return SweepLayersUnMarshal(js)
	case TOPO_TYPE_COMPOUND:
		return CompoundUnMarshal(js)
	case TOPO_TYPE_CROSS_POINT:
		inter = &TopoCrossPoint{}
	case TOPO_TYPE_SYMBOL:
		inter = &TopoSymbol{}
	case TOPO_TYPE_SYMBOL_PATH:
		inter = &TopoSymbolPath{}
	case TOPO_TYPE_SYMBOL_SURFACE:
		inter = &TopoSymbolSurface{}
	case TOPO_TYPE_TEXTURE_SURFACE:
		inter = &TopoTextureSurface{}
	case TOPO_TYPE_MATERIAL_SURFACE:
		inter = &TopoMaterialSurface{}
	case TOPO_TYPE_MASK:
		inter = &TopoMask{}
	case TOPO_TYPE_LIGHT:
		return LightUnMarshal(js)
	case TOPO_TYPE_LEVELED_SURFACE:
		inter = &TopoLeveledSurface{}
	case TOPO_TYPE_CAMERA:
		inter = &TopoCamera{}
	case TOPO_TYPE_CUSTOM:
		inter = &TopoCustom{}
	case TOPO_TYPE_CROSS_MULTI_POINT:
		inter = &TopoCrossMultiPoint{}
	case TOPO_TYPE_FEATURE:
		inter = &TopoFeature{}
	case TOPO_TYPE_PYRAMID:
		inter = &TopoPyramid{}
	default:
		return nil, errors.New("not support topo type")
	}
	e = json.Unmarshal(([]byte)(js), inter)
	if e != nil {
		return nil, e
	}
	return inter, nil
}
