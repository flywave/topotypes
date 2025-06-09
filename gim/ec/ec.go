package ec

import (
	"encoding/json"
	"fmt"
)

const Major = "EC"

type Point [3]float64
type Point2 [2]float64

type EcBase struct {
	Version string
	Type    string `json:"type"`
}

func (b *EcBase) GetType() string {
	return b.Type
}

// Enum types with conversion functions
type CableClampType string

const (
	CableClampSingle         CableClampType = "SINGLE"
	CableClampLinear         CableClampType = "LINEAR"
	CableClampContactTriple  CableClampType = "CONTACT_TRIPLE"
	CableClampSeparateTriple CableClampType = "SEPARATE_TRIPLE"
)

func (c CableClampType) ToInt() int {
	switch c {
	case CableClampSingle:
		return 1
	case CableClampLinear:
		return 2
	case CableClampContactTriple:
		return 3
	case CableClampSeparateTriple:
		return 4
	default:
		return 0
	}
}

type CableTerminalSort string

const (
	CableTerminalOutdoor CableTerminalSort = "OUTDOOR"
	CableTerminalGIS     CableTerminalSort = "GIS"
	CableTerminalDry     CableTerminalSort = "DRY"
)

func (c CableTerminalSort) ToInt() int {
	switch c {
	case CableTerminalOutdoor:
		return 1
	case CableTerminalGIS:
		return 2
	case CableTerminalDry:
		return 3
	default:
		return 0
	}
}

type SectionType string

const (
	SectionRectangular SectionType = "RECTANGULAR"
	SectionHorseshoe   SectionType = "HORSESHOE"
	SectionCircular    SectionType = "CIRCULAR"
)

func (s SectionType) ToInt() int {
	switch s {
	case SectionRectangular:
		return 1
	case SectionHorseshoe:
		return 2
	case SectionCircular:
		return 3
	default:
		return 0
	}
}

type ThreeWayWellType string

const (
	WellUndergroundTunnel ThreeWayWellType = "UNDERGROUND_TUNNEL"
	WellOpenCutTunnel     ThreeWayWellType = "OPEN_CUT_TUNNEL"
	WellWorkingWell       ThreeWayWellType = "WORKING_WELL"
)

func (w ThreeWayWellType) ToInt() int {
	switch w {
	case WellUndergroundTunnel:
		return 3
	case WellOpenCutTunnel:
		return 2
	case WellWorkingWell:
		return 1
	default:
		return 1
	}
}

type CornerType string

const (
	CornerRounded CornerType = "ROUNDED"
	CornerAngled  CornerType = "ANGLED"
)

func (c CornerType) ToInt() int {
	switch c {
	case CornerRounded:
		return 1
	case CornerAngled:
		return 2
	default:
		return 0
	}
}

type ShaftType string

const (
	ShaftCircular    ShaftType = "CIRCULAR"
	ShaftRectangular ShaftType = "RECTANGULAR"
)

func (s ShaftType) ToInt() int {
	switch s {
	case ShaftCircular:
		return 1
	case ShaftRectangular:
		return 2
	default:
		return 0
	}
}

type PipeType string

const (
	PipeNormal PipeType = "NORMAL"
	PipePull   PipeType = "PULL"
)

func (p PipeType) ToInt() int {
	switch p {
	case PipeNormal:
		return 1
	case PipePull:
		return 2
	default:
		return 0
	}
}

type TunnelStyle string

const (
	TunnelRectangular TunnelStyle = "RECTANGULAR"
	TunnelHorseshoe   TunnelStyle = "HORSESHOE"
	TunnelCircular    TunnelStyle = "CIRCULAR"
)

func (t TunnelStyle) ToInt() int {
	switch t {
	case TunnelRectangular:
		return 1
	case TunnelHorseshoe:
		return 2
	case TunnelCircular:
		return 3
	default:
		return 0
	}
}

type CableTrayStyle string

const (
	CableTrayArch CableTrayStyle = "ARCH"
	CableTrayBeam CableTrayStyle = "BEAM"
)

func (c CableTrayStyle) ToInt() int {
	switch c {
	case CableTrayArch:
		return 1
	case CableTrayBeam:
		return 2
	default:
		return 0
	}
}

type ManholeStyle string

const (
	ManholeCircular    ManholeStyle = "CIRCULAR"
	ManholeRectangular ManholeStyle = "RECTANGULAR"
)

func (m ManholeStyle) ToInt() int {
	switch m {
	case ManholeCircular:
		return 1
	case ManholeRectangular:
		return 2
	default:
		return 0
	}
}

type PartitionBoardStyle string

const (
	PartitionCircular    PartitionBoardStyle = "CIRCULAR"
	PartitionRectangular PartitionBoardStyle = "RECTANGULAR"
)

