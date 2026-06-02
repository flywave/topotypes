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
	case "RAILWAY/RegArmBracket":
		shape := RegArmBracket{}
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
