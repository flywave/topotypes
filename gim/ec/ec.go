package ec

import (
	"encoding/json"
	"fmt"
)

const Major = "EC"

type EcBase struct {
	Version int    `json:"version"`
	Type    string `json:"type"`
}

func (ec *EcBase) GetType() string {
	return ec.Type
}

// CableWireObject represents a cable wire
type CableWire struct {
	EcBase
	Points          [][3]float64 `json:"points"`          // Cable path points
	OutsideDiameter float64      `json:"outsideDiameter"` // Cable outer diameter
}

// NewCableWire creates a new CableWireObject
func NewCableWire() *CableWire {
	return &CableWire{
		EcBase: EcBase{Type: "GIM/EC/CableWire"},
	}
}

// CableJointObject represents a cable joint
type CableJoint struct {
	EcBase
	Length         float64 `json:"length"`         // Joint length
	OuterDiameter  float64 `json:"outerDiameter"`  // Outer diameter
	TerminalLength float64 `json:"terminalLength"` // Terminal length
	InnerDiameter  float64 `json:"innerDiameter"`  // Inner diameter
}

// NewCableJoint creates a new CableJointObject
func NewCableJoint() *CableJoint {
	return &CableJoint{
		EcBase: EcBase{Type: "GIM/EC/CableJoint"},
	}
}

// OpticalFiberBoxObject represents an optical fiber box
type OpticalFiberBox struct {
	EcBase
	Length float64 `json:"length"` // Length
	Height float64 `json:"height"` // Height
	Width  float64 `json:"width"`  // Width
}

// NewOpticalFiberBox creates a new OpticalFiberBoxObject
func NewOpticalFiberBox() *OpticalFiberBox {
	return &OpticalFiberBox{
		EcBase: EcBase{Type: "GIM/EC/OpticalFiberBox"},
	}
}

// CableTerminalObject represents a cable terminal
type CableTerminal struct {
	EcBase
	Sort                     int     `json:"sort"`                     // Type (1-3)
	Height                   float64 `json:"height"`                   // Total height
	TopDiameter              float64 `json:"topDiameter"`              // Top diameter
	BottomDiameter           float64 `json:"bottomDiameter"`           // Bottom diameter
	TailDiameter             float64 `json:"tailDiameter"`             // Tail diameter
	TailHeight               float64 `json:"tailHeight"`               // Tail height
	SkirtCount               int     `json:"skirtCount"`               // Skirt count
	UpperSkirtTopDiameter    float64 `json:"upperSkirtTopDiameter"`    // Upper skirt top diameter
	UpperSkirtBottomDiameter float64 `json:"upperSkirtBottomDiameter"` // Upper skirt bottom diameter
	LowerSkirtTopDiameter    float64 `json:"lowerSkirtTopDiameter"`    // Lower skirt top diameter
	LowerSkirtBottomDiameter float64 `json:"lowerSkirtBottomDiameter"` // Lower skirt bottom diameter
	SkirtSectionHeight       float64 `json:"skirtSectionHeight"`       // Skirt section height
	UpperTerminalLength      float64 `json:"upperTerminalLength"`      // Upper terminal length
	UpperTerminalDiameter    float64 `json:"upperTerminalDiameter"`    // Upper terminal diameter
	LowerTerminalLength      float64 `json:"lowerTerminalLength"`      // Lower terminal length
	LowerTerminalDiameter    float64 `json:"lowerTerminalDiameter"`    // Lower terminal diameter
	Hole1Diameter            float64 `json:"hole1Diameter"`            // Hole 1 diameter
	Hole2Diameter            float64 `json:"hole2Diameter"`            // Hole 2 diameter
	Hole1Distance            float64 `json:"hole1Distance"`            // Hole 1 distance
	HoleSpacing              float64 `json:"holeSpacing"`              // Hole spacing
	FlangeHoleDiameter       float64 `json:"flangeHoleDiameter"`       // Flange hole diameter
	FlangeHoleSpacing        float64 `json:"flangeHoleSpacing"`        // Flange hole spacing
	FlangeWidth              float64 `json:"flangeWidth"`              // Flange width
	FlangeCenterHoleRadius   float64 `json:"flangeCenterHoleRadius"`   // Flange center hole radius
	FlangeChamferRadius      float64 `json:"flangeChamferRadius"`      // Flange chamfer radius
	FlangeOpeningWidth       float64 `json:"flangeOpeningWidth"`       // Flange opening width
	FlangeBoltHeight         float64 `json:"flangeBoltHeight"`         // Flange bolt height
}

// NewCableTerminal creates a new CableTerminalObject
func NewCableTerminal() *CableTerminal {
	return &CableTerminal{
		EcBase: EcBase{Type: "GIM/EC/CableTerminal"},
	}
}

