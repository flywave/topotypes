package editor

import vec3d "github.com/flywave/go3d/float64/vec3"

type Anchor struct {
	Id             string  `json:"id"`
	Center         vec3d.T `json:"center"`
	Name           string  `json:"name"`
	Normal         vec3d.T `json:"normal"`
	Unit           float64 `json:"unit"`
	AngleFree      *bool   `json:"angle_free,omitempty"`
	ScaleFree      *bool   `json:"scale_free,omitempty"`
	Source         string  `json:"source"`
	SourceAnchorId string  `json:"source_anchor_id"`
}
