package railway

import (
	"encoding/json"
	"fmt"
)

const Major = "RAILWAY"

type Shape interface {
	GetType() string
}

type Base struct {
	Version int    `json:"version"`
	Type    string `json:"type"`
}

func (b *Base) GetType() string { return b.Type }

// 枚举常量, 与 go-topo src/primitives_railway.hh 一一对应

// rod_insulator_type
const (
	RodInsulatorSolid  = 1 // 实心棒式
	RodInsulatorHollow = 2 // 空心棒式
)

// end_fitting_type
const (
	EndFittingFlange = 1 // 法兰式
	EndFittingBall   = 2 // 球头
	EndFittingScrew  = 3 // 螺杆式
)

// cross_arm_type
const (
	CrossArmDoubleFork    = 1 // 双叉臂固定横担
	CrossArmTriangleBrace = 2 // 三角撑横担
	CrossArmTruss         = 3 // 桁架式横担
)

// curved_arm_type
const (
	CurvedArmArc    = 1 // 弧形弯臂
	CurvedArmLShape = 2 // L形弯臂
	CurvedArmDouble = 3 // 双弯臂
)

// registration_arm_type
const (
	RegistrationArmStraight = 1 // 直型
	RegistrationArmCurved   = 2 // 弯型
	RegistrationArmExtended = 3 // 加长型
)

// steel_mast_type
const (
	SteelMastLattice = 1 // 格构式钢柱
	SteelMastHBeam   = 2 // H型钢柱
)

// concrete_mast_section_type
const (
	ConcreteMastCircular         = 1 // 环形
	ConcreteMastRectangular      = 2 // 矩形
	ConcreteMastCircularHoled    = 3 // 带孔环形
	ConcreteMastRectangularHoled = 4 // 矩形挖孔
)

// foundation_type
const (
	FoundationDirectBuried  = 1 // 直埋式
	FoundationFlange        = 2 // 法兰盘基础
	FoundationBoredPile     = 3 // 钻孔灌注桩
	FoundationExcavatedPile = 4 // 挖孔桩
	FoundationAnchor        = 5 // 锚栓基础
)

// anchor_fitting_type
const (
	AnchorFittingRodAndRing = 1 // 杵环杆
	AnchorFittingDoubleEar  = 2 // 双耳连接器
	AnchorFittingWedgeClamp = 3 // 楔形线夹
)

// beam_section_type (硬横跨梁截面)
const (
	BeamSectionBox   = 1 // 箱型
	BeamSectionHBeam = 2 // H型
	BeamSectionTruss = 3 // 桁架式
	BeamSectionCombo = 4 // 组合式
)

// hanger_post_section_type
const (
	HangerPostRound  = 1 // 圆管
	HangerPostSquare = 2 // 方管
	HangerPostHBeam  = 3 // H型钢
)

// aux_bracket_type
const (
	AuxBracketCrossArm   = 1 // 横担式
	AuxBracketWallMount  = 2 // 壁挂式
	AuxBracketDoubleMast = 3 // 双支柱式
)

// sleeper_shape_type
const (
	SleeperRectangular = 1 // 矩形 (木枕/简易)
	SleeperTrapezoidal = 2 // 梯形收腰 (混凝土枕)
)

// mast_assembly 支柱/腕臂/补偿装置类型
const (
	MastAssemblyLattice  = 1 // 格构式钢柱
	MastAssemblyConcrete = 2 // 混凝土柱

	CantileverNone   = 0 // 无腕臂
	CantileverSingle = 1 // 单臂
	CantileverDouble = 2 // 双臂

	CompensatorNone    = 0 // 无补偿
	CompensatorRatchet = 1 // 棘轮
	CompensatorPulley  = 2 // 滑轮
)

// retarder_point 减速顶安装侧/设备类型/安装方式
const (
	RetarderSideLeft  = 1
	RetarderSideRight = 2
	RetarderSideBoth  = 3

	RetarderDeviceHydraulic    = 1 // 液压
	RetarderDeviceFriction     = 2 // 摩擦
	RetarderDeviceControllable = 3 // 可控

	RetarderMountInner = 1 // 轨内侧
	RetarderMountOuter = 2 // 轨外侧
	RetarderMountBoth  = 3 // 双轨双侧
)