func (p PartitionBoardStyle) ToInt() int {
	switch p {
	case PartitionCircular:
		return 1
	case PartitionRectangular:
		return 2
	default:
		return 0
	}
}

type PipeSupportStyle string

const (
	PipeSupportSingleSided PipeSupportStyle = "SINGLE_SIDED"
	PipeSupportDoubleSided PipeSupportStyle = "DOUBLE_SIDED"
)

func (p PipeSupportStyle) ToInt() int {
	switch p {
	case PipeSupportSingleSided:
		return 1
	case PipeSupportDoubleSided:
		return 2
	default:
		return 0
	}
}

type CoverPlateStyle string

const (
	CoverPlateRectangular CoverPlateStyle = "RECTANGULAR"
	CoverPlateSector      CoverPlateStyle = "SECTOR"
)

func (c CoverPlateStyle) ToInt() int {
	switch c {
	case CoverPlateRectangular:
		return 0
	case CoverPlateSector:
		return 1
	default:
		return 0
	}
}

type ChannelPointType string

const (
	ChannelPointLine ChannelPointType = "LINE"
	ChannelPointArc  ChannelPointType = "ARC"
)

func (c ChannelPointType) ToInt() int {
	switch c {
	case ChannelPointLine:
		return 0
	case ChannelPointArc:
		return 1
	default:
		return 0
	}
}

// Struct definitions with New functions (Type removed from individual structs)
type ChannelPoint struct {
	Position Point
	Type     ChannelPointType
}

type CableWire struct {
	EcBase
	Points          []Point `json:"points"`
	OutsideDiameter float64 `json:"outsideDiameter"`
}

func NewCableWire() *CableWire {
	return &CableWire{
		EcBase: EcBase{Type: "GIM/EC/CableWire"},
	}
}

type CableJoint struct {
	EcBase
	Length         float64 `json:"length"`
	OuterDiameter  float64 `json:"outerDiameter"`
	TerminalLength float64 `json:"terminalLength"`
	InnerDiameter  float64 `json:"innerDiameter"`
}

func NewCableJoint() *CableJoint {
	return &CableJoint{
		EcBase: EcBase{Type: "GIM/EC/CableJoint"},
	}
}

// [Rest of the struct definitions follow the same pattern...]
// Each struct now embeds EcBase and has Type initialized in its New function
// All enum types have ToInt() conversion methods

