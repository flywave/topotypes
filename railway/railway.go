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

// RodInsulator represents a rod insulator (棒式绝缘子)
type RodInsulator struct {
	Base
	Type               int     `json:"type"` // rod_insulator_type: 1=SOLID, 2=HOLLOW
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
	BeamLength    float64 `json:"beamLength"`
	BeamHeight    float64 `json:"beamHeight"`
	BeamWidth     float64 `json:"beamWidth"`
	BeamThickness float64 `json:"beamThickness"`
	BeamSpacing   float64 `json:"beamSpacing"`
	BraceDiameter float64 `json:"braceDiameter"`
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
	BoltSpacing          float64 `json:"boltSpacing"`
	BoltDiameter         float64 `json:"boltDiameter"`
	Height               float64 `json:"height"`
	Width                float64 `json:"width"`
	Thickness            float64 `json:"thickness"`
	InsulatorBoltSpacing  float64 `json:"insulatorBoltSpacing"`
	InsulatorBoltDiameter float64 `json:"insulatorBoltDiameter"`
	MountAngle           float64 `json:"mountAngle"`
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

// OcsFoundation represents an OCS foundation (支柱基础)
type OcsFoundation struct {
	Base
	Type            int     `json:"type"`
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
	Type            int     `json:"type"`
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
}

func NewConcreteMast() *ConcreteMast {
	return &ConcreteMast{
		Base: Base{Type: "RAILWAY/ConcreteMast"},
	}
}

// RegistrationArm represents a registration arm (定位器)
type RegistrationArm struct {
	Base
	Type          int     `json:"type"`
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

func Unmarshal(ty string, bt []byte) (Shape, error) {
	switch ty {
	case "RAILWAY/RodInsulator":
		shape := RodInsulator{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "RAILWAY/CrossArm":
		shape := CrossArm{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "RAILWAY/LevelCantilever":
		shape := LevelCantilever{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "RAILWAY/SlantCantilever":
		shape := SlantCantilever{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "RAILWAY/CantileverBrace":
		shape := CantileverBrace{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "RAILWAY/CurvedArm":
		shape := CurvedArm{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "RAILWAY/RegArmBracket":
		shape := RegArmBracket{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "RAILWAY/ContactWire":
		shape := ContactWire{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "RAILWAY/MessengerWire":
		shape := MessengerWire{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "RAILWAY/GuyWire":
		shape := GuyWire{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "RAILWAY/MastBracket":
		shape := MastBracket{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "RAILWAY/SteelMast":
		shape := SteelMast{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "RAILWAY/ConcreteMast":
		shape := ConcreteMast{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	case "RAILWAY/RegistrationArm":
		shape := RegistrationArm{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	default:
		return nil, fmt.Errorf("invalid railway type: %s", ty)
	}
}