// CableAccessoryObject represents a cable accessory
type CableAccessory struct {
	EcBase
	Length            float64 `json:"length"`            // Length
	Width             float64 `json:"width"`             // Width
	Height            float64 `json:"height"`            // Height
	PortCount         int     `json:"portCount"`         // Port count (3 or 6)
	PortDiameter      float64 `json:"portDiameter"`      // Port diameter
	PortSpacing       float64 `json:"portSpacing"`       // Port spacing
	BackPanelDistance float64 `json:"backPanelDistance"` // Back panel distance
	SidePanelDistance float64 `json:"sidePanelDistance"` // Side panel distance
}

// NewCableAccessory creates a new CableAccessoryObject
func NewCableAccessory() *CableAccessory {
	return &CableAccessory{
		EcBase: EcBase{Type: "GIM/EC/CableAccessory"},
	}
}

// CableBracketObject represents a cable bracket
type CableBracket struct {
	EcBase
	Length            float64      `json:"length"`            // Length
	RootHeight        float64      `json:"rootHeight"`        // Root height
	RootWidth         float64      `json:"rootWidth"`         // Root width
	Width             float64      `json:"width"`             // Width
	TopThickness      float64      `json:"topThickness"`      // Top thickness
	RootThickness     float64      `json:"rootThickness"`     // Root thickness
	ColumnMountPoints [][3]float64 `json:"columnMountPoints"` // Column mount points
	ClampMountPoints  [][3]float64 `json:"clampMountPoints"`  // Clamp mount points
}

// NewCableBracket creates a new CableBracketObject
func NewCableBracket() *CableBracket {
	return &CableBracket{
		EcBase: EcBase{Type: "GIM/EC/CableBracket"},
	}
}

// CableClampObject represents a cable clamp
type CableClamp struct {
	EcBase
	ClampType string  `json:"clampType"` // 'SINGLE' | 'LINEAR' | 'CONTACT_TRIPLE' | 'SEPARATE_TRIPLE'
	Diameter  float64 `json:"diameter"`  // Diameter
	Thickness float64 `json:"thickness"` // Thickness
	Width     float64 `json:"width"`     // Width
}

// NewCableClamp creates a new CableClampObject
func NewCableClamp() *CableClamp {
	return &CableClamp{
		EcBase: EcBase{Type: "GIM/EC/CableClamp"},
	}
}

// CablePoleObject represents a cable pole
type CablePole struct {
	EcBase
	Specification  string       `json:"specification"`  // Specification model
	Length         float64      `json:"length"`         // Length
	Radius         float64      `json:"radius"`         // Radius
	ArcAngle       float64      `json:"arcAngle"`       // Arc angle
	Width          float64      `json:"width"`          // Width
	FixedLegLength float64      `json:"fixedLegLength"` // Fixed leg length
	FixedLegWidth  float64      `json:"fixedLegWidth"`  // Fixed leg width
	Thickness      float64      `json:"thickness"`      // Thickness
	MountPoints    [][3]float64 `json:"mountPoints"`    // Mount points
}

// NewCablePole creates a new CablePoleObject
func NewCablePole() *CablePole {
	return &CablePole{
		EcBase: EcBase{Type: "GIM/EC/CablePole"},
	}
}

// GroundFlatIronObject represents a ground flat iron
type GroundFlatIron struct {
	EcBase
	Length    float64 `json:"length"`    // Length
	Height    float64 `json:"height"`    // Height
	Thickness float64 `json:"thickness"` // Thickness
}

// NewGroundFlatIron creates a new GroundFlatIronObject
func NewGroundFlatIron() *GroundFlatIron {
	return &GroundFlatIron{
		EcBase: EcBase{Type: "GIM/EC/GroundFlatIron"},
	}
}

// EmbeddedPartObject represents an embedded part
type EmbeddedPart struct {
	EcBase
	Length         float64 `json:"length"`         // Length
	Radius         float64 `json:"radius"`         // Radius
	Height         float64 `json:"height"`         // Height
	MaterialRadius float64 `json:"materialRadius"` // Material radius
	LowerLength    float64 `json:"lowerLength"`    // Lower length
}

// NewEmbeddedPart creates a new EmbeddedPartObject
func NewEmbeddedPart() *EmbeddedPart {
	return &EmbeddedPart{
		EcBase: EcBase{Type: "GIM/EC/EmbeddedPart"},
	}
}

// UShapedRingObject represents a U-shaped ring
type UShapedRing struct {
	EcBase
	Thickness float64 `json:"thickness"` // Thickness
	Height    float64 `json:"height"`    // Height
	Radius    float64 `json:"radius"`    // Radius
	Length    float64 `json:"length"`    // Length
}

// NewUShapedRing creates a new UShapedRingObject
func NewUShapedRing() *UShapedRing {
	return &UShapedRing{
		EcBase: EcBase{Type: "GIM/EC/UShapedRing"},
	}
}