// Example of another struct (all others would follow this pattern)
type OpticalFiberBox struct {
	EcBase
	Length float64 `json:"length"`
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

func NewOpticalFiberBox() *OpticalFiberBox {
	return &OpticalFiberBox{
		EcBase: EcBase{Type: "GIM/EC/OpticalFiberBox"},
	}
}

type CableTerminal struct {
	EcBase
	Sort                     CableTerminalSort `json:"sort"`
	Height                   float64           `json:"height"`
	TopDiameter              float64           `json:"topDiameter"`
	BottomDiameter           float64           `json:"bottomDiameter"`
	TailDiameter             float64           `json:"tailDiameter"`
	TailHeight               float64           `json:"tailHeight"`
	SkirtCount               int               `json:"skirtCount"`
	UpperSkirtTopDiameter    float64           `json:"upperSkirtTopDiameter"`
	UpperSkirtBottomDiameter float64           `json:"upperSkirtBottomDiameter"`
	LowerSkirtTopDiameter    float64           `json:"lowerSkirtTopDiameter"`
	LowerSkirtBottomDiameter float64           `json:"lowerSkirtBottomDiameter"`
	SkirtSectionHeight       float64           `json:"skirtSectionHeight"`
	UpperTerminalLength      float64           `json:"upperTerminalLength"`
	UpperTerminalDiameter    float64           `json:"upperTerminalDiameter"`
	LowerTerminalLength      float64           `json:"lowerTerminalLength"`
	LowerTerminalDiameter    float64           `json:"lowerTerminalDiameter"`
	Hole1Diameter            float64           `json:"hole1Diameter"`
	Hole2Diameter            float64           `json:"hole2Diameter"`
	Hole1Distance            float64           `json:"hole1Distance"`
	HoleSpacing              float64           `json:"holeSpacing"`
	FlangeHoleDiameter       float64           `json:"flangeHoleDiameter"`
	FlangeHoleSpacing        float64           `json:"flangeHoleSpacing"`
	FlangeWidth              float64           `json:"flangeWidth"`
	FlangeCenterHoleRadius   float64           `json:"flangeCenterHoleRadius"`
	FlangeChamferRadius      float64           `json:"flangeChamferRadius"`
	FlangeOpeningWidth       float64           `json:"flangeOpeningWidth"`
	FlangeBoltHeight         float64           `json:"flangeBoltHeight"`
}

func NewCableTerminal() *CableTerminal {
	return &CableTerminal{
		EcBase: EcBase{Type: "GIM/EC/CableTerminal"},
	}
}

type CableAccessoryType string

const (
	CableAccessoryDirectGround      CableAccessoryType = "DIRECT_GROUND"
	CableAccessoryProtectiveGround  CableAccessoryType = "PROTECTIVE_GROUND"
	CableAccessoryCrossInterconnect CableAccessoryType = "CROSS_INTERCONNECT"
)

func (c CableAccessoryType) ToInt() int {
	switch c {
	case CableAccessoryDirectGround:
		return 1
	case CableAccessoryProtectiveGround:
		return 2
	case CableAccessoryCrossInterconnect:
		return 3
	default:
		return 0
	}
}

type CableAccessory struct {
	EcBase
	CableAccessoryType CableAccessoryType `json:"cableAccessoryType"`
	Length             float64            `json:"length"`
	Width              float64            `json:"width"`
	Height             float64            `json:"height"`
	PortCount          int                `json:"portCount"`
	PortDiameter       float64            `json:"portDiameter"`
	PortSpacing        float64            `json:"portSpacing"`
	BackPanelDistance  float64            `json:"backPanelDistance"`
	SidePanelDistance  float64            `json:"sidePanelDistance"`
}

func NewCableAccessory() *CableAccessory {
	return &CableAccessory{
		EcBase: EcBase{Type: "GIM/EC/CableAccessory"},
	}
}

type CableBracket struct {
	EcBase
	Length            float64 `json:"length"`
	RootHeight        float64 `json:"rootHeight"`
	RootWidth         float64 `json:"rootWidth"`
	Width             float64 `json:"width"`
	TopThickness      float64 `json:"topThickness"`
	RootThickness     float64 `json:"rootThickness"`
	ColumnMountPoints []Point `json:"columnMountPoints"`
	ClampMountPoints  []Point `json:"clampMountPoints"`
}

func NewCableBracket() *CableBracket {
	return &CableBracket{
		EcBase: EcBase{Type: "GIM/EC/CableBracket"},
	}
}

type CableClamp struct {
	EcBase
	ClampType CableClampType `json:"clampType"`
	Diameter  float64        `json:"diameter"`
	Thickness float64        `json:"thickness"`
	Width     float64        `json:"width"`
}

func NewCableClamp() *CableClamp {
	return &CableClamp{
		EcBase: EcBase{Type: "GIM/EC/CableClamp"},
	}
}

type CablePole struct {
	EcBase
	Specification  string  `json:"specification"`
	Length         float64 `json:"length"`
	Radius         float64 `json:"radius"`
	ArcAngle       float64 `json:"arcAngle"`
	Width          float64 `json:"width"`
	FixedLegLength float64 `json:"fixedLegLength"`
	FixedLegWidth  float64 `json:"fixedLegWidth"`
	Thickness      float64 `json:"thickness"`
	MountPoints    []Point `json:"mountPoints"`
}

func NewCablePole() *CablePole {
	return &CablePole{
		EcBase: EcBase{Type: "GIM/EC/CablePole"},
	}
}

type GroundFlatIron struct {
	EcBase
	Length    float64 `json:"length"`
	Height    float64 `json:"height"`
	Thickness float64 `json:"thickness"`
}

func NewGroundFlatIron() *GroundFlatIron {
	return &GroundFlatIron{
		EcBase: EcBase{Type: "GIM/EC/GroundFlatIron"},
	}
}

type EmbeddedPart struct {
	EcBase
	Length         float64 `json:"length"`
	Radius         float64 `json:"radius"`
	Height         float64 `json:"height"`
	MaterialRadius float64 `json:"materialRadius"`
	LowerLength    float64 `json:"lowerLength"`
}

func NewEmbeddedPart() *EmbeddedPart {
	return &EmbeddedPart{
		EcBase: EcBase{Type: "GIM/EC/EmbeddedPart"},
	}
}

type UShapedRing struct {
	EcBase
	Thickness float64 `json:"thickness"`
	Height    float64 `json:"height"`
	Radius    float64 `json:"radius"`
	Length    float64 `json:"length"`
}

func NewUShapedRing() *UShapedRing {
	return &UShapedRing{
		EcBase: EcBase{Type: "GIM/EC/UShapedRing"},
	}
}

type LiftingEye struct {
	EcBase
	Height       float64 `json:"height"`
	RingRadius   float64 `json:"ringRadius"`
	PipeDiameter float64 `json:"pipeDiameter"`
}

func NewLiftingEye() *LiftingEye {
	return &LiftingEye{
		EcBase: EcBase{Type: "GIM/EC/LiftingEye"},
	}
}

type CornerWell struct {
	EcBase
	LeftLength       float64 `json:"leftLength"`
	RightLength      float64 `json:"rightLength"`
	Width            float64 `json:"width"`
	Height           float64 `json:"height"`
	TopThickness     float64 `json:"topThickness"`
	BottomThickness  float64 `json:"bottomThickness"`
	WallThickness    float64 `json:"wallThickness"`
	Angle            float64 `json:"angle"`
	CornerRadius     float64 `json:"cornerRadius"`
	CushionExtension float64 `json:"cushionExtension"`
	CushionThickness float64 `json:"cushionThickness"`
}

func NewCornerWell() *CornerWell {
	return &CornerWell{
		EcBase: EcBase{Type: "GIM/EC/CornerWell"},
	}
}

type TunnelWellType string

const (
	TunnelWellStraight       TunnelWellType = "STRAIGHT"
	TunnelWellStraightTunnel TunnelWellType = "STRAIGHT_TUNNEL"
)

func (t TunnelWellType) ToInt() int {
	switch t {
	case TunnelWellStraight:
		return 1
	case TunnelWellStraightTunnel:
		return 2
	}
	return 1
}

type TunnelWell struct {
	EcBase
	WellType           TunnelWellType `json:"wellType"`
	Length             float64        `json:"length"`
	Width              float64        `json:"width"`
	Height             float64        `json:"height"`
	TopThickness       float64        `json:"topThickness"`
	BottomThickness    float64        `json:"bottomThickness"`
	OuterWallThickness float64        `json:"outerWallThickness"`
	CushionExtension   float64        `json:"cushionExtension"`
	CushionThickness   float64        `json:"cushionThickness"`
	LeftSectionType    SectionType    `json:"leftSectionType"`
	LeftLength         float64        `json:"leftLength"`
	LeftWidth          float64        `json:"leftWidth"`
	LeftHeight         float64        `json:"leftHeight"`
	LeftArcHeight      float64        `json:"leftArcHeight"`
	RightSectionType   SectionType    `json:"rightSectionType"`
	RightLength        float64        `json:"rightLength"`
	RightWidth         float64        `json:"rightWidth"`
	RightHeight        float64        `json:"rightHeight"`
	RightArcHeight     float64        `json:"rightArcHeight"`
	Radius             float64        `json:"radius"`
	InnerWallThickness float64        `json:"innerWallThickness"`
}

func NewTunnelWell() *TunnelWell {
	return &TunnelWell{
		EcBase: EcBase{Type: "GIM/EC/TunnelWell"},
	}
}

type ThreeWayWell struct {
	EcBase
	WellType               ThreeWayWellType `json:"wellType"`
	CornerType             CornerType       `json:"cornerType"`
	ShaftType              ShaftType        `json:"shaftType"`
	Length                 float64          `json:"length"`
	Width                  float64          `json:"width"`
	Height                 float64          `json:"height"`
	ShaftRadius            float64          `json:"shaftRadius"`
	CornerRadius           float64          `json:"cornerRadius"`
	CornerLength           float64          `json:"cornerLength"`
	CornerWidth            float64          `json:"cornerWidth"`
	BranchLength           float64          `json:"branchLength"`
	BranchLeftLength       float64          `json:"branchLeftLength"`
	BranchWidth            float64          `json:"branchWidth"`
	TopThickness           float64          `json:"topThickness"`
	BottomThickness        float64          `json:"bottomThickness"`
	LeftSectionStyle       SectionType      `json:"leftSectionStyle"`
	LeftSectionLength      float64          `json:"leftSectionLength"`
	LeftSectionWidth       float64          `json:"leftSectionWidth"`
	LeftSectionHeight      float64          `json:"leftSectionHeight"`
	LeftSectionArcHeight   float64          `json:"leftSectionArcHeight"`
	RightSectionStyle      SectionType      `json:"rightSectionStyle"`
	RightSectionLength     float64          `json:"rightSectionLength"`
	RightSectionWidth      float64          `json:"rightSectionWidth"`
	RightSectionHeight     float64          `json:"rightSectionHeight"`
	RightSectionArcHeight  float64          `json:"rightSectionArcHeight"`
	BranchSectionStyle     SectionType      `json:"branchSectionStyle"`
	BranchSectionLength    float64          `json:"branchSectionLength"`
	BranchSectionWidth     float64          `json:"branchSectionWidth"`
	BranchSectionHeight    float64          `json:"branchSectionHeight"`
	BranchSectionArcHeight float64          `json:"branchSectionArcHeight"`
	OuterWallThickness     float64          `json:"outerWallThickness"`
	InnerWallThickness     float64          `json:"innerWallThickness"`
	IsDoubleShaft          bool             `json:"isDoubleShaft"`
	DoubleShaftSpacing     float64          `json:"doubleShaftSpacing"`
	OuterWallExtension     float64          `json:"outerWallExtension"`
	InnerWallExtension     float64          `json:"innerWallExtension"`
	CushionExtension       float64          `json:"cushionExtension"`
	CushionThickness       float64          `json:"cushionThickness"`
	InnerBottomThickness   float64          `json:"innerBottomThickness"`
	OuterBottomThickness   float64          `json:"outerBottomThickness"`
	Angle                  float64          `json:"angle"`
}

func NewThreeWayWell() *ThreeWayWell {
	return &ThreeWayWell{
		EcBase: EcBase{Type: "GIM/EC/ThreeWayWell"},
	}
}

type FourWayWellType string

const (
	FourWayWellTypeWorking           FourWayWellType = "WORKING_WELL"
	FourWayWellTypeOpenCutTunnel     FourWayWellType = "OPEN_CUT_TUNNEL"
	FourWayWellTypeUnderGroundTunnel FourWayWellType = "UNDERGROUND_TUNNEL"
)

func (t FourWayWellType) ToInt() int {
	switch t {
	case FourWayWellTypeWorking:
		return 1
	case FourWayWellTypeOpenCutTunnel:
		return 2
	case FourWayWellTypeUnderGroundTunnel:
		return 3
	default:
		return 1
	}
}

type FourWayWell struct {
	EcBase
	WellType                FourWayWellType `json:"wellType"`
	Length                  float64         `json:"length"`
	Width                   float64         `json:"width"`
	Height                  float64         `json:"height"`
	CornerStyle             CornerType      `json:"cornerStyle"`
	CornerRadius            float64         `json:"cornerRadius"`
	BranchLength            float64         `json:"branchLength"`
	BranchWidth             float64         `json:"branchWidth"`
	TopThickness            float64         `json:"topThickness"`
	BottomThickness         float64         `json:"bottomThickness"`
	OuterWallThickness      float64         `json:"outerWallThickness"`
	InnerWallThickness      float64         `json:"innerWallThickness"`
	CushionExtension        float64         `json:"cushionExtension"`
	CushionThickness        float64         `json:"cushionThickness"`
	LeftSectionStyle        SectionType     `json:"leftSectionStyle"`
	LeftSectionLength       float64         `json:"leftSectionLength"`
	LeftSectionWidth        float64         `json:"leftSectionWidth"`
	LeftSectionHeight       float64         `json:"leftSectionHeight"`
	LeftSectionArcHeight    float64         `json:"leftSectionArcHeight"`
	RightSectionStyle       SectionType     `json:"rightSectionStyle"`
	RightSectionLength      float64         `json:"rightSectionLength"`
	RightSectionWidth       float64         `json:"rightSectionWidth"`
	RightSectionHeight      float64         `json:"rightSectionHeight"`
	RightSectionArcHeight   float64         `json:"rightSectionArcHeight"`
	BranchSection1Style     SectionType     `json:"branchSection1Style"`
	BranchSection1Length    float64         `json:"branchSection1Length"`
	BranchSection1Width     float64         `json:"branchSection1Width"`
	BranchSection1Height    float64         `json:"branchSection1Height"`
	BranchSection1ArcHeight float64         `json:"branchSection1ArcHeight"`
	BranchSection2Style     SectionType     `json:"branchSection2Style"`
	BranchSection2Length    float64         `json:"branchSection2Length"`
	BranchSection2Width     float64         `json:"branchSection2Width"`
	BranchSection2Height    float64         `json:"branchSection2Height"`
	BranchSection2ArcHeight float64         `json:"branchSection2ArcHeight"`
	ShaftRadius             float64         `json:"shaftRadius"`
	CornerLength            float64         `json:"cornerLength"`
	CornerWidth             float64         `json:"cornerWidth"`
}

func NewFourWayWell() *FourWayWell {
	return &FourWayWell{
		EcBase: EcBase{Type: "GIM/EC/FourWayWell"},
	}
}

type ChannelPointIndex struct {
	Index int              `json:"end"`
	Type  ChannelPointType `json:"type"`
}
type PipeRow struct {
	EcBase
	PipeType              PipeType            `json:"pipeType"`
	HasEnclosure          bool                `json:"hasEnclosure"`
	EnclosureWidth        float64             `json:"enclosureWidth"`
	EnclosureHeight       float64             `json:"enclosureHeight"`
	BaseExtension         float64             `json:"baseExtension"`
	BaseThickness         float64             `json:"baseThickness"`
	CushionExtension      float64             `json:"cushionExtension"`
	CushionThickness      float64             `json:"cushionThickness"`
	PipePositions         []Point2            `json:"pipePositions"`
	PipeInnerDiameters    []float64           `json:"pipeInnerDiameters"`
	PipeWallThicknesses   []float64           `json:"pipeWallThicknesses"`
	PullPipeInnerDiameter float64             `json:"pullPipeInnerDiameter"`
	PullPipeThickness     float64             `json:"pullPipeThickness"`
	Points                []ChannelPoint      `json:"-"`
	PointIndexs           []ChannelPointIndex `json:"pointIndexs"`
}

func NewPipeRow() *PipeRow {
	return &PipeRow{
		EcBase: EcBase{Type: "GIM/EC/PipeRow"},
	}
}

type CableTrench struct {
	EcBase
	Width            float64             `json:"width"`
	Height           float64             `json:"height"`
	CoverWidth       float64             `json:"coverWidth"`
	CoverThickness   float64             `json:"coverThickness"`
	BaseExtension    float64             `json:"baseExtension"`
	BaseThickness    float64             `json:"baseThickness"`
	CushionExtension float64             `json:"cushionExtension"`
	CushionThickness float64             `json:"cushionThickness"`
	WallThickness    float64             `json:"wallThickness"`
	WallThickness2   float64             `json:"wallThickness2"`
	Points           []ChannelPoint      `json:"-"`
	PointIndexs      []ChannelPointIndex `json:"pointIndexs"`
}

func NewCableTrench() *CableTrench {
	return &CableTrench{
		EcBase: EcBase{Type: "GIM/EC/CableTrench"},
	}
}

type CableTunnel struct {
	EcBase
	Style                TunnelStyle         `json:"style"`
	Width                float64             `json:"width"`
	Height               float64             `json:"height"`
	TopThickness         float64             `json:"topThickness"`
	BottomThickness      float64             `json:"bottomThickness"`
	OuterWallThickness   float64             `json:"outerWallThickness"`
	InnerWallThickness   float64             `json:"innerWallThickness"`
	ArcHeight            float64             `json:"arcHeight"`
	BottomPlatformHeight float64             `json:"bottomPlatformHeight"`
	CushionExtension     float64             `json:"cushionExtension"`
	CushionThickness     float64             `json:"cushionThickness"`
	Points               []ChannelPoint      `json:"-"`
	PointIndexs          []ChannelPointIndex `json:"pointIndexs"`
}

func NewCableTunnel() *CableTunnel {
	return &CableTunnel{
		EcBase: EcBase{Type: "GIM/EC/CableTunnel"},
	}
}

type CableTray struct {
	EcBase
	Style               CableTrayStyle      `json:"style"`
	ColumnDiameter      float64             `json:"columnDiameter"`
	ColumnHeight        float64             `json:"columnHeight"`
	Span                float64             `json:"span"`
	Width               float64             `json:"width"`
	Height              float64             `json:"height"`
	TopPlateHeight      float64             `json:"topPlateHeight"`
	ArcHeight           float64             `json:"arcHeight"`
	WallThickness       float64             `json:"wallThickness"`
	PipePositions       []Point2            `json:"pipePositions"`
	PipeInnerDiameters  []float64           `json:"pipeInnerDiameters"`
	PipeWallThicknesses []float64           `json:"pipeWallThicknesses"`
	HasProtectionPlate  bool                `json:"hasProtectionPlate"`
	Points              []ChannelPoint      `json:"-"`
	PointIndexs         []ChannelPointIndex `json:"pointIndexs"`
}

func NewCableTray() *CableTray {
	return &CableTray{
		EcBase: EcBase{Type: "GIM/EC/CableTray"},
	}
}

type CableLBeam struct {
	EcBase
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func NewCableLBeam() *CableLBeam {
	return &CableLBeam{
		EcBase: EcBase{Type: "GIM/EC/CableLBeam"},
	}
}

type Manhole struct {
	EcBase
	Style         ManholeStyle `json:"style"`
	Length        float64      `json:"length"`
	Width         float64      `json:"width"`
	Height        float64      `json:"height"`
	WallThickness float64      `json:"wallThickness"`
}

func NewManhole() *Manhole {
	return &Manhole{
		EcBase: EcBase{Type: "GIM/EC/Manhole"},
	}
}

type ManholeCover struct {
	EcBase
	Style     ManholeStyle `json:"style"`
	Length    float64      `json:"length"`
	Width     float64      `json:"width"`
	Thickness float64      `json:"thickness"`
}

func NewManholeCover() *ManholeCover {
	return &ManholeCover{
		EcBase: EcBase{Type: "GIM/EC/ManholeCover"},
	}
}

type Ladder struct {
	EcBase
	Length    float64 `json:"length"`
	Width     float64 `json:"width"`
	Thickness float64 `json:"thickness"`
}

func NewLadder() *Ladder {
	return &Ladder{
		EcBase: EcBase{Type: "GIM/EC/Ladder"},
	}
}

type Sump struct {
	EcBase
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
	Depth           float64 `json:"depth"`
	BottomThickness float64 `json:"bottomThickness"`
}

func NewSump() *Sump {
	return &Sump{
		EcBase: EcBase{Type: "GIM/EC/Sump"},
	}
}

type Footpath struct {
	EcBase
	Height      float64             `json:"height"`
	Width       float64             `json:"width"`
	Points      []ChannelPoint      `json:"-"`
	PointIndexs []ChannelPointIndex `json:"pointIndexs"`
}

func NewFootpath() *Footpath {
	return &Footpath{
		EcBase: EcBase{Type: "GIM/EC/Footpath"},
	}
}

type ShaftChamber struct {
	EcBase
	SupportWallThickness float64 `json:"supportWallThickness"`
	SupportDiameter      float64 `json:"supportDiameter"`
	SupportHeight        float64 `json:"supportHeight"`
	TopThickness         float64 `json:"topThickness"`
	InnerDiameter        float64 `json:"innerDiameter"`
	WorkingHeight        float64 `json:"workingHeight"`
	OuterWallThickness   float64 `json:"outerWallThickness"`
	InnerWallThickness   float64 `json:"innerWallThickness"`
}

func NewShaftChamber() *ShaftChamber {
	return &ShaftChamber{
		EcBase: EcBase{Type: "GIM/EC/ShaftChamber"},
	}
}

type TunnelCompartmentPartition struct {
	EcBase
	Width     float64 `json:"width"`
	Thickness float64 `json:"thickness"`
}

func NewTunnelCompartmentPartition() *TunnelCompartmentPartition {
	return &TunnelCompartmentPartition{
		EcBase: EcBase{Type: "GIM/EC/TunnelCompartmentPartition"},
	}
}

type VentilationPavilion struct {
	EcBase
	TopLength    float64 `json:"topLength"`
	MiddleLength float64 `json:"middleLength"`
	BottomLength float64 `json:"bottomLength"`
	TopWidth     float64 `json:"topWidth"`
	MiddleWidth  float64 `json:"middleWidth"`
	BottomWidth  float64 `json:"bottomWidth"`
	TopHeight    float64 `json:"topHeight"`
	Height       float64 `json:"height"`
	BaseHeight   float64 `json:"baseHeight"`
}

func NewVentilationPavilion() *VentilationPavilion {
	return &VentilationPavilion{
		EcBase: EcBase{Type: "GIM/EC/VentilationPavilion"},
	}
}

type TunnelPartitionBoard struct {
	EcBase
	Style         PartitionBoardStyle `json:"style"`
	Length        float64             `json:"length"`
	Width         float64             `json:"width"`
	Thickness     float64             `json:"thickness"`
	HoleCount     int                 `json:"holeCount"`
	HolePositions []Point2            `json:"holePositions"`
	HoleStyles    []int               `json:"holeStyles"`
	HoleDiameters []float64           `json:"holeDiameters"`
	HoleWidths    []float64           `json:"holeWidths"`
}

func NewTunnelPartitionBoard() *TunnelPartitionBoard {
	return &TunnelPartitionBoard{
		EcBase: EcBase{Type: "GIM/EC/TunnelPartitionBoard"},
	}
}

type StraightVentilationDuct struct {
	EcBase
	Diameter      float64 `json:"diameter"`
	WallThickness float64 `json:"wallThickness"`
	Height        float64 `json:"height"`
}

func NewStraightVentilationDuct() *StraightVentilationDuct {
	return &StraightVentilationDuct{
		EcBase: EcBase{Type: "GIM/EC/StraightVentilationDuct"},
	}
}

type ObliqueVentilationDuct struct {
	EcBase
	HoodRoomLength        float64 `json:"hoodRoomLength"`
	HoodRoomWidth         float64 `json:"hoodRoomWidth"`
	HoodRoomHeight        float64 `json:"hoodRoomHeight"`
	HoodWallThickness     float64 `json:"hoodWallThickness"`
	DuctCenterHeight      float64 `json:"ductCenterHeight"`
	DuctLeftDistance      float64 `json:"ductLeftDistance"`
	DuctDiameter          float64 `json:"ductDiameter"`
	DuctWallThickness     float64 `json:"ductWallThickness"`
	DuctLength            float64 `json:"ductLength"`
	DuctHeightDifference  float64 `json:"ductHeightDifference"`
	BaseLength            float64 `json:"baseLength"`
	BaseWidth             float64 `json:"baseWidth"`
	BaseHeight            float64 `json:"baseHeight"`
	BaseRoomLength        float64 `json:"baseRoomLength"`
	BaseRoomWallThickness float64 `json:"baseRoomWallThickness"`
	BaseRoomWidth         float64 `json:"baseRoomWidth"`
	BaseRoomHeight        float64 `json:"baseRoomHeight"`
}

func NewObliqueVentilationDuct() *ObliqueVentilationDuct {
	return &ObliqueVentilationDuct{
		EcBase: EcBase{Type: "GIM/EC/ObliqueVentilationDuct"},
	}
}

type DrainageWell struct {
	EcBase
	Length           float64 `json:"length"`
	Width            float64 `json:"width"`
	Height           float64 `json:"height"`
	NeckDiameter     float64 `json:"neckDiameter"`
	NeckHeight       float64 `json:"neckHeight"`
	CushionExtension float64 `json:"cushionExtension"`
	BottomThickness  float64 `json:"bottomThickness"`
	WallThickness    float64 `json:"wallThickness"`
}

func NewDrainageWell() *DrainageWell {
	return &DrainageWell{
		EcBase: EcBase{Type: "GIM/EC/DrainageWell"},
	}
}

type PipeSupport struct {
	EcBase
	Style     PipeSupportStyle `json:"style"`
	Positions []Point2         `json:"positions"`
	Radii     []float64        `json:"radii"`
	Length    float64          `json:"length"`
	Width     float64          `json:"width"`
	Height    float64          `json:"height"`
}

func NewPipeSupport() *PipeSupport {
	return &PipeSupport{
		EcBase: EcBase{Type: "GIM/EC/PipeSupport"},
	}
}

type CoverPlate struct {
	EcBase
	Style       CoverPlateStyle `json:"style"`
	Length      float64         `json:"length"`
	Width       float64         `json:"width"`
	SmallRadius float64         `json:"smallRadius"`
	LargeRadius float64         `json:"largeRadius"`
	Thickness   float64         `json:"thickness"`
}

func NewCoverPlate() *CoverPlate {
	return &CoverPlate{
		EcBase: EcBase{Type: "GIM/EC/CoverPlate"},
	}
}

type CableRay struct {
	EcBase
	OuterLength    float64 `json:"outerLength"`
	OuterHeight    float64 `json:"outerHeight"`
	InnerLength    float64 `json:"innerLength"`
	InnerHeight    float64 `json:"innerHeight"`
	CoverThickness float64 `json:"coverThickness"`
}

func NewCableRay() *CableRay {
	return &CableRay{
		EcBase: EcBase{Type: "GIM/EC/CableRay"},
	}
}

type Shape interface {
	GetType() string
}

func Unmarshal(ty string, bt []byte) (Shape, error) {
	switch ty {
	case "GIM/EC/CableWire":
		t := &CableWire{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/CableJoint":
		t := &CableJoint{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/OpticalFiberBox":
		t := &OpticalFiberBox{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/CableTerminal":
		t := &CableTerminal{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/CableAccessory":
		t := &CableAccessory{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/CableBracket":
		t := &CableBracket{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/CableClamp":
		t := &CableClamp{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/CablePole":
		t := &CablePole{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/GroundFlatIron":
		t := &GroundFlatIron{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/EmbeddedPart":
		t := &EmbeddedPart{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/UShapedRing":
		t := &UShapedRing{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/LiftingEye":
		t := &LiftingEye{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/CornerWell":
		t := &CornerWell{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/TunnelWell":
		t := &TunnelWell{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/ThreeWayWell":
		t := &ThreeWayWell{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/FourWayWell":
		t := &FourWayWell{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/PipeRow":
		t := &PipeRow{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/CableTrench":
		t := &CableTrench{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/CableTunnel":
		t := &CableTunnel{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/CableTray":
		t := &CableTray{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/CableLBeam":
		t := &CableLBeam{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/Manhole":
		t := &Manhole{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/ManholeCover":
		t := &ManholeCover{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/Ladder":
		t := &Ladder{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/Sump":
		t := &Sump{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/Footpath":
		t := &Footpath{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/ShaftChamber":
		t := &ShaftChamber{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/TunnelCompartmentPartition":
		t := &TunnelCompartmentPartition{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/VentilationPavilion":
		t := &VentilationPavilion{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/TunnelPartitionBoard":
		t := &TunnelPartitionBoard{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/StraightVentilationDuct":
		t := &StraightVentilationDuct{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/ObliqueVentilationDuct":
		t := &ObliqueVentilationDuct{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/DrainageWell":
		t := &DrainageWell{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/PipeSupport":
		t := &PipeSupport{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/CoverPlate":
		t := &CoverPlate{}
		e := json.Unmarshal(bt, t)
		return t, e
	case "GIM/EC/CableRay":
		t := &CableRay{}
		e := json.Unmarshal(bt, t)
		return t, e

	default:
		return nil, fmt.Errorf("unknown type: %s", ty)
	}
}