// RodInsulator represents a rod insulator (棒式绝缘子)
type RodInsulator struct {
	Base
	InsulatorType      int     `json:"insulatorType"` // rod_insulator_type: 1=SOLID, 2=HOLLOW
	Height             float64 `json:"height"`
	OuterDiameter      float64 `json:"outerDiameter"`
	InnerDiameter      float64 `json:"innerDiameter"`
	ShedDiameter       float64 `json:"shedDiameter"`
	ShedSpacing        float64 `json:"shedSpacing"`
	ShedCount          int     `json:"shedCount"`
	EndFitting         int     `json:"endFitting"` // end_fitting_type: 1=FLANGE, 2=BALL, 3=SCREW
	FlangeDiameter     float64 `json:"flangeDiameter"`
	FlangeBoltSpacing  float64 `json:"flangeBoltSpacing"`
	FlangeBoltDiameter float64 `json:"flangeBoltDiameter"`
}

func NewRodInsulator() *RodInsulator {
	return &RodInsulator{
		Base: Base{Type: "RAILWAY/RodInsulator"},
	}
}

// CrossArm represents a cross arm (横担)
type CrossArm struct {
	Base
	CrossArmType  int     `json:"crossArmType"` // cross_arm_type: 1=双叉臂, 2=三角撑, 3=桁架式
	BeamLength    float64 `json:"beamLength"`
	BeamHeight    float64 `json:"beamHeight"`
	BeamWidth     float64 `json:"beamWidth"`
	BeamThickness float64 `json:"beamThickness"`
	BeamSpacing   float64 `json:"beamSpacing"`
	BraceLength   float64 `json:"braceLength"`
	BraceDiameter float64 `json:"braceDiameter"`
	MountHeight   float64 `json:"mountHeight"`
	BoltSpacing   float64 `json:"boltSpacing"`
	BoltDiameter  float64 `json:"boltDiameter"`
	BoltCount     int     `json:"boltCount"`
}

func NewCrossArm() *CrossArm {
	return &CrossArm{
		Base: Base{Type: "RAILWAY/CrossArm"},
	}
}

// LevelCantilever represents a level cantilever (平腕臂)
type LevelCantilever struct {
	Base
	Length        float64 `json:"length"`
	OuterDiameter float64 `json:"outerDiameter"`
	WallThickness float64 `json:"wallThickness"`
	MountHeight   float64 `json:"mountHeight"`
	RiseAngle     float64 `json:"riseAngle"`
}

func NewLevelCantilever() *LevelCantilever {
	return &LevelCantilever{
		Base: Base{Type: "RAILWAY/LevelCantilever"},
	}
}

// SlantCantilever represents a slanted cantilever (斜腕臂)
type SlantCantilever struct {
	Base
	Length        float64 `json:"length"`
	OuterDiameter float64 `json:"outerDiameter"`
	WallThickness float64 `json:"wallThickness"`
	SlantAngle    float64 `json:"slantAngle"`
}

func NewSlantCantilever() *SlantCantilever {
	return &SlantCantilever{
		Base: Base{Type: "RAILWAY/SlantCantilever"},
	}
}

// CantileverBrace represents a cantilever brace (斜撑)
type CantileverBrace struct {
	Base
	Length        float64 `json:"length"`
	OuterDiameter float64 `json:"outerDiameter"`
	WallThickness float64 `json:"wallThickness"`
	SlantAngle    float64 `json:"slantAngle"`
}

func NewCantileverBrace() *CantileverBrace {
	return &CantileverBrace{
		Base: Base{Type: "RAILWAY/CantileverBrace"},
	}
}

// CurvedArm represents a curved arm (弯臂)
type CurvedArm struct {
	Base
	CurvedArmType    int     `json:"curvedArmType"` // curved_arm_type: 1=弧形, 2=L形, 3=双弯臂
	VerticalLength   float64 `json:"verticalLength"`
	HorizontalLength float64 `json:"horizontalLength"`
	BendRadius       float64 `json:"bendRadius"`
	BendAngle        float64 `json:"bendAngle"`
	OuterDiameter    float64 `json:"outerDiameter"`
	WallThickness    float64 `json:"wallThickness"`
	FlangeThickness  float64 `json:"flangeThickness"`
	BoltSpacing      float64 `json:"boltSpacing"`
	BoltDiameter     float64 `json:"boltDiameter"`
}

