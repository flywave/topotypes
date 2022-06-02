package topotypes

import (
	"encoding/json"
	"errors"
	"strings"
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
	TOPO_TYPE_SWEEP_LAYERS
	TOPO_TYPE_SWEEP_LAYERS_INTERSECTION
	TOPO_TYPE_DECAL
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
	case TOPO_TYPE_SWEEP_LAYERS:
		return "sweep-layers"
	case TOPO_TYPE_SWEEP_LAYERS_INTERSECTION:
		return "sweep-layers-intersection"
	case TOPO_TYPE_DECAL:
		return "decal"
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
	} else if strEquals(tp, "sweep-layers") {
		return TOPO_TYPE_SWEEP_LAYERS
	} else if strEquals(tp, "sweep-layers-intersection") {
		return TOPO_TYPE_SWEEP_LAYERS_INTERSECTION
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
	case TOPO_TYPE_SWEEP_LAYERS_INTERSECTION:
		inter = &TopoSweepLayersIntersection{}
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
	default:
		return nil, errors.New("not support topo type")
	}
	e = json.Unmarshal(([]byte)(js), inter)
	if e != nil {
		return nil, e
	}
	return inter, nil
}