// LiftingEyeObject represents a lifting eye
type LiftingEye struct {
	EcBase
	Height       float64 `json:"height"`       // Height
	RingRadius   float64 `json:"ringRadius"`   // Ring radius
	PipeDiameter float64 `json:"pipeDiameter"` // Pipe diameter
}

// NewLiftingEye creates a new LiftingEyeObject
func NewLiftingEye() *LiftingEye {
	return &LiftingEye{
		EcBase: EcBase{Type: "GIM/EC/LiftingEye"},
	}
}

// CornerWellObject represents a corner well
type CornerWell struct {
	EcBase
	LeftLength       float64 `json:"leftLength"`       // Left length
	RightLength      float64 `json:"rightLength"`      // Right length
	Width            float64 `json:"width"`            // Width
	Height           float64 `json:"height"`           // Height
	TopThickness     float64 `json:"topThickness"`     // Top thickness
	BottomThickness  float64 `json:"bottomThickness"`  // Bottom thickness
	WallThickness    float64 `json:"wallThickness"`    // Wall thickness
	Angle            float64 `json:"angle"`            // Angle
	CornerRadius     float64 `json:"cornerRadius"`     // Corner radius
	CushionExtension float64 `json:"cushionExtension"` // Cushion extension
	CushionThickness float64 `json:"cushionThickness"` // Cushion thickness
}

// NewCornerWell creates a new CornerWellObject
func NewCornerWell() *CornerWell {
	return &CornerWell{
		EcBase: EcBase{Type: "GIM/EC/CornerWell"},
	}
}

// TunnelWellObject represents a tunnel well
type TunnelWell struct {
	EcBase
	WellType           string  `json:"wellType"`           // 'STRAIGHT' | 'STRAIGHT_TUNNEL'
	Length             float64 `json:"length"`             // Length
	Width              float64 `json:"width"`              // Width
	Height             float64 `json:"height"`             // Height
	TopThickness       float64 `json:"topThickness"`       // Top thickness
	BottomThickness    float64 `json:"bottomThickness"`    // Bottom thickness
	OuterWallThickness float64 `json:"outerWallThickness"` // Outer wall thickness
	CushionExtension   float64 `json:"cushionExtension"`   // Cushion extension
	CushionThickness   float64 `json:"cushionThickness"`   // Cushion thickness
	LeftSectionType    string  `json:"leftSectionType"`    // 'RECTANGULAR' | 'HORSESHOE' | 'CIRCULAR'
	LeftLength         float64 `json:"leftLength"`         // Left length
	LeftWidth          float64 `json:"leftWidth"`          // Left width
	LeftHeight         float64 `json:"leftHeight"`         // Left height
	LeftArcHeight      float64 `json:"leftArcHeight"`      // Left arc height
	RightSectionType   string  `json:"rightSectionType"`   // 'RECTANGULAR' | 'HORSESHOE' | 'CIRCULAR'
	RightLength        float64 `json:"rightLength"`        // Right length
	RightWidth         float64 `json:"rightWidth"`         // Right width
	RightHeight        float64 `json:"rightHeight"`        // Right height
	RightArcHeight     float64 `json:"rightArcHeight"`     // Right arc height
	Radius             float64 `json:"radius"`             // Radius
	InnerWallThickness float64 `json:"innerWallThickness"` // Inner wall thickness
}

// NewTunnelWell creates a new TunnelWellObject
func NewTunnelWell() *TunnelWell {
	return &TunnelWell{
		EcBase: EcBase{Type: "GIM/EC/TunnelWell"},
	}
}

