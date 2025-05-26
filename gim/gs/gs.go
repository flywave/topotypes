package gs

import (
	"encoding/json"
	"fmt"
)

const Major = "GS"

type GsBase struct {
	Version int    `json:"version"`
	Type    string `json:"type"`
}

func (b *GsBase) GetType() string {
	return b.Type
}

// SphereObject represents a sphere
type Sphere struct {
	GsBase
	Radius float64 `json:"radius"`
}

func NewSphere() *Sphere {
	return &Sphere{
		GsBase: GsBase{Type: "GIM/GS/Sphere"},
	}
}

// RotationalEllipsoidObject represents a rotational ellipsoid
type RotationalEllipsoid struct {
	GsBase
	PolarRadius      float64 `json:"polarRadius"`
	EquatorialRadius float64 `json:"equatorialRadius"`
	Height           float64 `json:"height"`
}

func NewRotationalEllipsoid() *RotationalEllipsoid {
	return &RotationalEllipsoid{
		GsBase: GsBase{Type: "GIM/GS/RotationalEllipsoid"},
	}
}

// CuboidObject represents a cuboid
type Cuboid struct {
	GsBase
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func NewCuboid() *Cuboid {
	return &Cuboid{
		GsBase: GsBase{Type: "GIM/GS/Cuboid"},
	}
}

// DiamondFrustumObject represents a diamond frustum
type DiamondFrustum struct {
	GsBase
	TopDiag1    float64 `json:"topDiag1"`
	TopDiag2    float64 `json:"topDiag2"`
	BottomDiag1 float64 `json:"bottomDiag1"`
	BottomDiag2 float64 `json:"bottomDiag2"`
	Height      float64 `json:"height"`
}

func NewDiamondFrustum() *DiamondFrustum {
	return &DiamondFrustum{
		GsBase: GsBase{Type: "GIM/GS/DiamondFrustum"},
	}
}

// OffsetRectangularTableObject represents an offset rectangular table
type OffsetRectangularTable struct {
	GsBase
	TopLength    float64 `json:"topLength"`
	TopWidth     float64 `json:"topWidth"`
	BottomLength float64 `json:"bottomLength"`
	BottomWidth  float64 `json:"bottomWidth"`
	Height       float64 `json:"height"`
	XOffset      float64 `json:"xOffset"`
	YOffset      float64 `json:"yOffset"`
}

func NewOffsetRectangularTable() *OffsetRectangularTable {
	return &OffsetRectangularTable{
		GsBase: GsBase{Type: "GIM/GS/OffsetRectangularTable"},
	}
}

// CylinderObject represents a cylinder
type Cylinder struct {
	GsBase
	Radius float64 `json:"radius"`
	Height float64 `json:"height"`
}

func NewCylinder() *Cylinder {
	return &Cylinder{
		GsBase: GsBase{Type: "GIM/GS/Cylinder"},
	}
}

// SharpBentCylinderObject represents a sharp bent cylinder
type SharpBentCylinder struct {
	GsBase
	Radius    float64 `json:"radius"`
	Length    float64 `json:"length"`
	BendAngle float64 `json:"bendAngle"`
}

func NewSharpBentCylinder() *SharpBentCylinder {
	return &SharpBentCylinder{
		GsBase: GsBase{Type: "GIM/GS/SharpBentCylinder"},
	}
}

// TruncatedConeObject represents a truncated cone
type TruncatedCone struct {
	GsBase
	TopRadius    float64 `json:"topRadius"`
	BottomRadius float64 `json:"bottomRadius"`
	Height       float64 `json:"height"`
}

func NewTruncatedCone() *TruncatedCone {
	return &TruncatedCone{
		GsBase: GsBase{Type: "GIM/GS/TruncatedCone"},
	}
}

// EccentricTruncatedConeObject represents an eccentric truncated cone
type EccentricTruncatedCone struct {
	GsBase
	TopRadius    float64 `json:"topRadius"`
	BottomRadius float64 `json:"bottomRadius"`
	Height       float64 `json:"height"`
	TopXOffset   float64 `json:"topXOffset"`
	TopYOffset   float64 `json:"topYOffset"`
}

func NewEccentricTruncatedCone() *EccentricTruncatedCone {
	return &EccentricTruncatedCone{
		GsBase: GsBase{Type: "GIM/GS/EccentricTruncatedCone"},
	}
}

// RingObject represents a ring
type Ring struct {
	GsBase
	RingRadius float64 `json:"ringRadius"`
	TubeRadius float64 `json:"tubeRadius"`
	Angle      float64 `json:"angle"`
}

func NewRing() *Ring {
	return &Ring{
		GsBase: GsBase{Type: "GIM/GS/Ring"},
	}
}

// RectangularRingObject represents a rectangular ring
type RectangularRing struct {
	GsBase
	TubeRadius   float64 `json:"tubeRadius"`
	FilletRadius float64 `json:"filletRadius"`
	Length       float64 `json:"length"`
	Width        float64 `json:"width"`
}

func NewRectangularRing() *RectangularRing {
	return &RectangularRing{
		GsBase: GsBase{Type: "GIM/GS/RectangularRing"},
	}
}

// EllipticRingObject represents an elliptic ring
type EllipticRing struct {
	GsBase
	TubeRadius  float64 `json:"tubeRadius"`
	MajorRadius float64 `json:"majorRadius"`
	MinorRadius float64 `json:"minorRadius"`
}

func NewEllipticRing() *EllipticRing {
	return &EllipticRing{
		GsBase: GsBase{Type: "GIM/GS/EllipticRing"},
	}
}

// CircularGasketObject represents a circular gasket
type CircularGasket struct {
	GsBase
	OuterRadius float64 `json:"outerRadius"`
	InnerRadius float64 `json:"innerRadius"`
	Height      float64 `json:"height"`
	Angle       float64 `json:"angle"`
}

func NewCircularGasket() *CircularGasket {
	return &CircularGasket{
		GsBase: GsBase{Type: "GIM/GS/CircularGasket"},
	}
}

// TableGasketObject represents a table gasket
type TableGasket struct {
	GsBase
	TopRadius   float64 `json:"topRadius"`
	OuterRadius float64 `json:"outerRadius"`
	InnerRadius float64 `json:"innerRadius"`
	Height      float64 `json:"height"`
	Angle       float64 `json:"angle"`
}

func NewTableGasket() *TableGasket {
	return &TableGasket{
		GsBase: GsBase{Type: "GIM/GS/TableGasket"},
	}
}

// SquareGasketObject represents a square gasket
type SquareGasket struct {
	GsBase
	OuterLength float64 `json:"outerLength"`
	OuterWidth  float64 `json:"outerWidth"`
	InnerLength float64 `json:"innerLength"`
	InnerWidth  float64 `json:"innerWidth"`
	Height      float64 `json:"height"`
	CornerType  int     `json:"cornerType"`
	CornerParam float64 `json:"cornerParam"`
}

func NewSquareGasket() *SquareGasket {
	return &SquareGasket{
		GsBase: GsBase{Type: "GIM/GS/SquareGasket"},
	}
}

// StretchedBodyObject represents a stretched body
type StretchedBody struct {
	GsBase
	Points [][3]float64 `json:"points"`
	Normal [3]float64   `json:"normal"`
	Length float64      `json:"length"`
}

func NewStretchedBody() *StretchedBody {
	return &StretchedBody{
		GsBase: GsBase{Type: "GIM/GS/StretchedBody"},
	}
}

// PorcelainBushingObject represents a porcelain bushing
type PorcelainBushing struct {
	GsBase
	Height           float64 `json:"height"`
	Radius           float64 `json:"radius"`
	BigSkirtRadius   float64 `json:"bigSkirtRadius"`
	SmallSkirtRadius float64 `json:"smallSkirtRadius"`
	Count            int     `json:"count"`
}

func NewPorcelainBushing() *PorcelainBushing {
	return &PorcelainBushing{
		GsBase: GsBase{Type: "GIM/GS/PorcelainBushing"},
	}
}

// ConePorcelainBushingObject represents a cone porcelain bushing
type ConePorcelainBushing struct {
	GsBase
	Height             float64 `json:"height"`
	BottomRadius       float64 `json:"bottomRadius"`
	TopRadius          float64 `json:"topRadius"`
	BottomSkirtRadius1 float64 `json:"bottomSkirtRadius1"`
	BottomSkirtRadius2 float64 `json:"bottomSkirtRadius2"`
	TopSkirtRadius1    float64 `json:"topSkirtRadius1"`
	TopSkirtRadius2    float64 `json:"topSkirtRadius2"`
	Count              int     `json:"count"`
}

func NewConePorcelainBushing() *ConePorcelainBushing {
	return &ConePorcelainBushing{
		GsBase: GsBase{Type: "GIM/GS/ConePorcelainBushing"},
	}
}

// InsulatorStringObject represents an insulator string
type InsulatorString struct {
	GsBase
	Count            int     `json:"count"`
	Spacing          float64 `json:"spacing"`
	InsulatorCount   int     `json:"insulatorCount"`
	Height           float64 `json:"height"`
	BigSkirtRadius   float64 `json:"bigSkirtRadius"`
	SmallSkirtRadius float64 `json:"smallSkirtRadius"`
	Radius           float64 `json:"radius"`
	FrontLength      float64 `json:"frontLength"`
	BackLength       float64 `json:"backLength"`
	SplitCount       int     `json:"splitCount"`
}

func NewInsulatorString() *InsulatorString {
	return &InsulatorString{
		GsBase: GsBase{Type: "GIM/GS/InsulatorString"},
	}
}

// VTypeInsulatorObject represents a V-type insulator
type VTypeInsulator struct {
	GsBase
	FrontSpacing     float64 `json:"frontSpacing"`
	BackSpacing      float64 `json:"backSpacing"`
	InsulatorCount   int     `json:"insulatorCount"`
	Height           float64 `json:"height"`
	Radius           float64 `json:"radius"`
	BigSkirtRadius   float64 `json:"bigSkirtRadius"`
	SmallSkirtRadius float64 `json:"smallSkirtRadius"`
	FrontLength      float64 `json:"frontLength"`
	BackLength       float64 `json:"backLength"`
	SplitCount       int     `json:"splitCount"`
}

func NewVTypeInsulator() *VTypeInsulator {
	return &VTypeInsulator{
		GsBase: GsBase{Type: "GIM/GS/VTypeInsulator"},
	}
}

// TerminalBlockObject represents a terminal block
type TerminalBlock struct {
	GsBase
	Length        float64 `json:"length"`
	Width         float64 `json:"width"`
	Thickness     float64 `json:"thickness"`
	ChamferLength float64 `json:"chamferLength"`
	ColumnSpacing float64 `json:"columnSpacing"`
	RowSpacing    float64 `json:"rowSpacing"`
	HoleRadius    float64 `json:"holeRadius"`
	ColumnCount   int     `json:"columnCount"`
	RowCount      int     `json:"rowCount"`
	BottomOffset  float64 `json:"bottomOffset"`
}

func NewTerminalBlock() *TerminalBlock {
	return &TerminalBlock{
		GsBase: GsBase{Type: "GIM/GS/TerminalBlock"},
	}
}

// RectangularHolePlateObject represents a rectangular hole plate
type RectangularHolePlate struct {
	GsBase
	Length        float64 `json:"length"`
	Width         float64 `json:"width"`
	Thickness     float64 `json:"thickness"`
	ColumnSpacing float64 `json:"columnSpacing"`
	RowSpacing    float64 `json:"rowSpacing"`
	ColumnCount   int     `json:"columnCount"`
	RowCount      int     `json:"rowCount"`
	HasMiddleHole bool    `json:"hasMiddleHole"`
	HoleDiameter  float64 `json:"holeDiameter"`
}

func NewRectangularHolePlate() *RectangularHolePlate {
	return &RectangularHolePlate{
		GsBase: GsBase{Type: "GIM/GS/RectangularHolePlate"},
	}
}

// CircularFixedPlateObject represents a circular fixed plate
type CircularFixedPlate struct {
	GsBase
	Length        float64 `json:"length"`
	Width         float64 `json:"width"`
	Thickness     float64 `json:"thickness"`
	RingRadius    float64 `json:"ringRadius"`
	HoleCount     int     `json:"holeCount"`
	HasMiddleHole bool    `json:"hasMiddleHole"`
	HoleDiameter  float64 `json:"holeDiameter"`
}

func NewCircularFixedPlate() *CircularFixedPlate {
	return &CircularFixedPlate{
		GsBase: GsBase{Type: "GIM/GS/CircularFixedPlate"},
	}
}

// WireObject represents a wire
type Wire struct {
	GsBase
	StartPoint [3]float64   `json:"startPoint"`
	EndPoint   [3]float64   `json:"endPoint"`
	StartDir   [3]float64   `json:"startDir"`
	EndDir     [3]float64   `json:"endDir"`
	Sag        float64      `json:"sag"`
	Diameter   float64      `json:"diameter"`
	FitPoints  [][3]float64 `json:"fitPoints"`
}

func NewWire() *Wire {
	return &Wire{
		GsBase: GsBase{Type: "GIM/GS/Wire"},
	}
}

// CableObject represents a cable
type Cable struct {
	GsBase
	StartPoint       [3]float64   `json:"startPoint"`
	EndPoint         [3]float64   `json:"endPoint"`
	InflectionPoints [][3]float64 `json:"inflectionPoints"`
	Radii            []float64    `json:"radii"`
	Diameter         float64      `json:"diameter"`
}

func NewCable() *Cable {
	return &Cable{
		GsBase: GsBase{Type: "GIM/GS/Cable"},
	}
}

type CurveType string

const (
	CurveTypeLine   CurveType = "LINE"
	CurveTypeArc    CurveType = "ARC"
	CurveTypeSpline CurveType = "SPLINE"
)

// CurveCableObject represents a curve cable
type CurveCable struct {
	GsBase
	ControlPoints [][][3]float64 `json:"controlPoints"`
	CurveTypes    []CurveType    `json:"curveTypes"`
	Diameter      float64        `json:"diameter"`
}

func NewCurveCable() *CurveCable {
	return &CurveCable{
		GsBase: GsBase{Type: "GIM/GS/CurveCable"},
	}
}

// AngleSteelObject represents angle steel
type AngleSteel struct {
	GsBase
	L1     float64 `json:"L1"`
	L2     float64 `json:"L2"`
	X      float64 `json:"X"`
	Length float64 `json:"length"`
}

func NewAngleSteel() *AngleSteel {
	return &AngleSteel{
		GsBase: GsBase{Type: "GIM/GS/AngleSteel"},
	}
}

// IShapedSteelObject represents I-shaped steel
type IShapedSteel struct {
	GsBase
	Height          float64 `json:"height"`
	FlangeWidth     float64 `json:"flangeWidth"`
	WebThickness    float64 `json:"webThickness"`
	FlangeThickness float64 `json:"flangeThickness"`
	Length          float64 `json:"length"`
}

func NewIShapedSteel() *IShapedSteel {
	return &IShapedSteel{
		GsBase: GsBase{Type: "GIM/GS/IShapedSteel"},
	}
}

// ChannelSteelObject represents channel steel
type ChannelSteel struct {
	GsBase
	Height          float64 `json:"height"`
	FlangeWidth     float64 `json:"flangeWidth"`
	WebThickness    float64 `json:"webThickness"`
	FlangeThickness float64 `json:"flangeThickness"`
	Length          float64 `json:"length"`
}

func NewChannelSteel() *ChannelSteel {
	return &ChannelSteel{
		GsBase: GsBase{Type: "GIM/GS/ChannelSteel"},
	}
}

// TSteelObject represents T-shaped steel
type TSteel struct {
	GsBase
	Height          float64 `json:"height"`
	Width           float64 `json:"width"`
	WebThickness    float64 `json:"webThickness"`
	FlangeThickness float64 `json:"flangeThickness"`
	Length          float64 `json:"length"`
}

func NewTSteel() *TSteel {
	return &TSteel{
		GsBase: GsBase{Type: "GIM/GS/TSteel"},
	}
}

type Shape interface {
	GetType() string
}

func Unmarshal(ty string, bt []byte) (Shape, error) {
	switch ty {
	case "GIM/GS/Sphere":
		shape := Sphere{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/RotationalEllipsoid":
		shape := RotationalEllipsoid{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/Cuboid":
		shape := Cuboid{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/DiamondFrustum":
		shape := DiamondFrustum{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/OffsetRectangularTable":
		shape := OffsetRectangularTable{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/Cylinder":
		shape := Cylinder{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/SharpBentCylinder":
		shape := SharpBentCylinder{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/TruncatedCone":
		shape := TruncatedCone{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/EccentricTruncatedCone":
		shape := EccentricTruncatedCone{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/Ring":
		shape := Ring{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/RectangularRing":
		shape := RectangularRing{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/EllipticRing":
		shape := EllipticRing{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/CircularGasket":
		shape := CircularGasket{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/TableGasket":
		shape := TableGasket{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/SquareGasket":
		shape := SquareGasket{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/StretchedBody":
		shape := StretchedBody{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/PorcelainBushing":
		shape := PorcelainBushing{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/ConePorcelainBushing":
		shape := ConePorcelainBushing{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/InsulatorString":
		shape := InsulatorString{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/VTypeInsulator":
		shape := VTypeInsulator{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/TerminalBlock":
		shape := TerminalBlock{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/RectangularHolePlate":
		shape := RectangularHolePlate{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/CircularFixedPlate":
		shape := CircularFixedPlate{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/Wire":
		shape := Wire{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/Cable":
		shape := Cable{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/CurveCable":
		shape := CurveCable{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/AngleSteel":
		shape := AngleSteel{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/IShapedSteel":
		shape := IShapedSteel{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/ChannelSteel":
		shape := ChannelSteel{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "GIM/GS/TSteel":
		shape := TSteel{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	default:
		return nil, fmt.Errorf("invalid type: %s", ty)
	}
}
