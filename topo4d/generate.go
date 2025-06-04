package topo4d

type GenerateType string

const (
	GenerateTypeAssemble GenerateType = "ASSEMBLE" // 装配
	GenerateTypeBuild    GenerateType = "BUILD"    // 建造
)

type Generate4D struct {
	Type       GenerateType  `json:"type"`
	Direction  *[3]float64   `json:"direction,omitempty"`
	Radius     *float64      `json:"radius,omitempty"`
	Centerline *[][3]float64 `json:"centerline,omitempty"`
}