// ThreeWayWellObject represents a three-way well
type ThreeWayWell struct {
	EcBase
	WellType               string  `json:"wellType"`               // 'UNDERGROUND_TUNNEL' | 'OPEN_CUT_TUNNEL' | 'WORKING_WELL'
	CornerType             string  `json:"cornerType"`             // 'ROUNDED' | 'ANGLED'
	ShaftType              string  `json:"shaftType"`              // 'CIRCULAR' | 'RECTANGULAR'
	Length                 float64 `json:"length"`                 // Length
	Width                  float64 `json:"width"`                  // Width
	Height                 float64 `json:"height"`                 // Height
	ShaftRadius            float64 `json:"shaftRadius"`            // Shaft radius
	CornerRadius           float64 `json:"cornerRadius"`           // Corner radius
	CornerLength           float64 `json:"cornerLength"`           // Corner length
	CornerWidth            float64 `json:"cornerWidth"`            // Corner width
	BranchLength           float64 `json:"branchLength"`           // Branch length
	BranchLeftLength       float64 `json:"branchLeftLength"`       // Branch left length
	BranchWidth            float64 `json:"branchWidth"`            // Branch width
	TopThickness           float64 `json:"topThickness"`           // Top thickness
	BottomThickness        float64 `json:"bottomThickness"`        // Bottom thickness
	LeftSectionStyle       string  `json:"leftSectionStyle"`       // 'RECTANGULAR' | 'HORSESHOE' | 'CIRCULAR'
	LeftSectionLength      float64 `json:"leftSectionLength"`      // Left section length
	LeftSectionWidth       float64 `json:"leftSectionWidth"`       // Left section width
	LeftSectionHeight      float64 `json:"leftSectionHeight"`      // Left section height
	LeftSectionArcHeight   float64 `json:"leftSectionArcHeight"`   // Left section arc height
	RightSectionStyle      string  `json:"rightSectionStyle"`      // 'RECTANGULAR' | 'HORSESHOE' | 'CIRCULAR'
	RightSectionLength     float64 `json:"rightSectionLength"`     // Right section length
	RightSectionWidth      float64 `json:"rightSectionWidth"`      // Right section width
	RightSectionHeight     float64 `json:"rightSectionHeight"`     // Right section height
	RightSectionArcHeight  float64 `json:"rightSectionArcHeight"`  // Right section arc height
	BranchSectionStyle     string  `json:"branchSectionStyle"`     // 'RECTANGULAR' | 'HORSESHOE' | 'CIRCULAR'
	BranchSectionLength    float64 `json:"branchSectionLength"`    // Branch section length
	BranchSectionWidth     float64 `json:"branchSectionWidth"`     // Branch section width
	BranchSectionHeight    float64 `json:"branchSectionHeight"`    // Branch section height
	BranchSectionArcHeight float64 `json:"branchSectionArcHeight"` // Branch section arc height
	OuterWallThickness     float64 `json:"outerWallThickness"`     // Outer wall thickness
	InnerWallThickness     float64 `json:"innerWallThickness"`     // Inner wall thickness
	IsDoubleShaft          bool    `json:"isDoubleShaft"`          // Is double shaft
	DoubleShaftSpacing     float64 `json:"doubleShaftSpacing"`     // Double shaft spacing
	OuterWallExtension     float64 `json:"outerWallExtension"`     // Outer wall extension
	InnerWallExtension     float64 `json:"innerWallExtension"`     // Inner wall extension
	CushionExtension       float64 `json:"cushionExtension"`       // Cushion extension
	CushionThickness       float64 `json:"cushionThickness"`       // Cushion thickness
	InnerBottomThickness   float64 `json:"innerBottomThickness"`   // Inner bottom thickness
	OuterBottomThickness   float64 `json:"outerBottomThickness"`   // Outer bottom thickness
	Angle                  float64 `json:"angle"`                  // Angle
}

// NewThreeWayWell creates a new ThreeWayWellObject
func NewThreeWayWell() *ThreeWayWell {
	return &ThreeWayWell{
		EcBase: EcBase{Type: "GIM/EC/ThreeWayWell"},
	}
}

// FourWayWellObject represents a four-way well

// FourWayWell represents a four-way well structure
type FourWayWell struct {
	EcBase
	WellType           string  `json:"wellType"`           // 类型(1-地下隧道,2-其他)
	Length             float64 `json:"length"`             // 长度
	Width              float64 `json:"width"`              // 宽度
	Height             float64 `json:"height"`             // 高度
	CornerStyle        string  `json:"cornerStyle"`        // 转角类型(1-圆形,2-方形)
	CornerRadius       float64 `json:"cornerRadius"`       // 转角半径
	BranchLength       float64 `json:"branchLength"`       // 分支长度
	BranchWidth        float64 `json:"branchWidth"`        // 分支宽度
	TopThickness       float64 `json:"topThickness"`       // 顶部厚度
	BottomThickness    float64 `json:"bottomThickness"`    // 底部厚度
	OuterWallThickness float64 `json:"outerWallThickness"` // 外壁厚度
	InnerWallThickness float64 `json:"innerWallThickness"` // 内壁厚度
	CushionExtension   float64 `json:"cushionExtension"`   // 垫层延伸
	CushionThickness   float64 `json:"cushionThickness"`   // 垫层厚度

	// Left section properties
	LeftSectionStyle     string  `json:"leftSectionStyle"`     // 左侧截面类型(1-矩形,2-马蹄形,3-圆形)
	LeftSectionLength    float64 `json:"leftSectionLength"`    // 左侧截面长度
	LeftSectionWidth     float64 `json:"leftSectionWidth"`     // 左侧截面宽度
	LeftSectionHeight    float64 `json:"leftSectionHeight"`    // 左侧截面高度
	LeftSectionArcHeight float64 `json:"leftSectionArcHeight"` // 左侧截面弧高

	// Right section properties
	RightSectionStyle     string  `json:"rightSectionStyle"`     // 右侧截面类型
	RightSectionLength    float64 `json:"rightSectionLength"`    // 右侧截面长度
	RightSectionWidth     float64 `json:"rightSectionWidth"`     // 右侧截面宽度
	RightSectionHeight    float64 `json:"rightSectionHeight"`    // 右侧截面高度
	RightSectionArcHeight float64 `json:"rightSectionArcHeight"` // 右侧截面弧高

	// Branch 1 properties
	BranchSection1Style     string  `json:"branchSection1Style"`     // 分支1截面类型
	BranchSection1Length    float64 `json:"branchSection1Length"`    // 分支1截面长度
	BranchSection1Width     float64 `json:"branchSection1Width"`     // 分支1截面宽度
	BranchSection1Height    float64 `json:"branchSection1Height"`    // 分支1截面高度
	BranchSection1ArcHeight float64 `json:"branchSection1ArcHeight"` // 分支1截面弧高

	// Branch 2 properties
	BranchSection2Style     string  `json:"branchSection2Style"`     // 分支2截面类型
	BranchSection2Length    float64 `json:"branchSection2Length"`    // 分支2截面长度
	BranchSection2Width     float64 `json:"branchSection2Width"`     // 分支2截面宽度
	BranchSection2Height    float64 `json:"branchSection2Height"`    // 分支2截面高度
	BranchSection2ArcHeight float64 `json:"branchSection2ArcHeight"` // 分支2截面弧高

	ShaftRadius  float64 `json:"shaftRadius"`  // 井筒半径
	CornerLength float64 `json:"cornerLength"` // 转角长度
	CornerWidth  float64 `json:"cornerWidth"`  // 转角宽度
}

