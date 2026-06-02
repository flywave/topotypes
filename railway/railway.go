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

func Unmarshal(ty string, bt []byte) (Shape, error) {
	switch ty {
	case "RAILWAY/RodInsulator":
		shape := RodInsulator{}
		err := json.Unmarshal(bt, &shape)
		return &shape, err
	default:
		return nil, fmt.Errorf("invalid railway type: %s", ty)
	}
}