func NewCurvedArm() *CurvedArm {
	return &CurvedArm{
		Base: Base{Type: "RAILWAY/CurvedArm"},
	}
}

// RegArmBracket represents a registration arm bracket (定位器底座 L型金具)
type RegArmBracket struct {
	Base
	TubeDiameter      float64 `json:"tubeDiameter"`
	BandWidth         float64 `json:"bandWidth"`
	BandThickness     float64 `json:"bandThickness"`
	BracketHeight     float64 `json:"bracketHeight"`
	BracketThickness  float64 `json:"bracketThickness"`
	BracketWidth      float64 `json:"bracketWidth"`
	MountHoleDiameter float64 `json:"mountHoleDiameter"`
}

func NewRegArmBracket() *RegArmBracket {
	return &RegArmBracket{
		Base: Base{Type: "RAILWAY/RegArmBracket"},
	}
}

// ContactWire represents a contact wire (接触线)
type ContactWire struct {
	Base
	SectionalArea float64 `json:"sectionalArea"`
	Diameter      float64 `json:"diameter"`
	RatedTension  float64 `json:"ratedTension"`
	GrooveDepth   float64 `json:"grooveDepth"`
	GrooveWidth   float64 `json:"grooveWidth"`
	BottomRadius  float64 `json:"bottomRadius"`
	TopRadius     float64 `json:"topRadius"`
	Sag           float64 `json:"sag"`
}

func NewContactWire() *ContactWire {
	return &ContactWire{
		Base: Base{Type: "RAILWAY/ContactWire"},
	}
}

// MessengerWire represents a messenger wire (承力索)
type MessengerWire struct {
	Base
	Diameter         float64 `json:"diameter"`
	RatedTension     float64 `json:"ratedTension"`
	StructuralHeight float64 `json:"structuralHeight"`
	Sag              float64 `json:"sag"`
}

func NewMessengerWire() *MessengerWire {
	return &MessengerWire{
		Base: Base{Type: "RAILWAY/MessengerWire"},
	}
}

// MastBracket represents a mast bracket (支柱连接座)
type MastBracket struct {
	Base
	BoltSpacing           float64 `json:"boltSpacing"`
	BoltDiameter          float64 `json:"boltDiameter"`
	Height                float64 `json:"height"`
	Width                 float64 `json:"width"`
	Thickness             float64 `json:"thickness"`
	InsulatorBoltSpacing  float64 `json:"insulatorBoltSpacing"`
	InsulatorBoltDiameter float64 `json:"insulatorBoltDiameter"`
	MountAngle            float64 `json:"mountAngle"`
	MastDiameter          float64 `json:"mastDiameter"`
}

func NewMastBracket() *MastBracket {
	return &MastBracket{
		Base: Base{Type: "RAILWAY/MastBracket"},
	}
}

// GuyWire represents a guy wire (下锚拉线)
type GuyWire struct {
	Base
	Length            float64 `json:"length"`
	Diameter          float64 `json:"diameter"`
	Angle             float64 `json:"angle"`
	RatedTension      float64 `json:"ratedTension"`
	HasInsulator      bool    `json:"hasInsulator"`
	InsulatorCount    int     `json:"insulatorCount"`
	AnchorRodDiameter float64 `json:"anchorRodDiameter"`
	AnchorRodLength   float64 `json:"anchorRodLength"`
	AnchorPlateLength float64 `json:"anchorPlateLength"`
	AnchorPlateWidth  float64 `json:"anchorPlateWidth"`
}

func NewGuyWire() *GuyWire {
	return &GuyWire{
		Base: Base{Type: "RAILWAY/GuyWire"},
	}
}

// Dropper represents a dropper (吊弦)
type Dropper struct {
	Base
	Length         float64 `json:"length"`
	WireDiameter   float64 `json:"wireDiameter"`
	ClampLength    float64 `json:"clampLength"`
	ClampWidth     float64 `json:"clampWidth"`
	ClampThickness float64 `json:"clampThickness"`
	Conductive     bool    `json:"conductive"`
}