// NewFourWayWell creates a new FourWayWell instance
func NewFourWayWell() *FourWayWell {
	return &FourWayWell{
		EcBase: EcBase{Type: "GIM/EC/FourWayWell"},
	}
}

// Constants for well types
const (
	WellTypeUndergroundTunnel = "UNDERGROUND_TUNNEL"
	WellTypeOpenCutTunnel     = "OPEN_CUT_TUNNEL"
	WellTypeWorkingWell       = "WORKING_WELL"
)

// Constants for corner styles
const (
	CornerStyleRounded = "ROUNDED"
	CornerStyleAngled  = "ANGLED"
)

// Constants for section styles
const (
	SectionStyleRectangular = "RECTANGULAR"
	SectionStyleHorseshoe   = "HORSESHOE"
	SectionStyleCircular    = "CIRCULAR"
)

type Point struct {
	Position [3]float64 `json:"position"`
	Type     int        `json:"type"`
}

// PipeRowObject represents a pipe row
type PipeRow struct {
	EcBase
	PipeType              int          `json:"pipeType"`              // Pipe type
	HasEnclosure          bool         `json:"hasEnclosure"`          // Has enclosure
	EnclosureWidth        float64      `json:"enclosureWidth"`        // Enclosure width
	EnclosureHeight       float64      `json:"enclosureHeight"`       // Enclosure height
	BaseExtension         float64      `json:"baseExtension"`         // EcBase extension
	BaseThickness         float64      `json:"baseThickness"`         // EcBase thickness
	CushionExtension      float64      `json:"cushionExtension"`      // Cushion extension
	CushionThickness      float64      `json:"cushionThickness"`      // Cushion thickness
	PipePositions         [][2]float64 `json:"pipePositions"`         // Pipe positions
	PipeInnerDiameters    []float64    `json:"pipeInnerDiameters"`    // Pipe inner diameters
	PipeWallThicknesses   []float64    `json:"pipeWallThicknesses"`   // Pipe wall thicknesses
	PullPipeInnerDiameter float64      `json:"pullPipeInnerDiameter"` // Pull pipe inner diameter
	PullPipeThickness     float64      `json:"pullPipeThickness"`     // Pull pipe thickness
	Points                []*Point     `json:"points"`                // Path points
}

// NewPipeRow creates a new PipeRowObject
func NewPipeRow() *PipeRow {
	return &PipeRow{
		EcBase: EcBase{Type: "GIM/EC/PipeRow"},
	}
}

// CableTrenchObject represents a cable trench
type CableTrench struct {
	EcBase
	Width            float64  `json:"width"`            // Width
	Height           float64  `json:"height"`           // Height
	CoverWidth       float64  `json:"coverWidth"`       // Cover width
	CoverThickness   float64  `json:"coverThickness"`   // Cover thickness
	BaseExtension    float64  `json:"baseExtension"`    // EcBase extension
	BaseThickness    float64  `json:"baseThickness"`    // EcBase thickness
	CushionExtension float64  `json:"cushionExtension"` // Cushion extension
	CushionThickness float64  `json:"cushionThickness"` // Cushion thickness
	WallThickness    float64  `json:"wallThickness"`    // Wall thickness
	WallThickness2   float64  `json:"wallThickness2"`   // Wall thickness 2
	Points           []*Point `json:"points"`           // Path points
}

