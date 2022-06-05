package editor

import (
	quatd "github.com/flywave/go3d/float64/quaternion"
	vec3d "github.com/flywave/go3d/float64/vec3"
)

type Metadata struct {
	Scale       float64      `json:"scale"`
	Rotation    quatd.T      `json:"rotation"`
	Offset      vec3d.T      `json:"offset"`
	Anchors     []*Anchor    `json:"anchors"`
	AnchorCount int          `json:"anchorcount"`
	Boards      []*Board     `json:"boards"`
	BoardCount  int          `json:"boardcount"`
	Components  []*Component `json:"components,omitempty"`
}