func NewDropper() *Dropper {
	return &Dropper{
		Base: Base{Type: "RAILWAY/Dropper"},
	}
}

// OcsFoundation represents an OCS foundation (支柱基础)
type OcsFoundation struct {
	Base
	FoundationType  int     `json:"foundationType"` // foundation_type: 1=直埋, 2=法兰盘, 3=钻孔灌注桩, 4=挖孔桩, 5=锚栓
	Height          float64 `json:"height"`
	Width           float64 `json:"width"`
	Length          float64 `json:"length"`
	FlangeThickness float64 `json:"flangeThickness"`
	AnchorCount     int     `json:"anchorCount"`
	AnchorDiameter  float64 `json:"anchorDiameter"`
	AnchorLength    float64 `json:"anchorLength"`
	AnchorSpacing   float64 `json:"anchorSpacing"`
}

func NewOcsFoundation() *OcsFoundation {
	return &OcsFoundation{
		Base: Base{Type: "RAILWAY/OcsFoundation"},
	}
}

// SteelMast represents a steel mast (钢支柱)
type SteelMast struct {
	Base
	MastType        int     `json:"mastType"` // steel_mast_type: 1=格构式, 2=H型钢
	Height          float64 `json:"height"`
	TopWidth        float64 `json:"topWidth"`
	BottomWidth     float64 `json:"bottomWidth"`
	WallThickness   float64 `json:"wallThickness"`
	FlangeThickness float64 `json:"flangeThickness"`
	FlangeWidth     float64 `json:"flangeWidth"`
	AnchorSpacing   float64 `json:"anchorSpacing"`
	AnchorDiameter  float64 `json:"anchorDiameter"`
	SegmentCount    int     `json:"segmentCount"`
}

func NewSteelMast() *SteelMast {
	return &SteelMast{
		Base: Base{Type: "RAILWAY/SteelMast"},
	}
}

// ConcreteMast represents a concrete mast (混凝土支柱)
type ConcreteMast struct {
	Base
	SectionType     int     `json:"sectionType"`
	Height          float64 `json:"height"`
	TopWidth        float64 `json:"topWidth"`
	BottomWidth     float64 `json:"bottomWidth"`
	WallThickness   float64 `json:"wallThickness"`
	HoleDiameter    float64 `json:"holeDiameter"`
	HoleSpacingV    float64 `json:"holeSpacingV"`
	HoleSpacingH    float64 `json:"holeSpacingH"`
	FirstHoleOffset float64 `json:"firstHoleOffset"`
	HoleRowCount    int     `json:"holeRowCount"`
	HolesPerRow     int     `json:"holesPerRow"`
	HoleLength      float64 `json:"holeLength"`
}

func NewConcreteMast() *ConcreteMast {
	return &ConcreteMast{
		Base: Base{Type: "RAILWAY/ConcreteMast"},
	}
}

// RegistrationArm represents a registration arm (定位器)
type RegistrationArm struct {
	Base
	ArmType       int     `json:"armType"` // registration_arm_type: 1=直型, 2=弯型, 3=加长型
	Length        float64 `json:"length"`
	TubeWidth     float64 `json:"tubeWidth"`
	TubeHeight    float64 `json:"tubeHeight"`
	WallThickness float64 `json:"wallThickness"`
	Angle         float64 `json:"angle"`
	IsReverse     bool    `json:"isReverse"`
}

func NewRegistrationArm() *RegistrationArm {
	return &RegistrationArm{
		Base: Base{Type: "RAILWAY/RegistrationArm"},
	}
}

