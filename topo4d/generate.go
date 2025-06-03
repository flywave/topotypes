package topo4d

type GenerateType string

const (
	GenerateTypeAssemble GenerateType = "ASSEMBLE"
	GenerateTypeBuild    GenerateType = "BUILD"
)

type Generate struct {
	Type       GenerateType  `json:"type"`
	Dir        *[3]float64   `json:"dir,omitempty"`
	Radius     *float64      `json:"radius,omitempty"`
	Centerline *[][3]float64 `json:"centerline,omitempty"`
}