// NewCableTrench creates a new CableTrenchObject
func NewCableTrench() *CableTrench {
	return &CableTrench{
		EcBase: EcBase{Type: "GIM/EC/CableTrench"},
	}
}

// CableTunnelObject represents a cable tunnel
type CableTunnel struct {
	EcBase
	Style                string   `json:"style"`                // 'RECTANGULAR' | 'HORSESHOE' | 'CIRCULAR'
	Width                float64  `json:"width"`                // Width
	Height               float64  `json:"height"`               // Height
	TopThickness         float64  `json:"topThickness"`         // Top thickness
	BottomThickness      float64  `json:"bottomThickness"`      // Bottom thickness
	OuterWallThickness   float64  `json:"outerWallThickness"`   // Outer wall thickness
	InnerWallThickness   float64  `json:"innerWallThickness"`   // Inner wall thickness
	ArcHeight            float64  `json:"arcHeight"`            // Arc height
	BottomPlatformHeight float64  `json:"bottomPlatformHeight"` // Bottom platform height
	CushionExtension     float64  `json:"cushionExtension"`     // Cushion extension
	CushionThickness     float64  `json:"cushionThickness"`     // Cushion thickness
	Points               []*Point `json:"points"`               // Path points
}

// NewCableTunnel creates a new CableTunnelObject
func NewCableTunnel() *CableTunnel {
	return &CableTunnel{
		EcBase: EcBase{Type: "GIM/EC/CableTunnel"},
	}
}

// CableTrayObject represents a cable tray
type CableTray struct {
	EcBase
	Style               string       `json:"style"`               // 'ARCH' | 'BEAM'
	ColumnDiameter      float64      `json:"columnDiameter"`      // Column diameter
	ColumnHeight        float64      `json:"columnHeight"`        // Column height
	Span                float64      `json:"span"`                // Span
	Width               float64      `json:"width"`               // Width
	Height              float64      `json:"height"`              // Height
	TopPlateHeight      float64      `json:"topPlateHeight"`      // Top plate height
	ArcHeight           float64      `json:"arcHeight"`           // Arc height
	WallThickness       float64      `json:"wallThickness"`       // Wall thickness
	PipeCount           int          `json:"pipeCount"`           // Pipe count
	PipePositions       [][2]float64 `json:"pipePositions"`       // Pipe positions
	PipeInnerDiameters  []float64    `json:"pipeInnerDiameters"`  // Pipe inner diameters
	PipeWallThicknesses []float64    `json:"pipeWallThicknesses"` // Pipe wall thicknesses
	HasProtectionPlate  bool         `json:"hasProtectionPlate"`  // Has protection plate
	Points              []*Point     `json:"points"`              // Path points
}

// NewCableTray creates a new CableTrayObject
func NewCableTray() *CableTray {
	return &CableTray{
		EcBase: EcBase{Type: "GIM/EC/CableTray"},
	}
}

// CableLBeamObject represents a cable L-beam
type CableLBeam struct {
	EcBase
	Length float64 `json:"length"` // Length
	Width  float64 `json:"width"`  // Width
	Height float64 `json:"height"` // Height
}

// NewCableLBeam creates a new CableLBeamObject
func NewCableLBeam() *CableLBeam {
	return &CableLBeam{
		EcBase: EcBase{Type: "GIM/EC/CableLBeam"},
	}
}

// ManholeObject represents a manhole
type Manhole struct {
	EcBase
	Style         string  `json:"style"`         // 'CIRCULAR' | 'RECTANGULAR'
	Length        float64 `json:"length"`        // Length
	Width         float64 `json:"width"`         // Width
	Height        float64 `json:"height"`        // Height
	WallThickness float64 `json:"wallThickness"` // Wall thickness
}

// NewManhole creates a new ManholeObject
func NewManhole() *Manhole {
	return &Manhole{
		EcBase: EcBase{Type: "GIM/EC/Manhole"},
	}
}

// ManholeCoverObject represents a manhole cover
type ManholeCover struct {
	EcBase
	Style     string  `json:"style"`     // 'CIRCULAR' | 'RECTANGULAR'
	Length    float64 `json:"length"`    // Length
	Width     float64 `json:"width"`     // Width
	Thickness float64 `json:"thickness"` // Thickness
}

// NewManholeCover creates a new ManholeCoverObject
func NewManholeCover() *ManholeCover {
	return &ManholeCover{
		EcBase: EcBase{Type: "GIM/EC/ManholeCover"},
	}
}

// LadderObject represents a ladder
type Ladder struct {
	EcBase
	Length    float64 `json:"length"`    // Height
	Width     float64 `json:"width"`     // Width
	Thickness float64 `json:"thickness"` // Thickness
}