// Unmarshal 按 "RAILWAY/Xxx" 判别串反序列化; Base.Type 是唯一 "type" 字段
func Unmarshal(ty string, bt []byte) (Shape, error) {
	var shape Shape
	switch ty {
	case "RAILWAY/RodInsulator":
		shape = &RodInsulator{}
	case "RAILWAY/CrossArm":
		shape = &CrossArm{}
	case "RAILWAY/LevelCantilever":
		shape = &LevelCantilever{}
	case "RAILWAY/SlantCantilever":
		shape = &SlantCantilever{}
	case "RAILWAY/CantileverBrace":
		shape = &CantileverBrace{}
	case "RAILWAY/CurvedArm":
		shape = &CurvedArm{}
	case "RAILWAY/RegArmBracket":
		shape = &RegArmBracket{}
	case "RAILWAY/ContactWire":
		shape = &ContactWire{}
	case "RAILWAY/MessengerWire":
		shape = &MessengerWire{}
	case "RAILWAY/GuyWire":
		shape = &GuyWire{}
	case "RAILWAY/Dropper":
		shape = &Dropper{}
	case "RAILWAY/MastBracket":
		shape = &MastBracket{}
	case "RAILWAY/SteelMast":
		shape = &SteelMast{}
	case "RAILWAY/ConcreteMast":
		shape = &ConcreteMast{}
	case "RAILWAY/OcsFoundation":
		shape = &OcsFoundation{}
	case "RAILWAY/RegistrationArm":
		shape = &RegistrationArm{}
	case "RAILWAY/CantileverBase":
		shape = &CantileverBase{}
	case "RAILWAY/MWSaddle":
		shape = &MWSaddle{}
	case "RAILWAY/BalanceWeight":
		shape = &BalanceWeight{}
	case "RAILWAY/WeightRod":
		shape = &WeightRod{}
	case "RAILWAY/AnchorFitting":
		shape = &AnchorFitting{}
	case "RAILWAY/Crossing":
		shape = &Crossing{}
	case "RAILWAY/HeadSpan":
		shape = &HeadSpan{}
	case "RAILWAY/TransverseSpan":
		shape = &TransverseSpan{}
	case "RAILWAY/HangerPost":
		shape = &HangerPost{}
	case "RAILWAY/PortalFrame":
		shape = &PortalFrame{}
	case "RAILWAY/SuspensionHardSpan":
		shape = &SuspensionHardSpan{}
	case "RAILWAY/PositioningCable":
		shape = &PositioningCable{}
	case "RAILWAY/AuxBracket":
		shape = &AuxBracket{}
	case "RAILWAY/Rail":
		shape = &Rail{}
	case "RAILWAY/Sleeper":
		shape = &Sleeper{}
	case "RAILWAY/Ballast":
		shape = &Ballast{}
	case "RAILWAY/TrackSlab":
		shape = &TrackSlab{}
	case "RAILWAY/Fastener":
		shape = &Fastener{}
	case "RAILWAY/GuardRail":
		shape = &GuardRail{}
	case "RAILWAY/MastAssembly":
		shape = &MastAssembly{}
	case "RAILWAY/WeightStack":
		shape = &WeightStack{}
	case "RAILWAY/RatchetCompensator":
		shape = &RatchetCompensator{}
	case "RAILWAY/AuxiliaryWire":
		shape = &AuxiliaryWire{}
	case "RAILWAY/Disconnector":
		shape = &Disconnector{}
	case "RAILWAY/Arrester":
		shape = &Arrester{}
	case "RAILWAY/PulleyCompensator":
		shape = &PulleyCompensator{}
	case "RAILWAY/SleeveConnector":
		shape = &SleeveConnector{}
	case "RAILWAY/SleeveEar":
		shape = &SleeveEar{}
	case "RAILWAY/SwitchRail":
		shape = &SwitchRail{}
	case "RAILWAY/Frog":
		shape = &Frog{}
	case "RAILWAY/Turnout":
		shape = &Turnout{}
	case "RAILWAY/StraightTrack":
		shape = &StraightTrack{}
	case "RAILWAY/CurveTrack":
		shape = &CurveTrack{}
	case "RAILWAY/RailPair":
		shape = &RailPair{}
	case "RAILWAY/SleeperLayout":
		shape = &SleeperLayout{}
	case "RAILWAY/RetarderPoint":
		shape = &RetarderPoint{}
	case "RAILWAY/AnchorSection":
		shape = &AnchorSection{}
	case "RAILWAY/Yard":
		shape = &Yard{}
	default:
		return nil, fmt.Errorf("invalid railway type: %s", ty)
	}
	if err := json.Unmarshal(bt, shape); err != nil {
		return nil, err
	}
	return shape, nil
}
