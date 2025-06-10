package hydropower

import (
	"encoding/json"
	"errors"
)

const Major = "HYDROPOWER"

type Shape interface {
	GetType() string
}
type Base struct {
	Version int    `json:"version"`
	Type    string `json:"type"`
}

func (b *Base) GetType() string {
	return b.Type
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
	Position [3]float64
	Type     ChannelPointType
}

type TunnelStyle string

const (
	TunnelStyleRectangular TunnelStyle = "RECTANGULAR"
	TunnelStyleCityOpening TunnelStyle = "CITYOPENING"
	TunnelStyleCircular    TunnelStyle = "CIRCULAR"
	TunnelStyleHorseshoe   TunnelStyle = "HORSESHOE"
)

func (t TunnelStyle) ToInt() int {
	switch t {
	case TunnelStyleRectangular:
		return 1
	case TunnelStyleCityOpening:
		return 2
	case TunnelStyleCircular:
		return 3
	case TunnelStyleHorseshoe:
		return 4
	default:
		return 1
	}
}

type PointIndex struct {
	Index int
	Type  ChannelPointType
}

type WaterTunnel struct {
	Base
	Style                TunnelStyle     `json:"style"`                          // 'RECTANGULAR' | 'CITYOPENING' | 'CIRCULAR' | 'HORSESHOE'
	Width                float64         `json:"width"`                          // Width of the tunnel
	Height               float64         `json:"height"`                         // Height of the tunnel
	TopThickness         float64         `json:"topThickness"`                   // Thickness of the top part
	BottomThickness      float64         `json:"bottomThickness"`                // Thickness of the bottom part
	OuterWallThickness   float64         `json:"outerWallThickness"`             // Thickness of outer wall
	InnerWallThickness   float64         `json:"innerWallThickness"`             // Thickness of inner wall
	ArcHeight            float64         `json:"arcHeight,omitempty"`            // Height of the arc (for certain styles)
	ArcRadius            float64         `json:"arcRadius,omitempty"`            // Radius of the arc (for certain styles)
	ArcAngle             float64         `json:"arcAngle,omitempty"`             // Angle of the arc (for certain styles)
	BottomPlatformHeight float64         `json:"bottomPlatformHeight,omitempty"` // Height of bottom platform
	CushionExtension     float64         `json:"cushionExtension,omitempty"`     // Extension of cushion
	CushionThickness     float64         `json:"cushionThickness,omitempty"`     // Thickness of cushion
	Points               []*ChannelPoint `json:"-"`                              // Path points defining the tunnel
	PointIndexs          []PointIndex    `json:"pointIndex"`                     // Index of points defining the tunnel
}

// NewWaterTunnel creates a new WaterTunnel instance
func NewWaterTunnel() *WaterTunnel {
	return &WaterTunnel{
		Base: Base{Type: "HYDROPOWER/WaterTunnel"},
	}
}

func Unmarshal(ty string, bt []byte) (Shape, error) {
	if ty == "HYDROPOWER/WaterTunnel" {
		var t WaterTunnel
		err := json.Unmarshal(bt, &t)
		if err != nil {
			return nil, err
		}
		return &t, nil
	}
	return nil, errors.New("invalid type")
}