// NewLadder creates a new LadderObject
func NewLadder() *Ladder {
	return &Ladder{
		EcBase: EcBase{Type: "GIM/EC/Ladder"},
	}
}

// SumpObject represents a sump
type Sump struct {
	EcBase
	Length          float64 `json:"length"`          // Length
	Width           float64 `json:"width"`           // Width
	Depth           float64 `json:"depth"`           // Depth
	BottomThickness float64 `json:"bottomThickness"` // Bottom thickness
}

// NewSump creates a new SumpObject
func NewSump() *Sump {
	return &Sump{
		EcBase: EcBase{Type: "GIM/EC/Sump"},
	}
}

// FootpathObject represents a footpath
type Footpath struct {
	EcBase
	Height float64  `json:"height"` // Height
	Width  float64  `json:"width"`  // Width
	Points []*Point `json:"points"` // Path points
}

// NewFootpath creates a new FootpathObject
func NewFootpath() *Footpath {
	return &Footpath{
		EcBase: EcBase{Type: "GIM/EC/Footpath"},
	}
}

// ShaftChamberObject represents a shaft chamber
type ShaftChamber struct {
	EcBase
	SupportWallThickness float64 `json:"supportWallThickness"` // Support wall thickness
	SupportDiameter      float64 `json:"supportDiameter"`      // Support diameter
	SupportHeight        float64 `json:"supportHeight"`        // Support height
	TopThickness         float64 `json:"topThickness"`         // Top thickness
	InnerDiameter        float64 `json:"innerDiameter"`        // Inner diameter
	WorkingHeight        float64 `json:"workingHeight"`        // Working height
	OuterWallThickness   float64 `json:"outerWallThickness"`   // Outer wall thickness
	InnerWallThickness   float64 `json:"innerWallThickness"`   // Inner wall thickness
}

// NewShaftChamber creates a new ShaftChamberObject
func NewShaftChamber() *ShaftChamber {
	return &ShaftChamber{
		EcBase: EcBase{Type: "GIM/EC/ShaftChamber"},
	}
}

// TunnelCompartmentPartitionObject represents a tunnel compartment partition
type TunnelCompartmentPartition struct {
	EcBase
	Width     float64 `json:"width"`     // Width
	Thickness float64 `json:"thickness"` // Thickness
}

// NewTunnelCompartmentPartition creates a new TunnelCompartmentPartitionObject
func NewTunnelCompartmentPartition() *TunnelCompartmentPartition {
	return &TunnelCompartmentPartition{
		EcBase: EcBase{Type: "GIM/EC/TunnelCompartmentPartition"},
	}
}

// VentilationPavilionObject represents a ventilation pavilion
type VentilationPavilion struct {
	EcBase
	TopLength    float64 `json:"topLength"`    // Top length
	MiddleLength float64 `json:"middleLength"` // Middle length
	BottomLength float64 `json:"bottomLength"` // Bottom length
	TopWidth     float64 `json:"topWidth"`     // Top width
	MiddleWidth  float64 `json:"middleWidth"`  // Middle width
	BottomWidth  float64 `json:"bottomWidth"`  // Bottom width
	TopHeight    float64 `json:"topHeight"`    // Top height
	Height       float64 `json:"height"`       // Total height
	BaseHeight   float64 `json:"baseHeight"`   // EcBase height
}

// NewVentilationPavilion creates a new VentilationPavilionObject
func NewVentilationPavilion() *VentilationPavilion {
	return &VentilationPavilion{
		EcBase: EcBase{Type: "GIM/EC/VentilationPavilion"},
	}
}

// TunnelPartitionBoardObject represents a tunnel partition board
type TunnelPartitionBoard struct {
	EcBase
	Style         int          `json:"style"`         // Type
	Length        float64      `json:"length"`        // Length
	Width         float64      `json:"width"`         // Width
	Thickness     float64      `json:"thickness"`     // Thickness
	HoleCount     int          `json:"holeCount"`     // Hole count
	HolePositions [][2]float64 `json:"holePositions"` // Hole positions
	HoleStyles    []int        `json:"holeStyles"`    // Hole types
	HoleDiameters []float64    `json:"holeDiameters"` // Hole diameters
	HoleWidths    []float64    `json:"holeWidths"`    // Hole widths
}

// NewTunnelPartitionBoard creates a new TunnelPartitionBoardObject
func NewTunnelPartitionBoard() *TunnelPartitionBoard {
	return &TunnelPartitionBoard{
		EcBase: EcBase{Type: "GIM/EC/TunnelPartitionBoard"},
	}
}

// StraightVentilationDuctObject represents a straight ventilation duct
type StraightVentilationDuct struct {
	EcBase
	Diameter      float64 `json:"diameter"`      // Diameter
	WallThickness float64 `json:"wallThickness"` // Wall thickness
	Height        float64 `json:"height"`        // Height
}

// NewStraightVentilationDuct creates a new StraightVentilationDuctObject
func NewStraightVentilationDuct() *StraightVentilationDuct {
	return &StraightVentilationDuct{
		EcBase: EcBase{Type: "GIM/EC/StraightVentilationDuct"},
	}
}

// ObliqueVentilationDuctObject represents an oblique ventilation duct
type ObliqueVentilationDuct struct {
	EcBase
	HoodRoomLength        float64 `json:"hoodRoomLength"`        // Hood room length
	HoodRoomWidth         float64 `json:"hoodRoomWidth"`         // Hood room width
	HoodRoomHeight        float64 `json:"hoodRoomHeight"`        // Hood room height
	HoodWallThickness     float64 `json:"hoodWallThickness"`     // Hood wall thickness
	DuctCenterHeight      float64 `json:"ductCenterHeight"`      // Duct center height
	DuctLeftDistance      float64 `json:"ductLeftDistance"`      // Duct left distance
	DuctDiameter          float64 `json:"ductDiameter"`          // Duct diameter
	DuctWallThickness     float64 `json:"ductWallThickness"`     // Duct wall thickness
	DuctLength            float64 `json:"ductLength"`            // Duct length
	DuctHeightDifference  float64 `json:"ductHeightDifference"`  // Duct height difference
	BaseLength            float64 `json:"baseLength"`            // EcBase length
	BaseWidth             float64 `json:"baseWidth"`             // EcBase width
	BaseHeight            float64 `json:"baseHeight"`            // EcBase height
	BaseRoomLength        float64 `json:"baseRoomLength"`        // EcBase room length
	BaseRoomWallThickness float64 `json:"baseRoomWallThickness"` // EcBase room wall thickness
	BaseRoomWidth         float64 `json:"baseRoomWidth"`         // EcBase room width
	BaseRoomHeight        float64 `json:"baseRoomHeight"`        // EcBase room height
}

// NewObliqueVentilationDuct creates a new ObliqueVentilationDuctObject
func NewObliqueVentilationDuct() *ObliqueVentilationDuct {
	return &ObliqueVentilationDuct{
		EcBase: EcBase{Type: "GIM/EC/ObliqueVentilationDuct"},
	}
}

// DrainageWellObject represents a drainage well
type DrainageWell struct {
	EcBase
	Length           float64 `json:"length"`           // Length
	Width            float64 `json:"width"`            // Width
	Height           float64 `json:"height"`           // Height
	NeckDiameter     float64 `json:"neckDiameter"`     // Neck diameter
	NeckHeight       float64 `json:"neckHeight"`       // Neck height
	CushionExtension float64 `json:"cushionExtension"` // Cushion extension
	BottomThickness  float64 `json:"bottomThickness"`  // Bottom thickness
	WallThickness    float64 `json:"wallThickness"`    // Wall thickness
}

// NewDrainageWell creates a new DrainageWellObject
func NewDrainageWell() *DrainageWell {
	return &DrainageWell{
		EcBase: EcBase{Type: "GIM/EC/DrainageWell"},
	}
}

// PipeSupportObject represents a pipe support
type PipeSupport struct {
	EcBase
	Style     int          `json:"style"`     // Type (1-2)
	Count     int          `json:"count"`     // Count
	Positions [][2]float64 `json:"positions"` // Positions
	Radii     []float64    `json:"radii"`     // Radii
	Length    float64      `json:"length"`    // Length
	Width     float64      `json:"width"`     // Width
	Height    float64      `json:"height"`    // Height
}

// NewPipeSupport creates a new PipeSupportObject
func NewPipeSupport() *PipeSupport {
	return &PipeSupport{
		EcBase: EcBase{Type: "GIM/EC/PipeSupport"},
	}
}

// CoverPlateObject represents a cover plate
type CoverPlate struct {
	EcBase
	Style       string  `json:"style"`       // Type
	Length      float64 `json:"length"`      // Length
	Width       float64 `json:"width"`       // Width
	SmallRadius float64 `json:"smallRadius"` // Small radius
	LargeRadius float64 `json:"largeRadius"` // Large radius
	Thickness   float64 `json:"thickness"`   // Thickness
}

// NewCoverPlate creates a new CoverPlateObject
func NewCoverPlate() *CoverPlate {
	return &CoverPlate{
		EcBase: EcBase{Type: "GIM/EC/CoverPlate"},
	}
}

// CableRayObject represents a cable ray
type CableRay struct {
	EcBase
	OuterLength    float64 `json:"outerLength"`    // Outer length
	OuterHeight    float64 `json:"outerHeight"`    // Outer height
	InnerLength    float64 `json:"innerLength"`    // Inner length
	InnerHeight    float64 `json:"innerHeight"`    // Inner height
	CoverThickness float64 `json:"coverThickness"` // Cover thickness
}

// NewCableRay creates a new CableRayObject
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
